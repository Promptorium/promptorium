package gitcontext

import (
	"os/exec"
	"promptorium/internal/utils"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// GitContext

type GitContext struct {
	IsGitRepo       bool
	IsDirty         bool
	IsDetachedHead  bool
	HasUpstream     bool
	LocalBranch     string
	UpstreamBranch  string
	Remote          string
	Ahead           int
	Behind          int
	UnstagedChanges int
	StagedChanges   int
	UntrackedFiles  int

	gitRoot utils.CachedData[string]
	GitRoot func() string
}

type changes struct {
	StagedChanges   int
	UnstagedChanges int
	UntrackedFiles  int
}

func GetGitState(gitContext chan GitContext) {
	// Get the git context by shelling out to git
	cmd := exec.Command("git", "status", "--porcelain=v2", "--branch", "-z")
	output, err := cmd.Output()
	if err != nil {
		log.Trace().Msgf("Error getting git context: %s", err)
		gitContext <- GitContext{
			IsGitRepo: false,
		}
		return
	}
	changes, err := getchanges(output)
	if err != nil {
		log.Trace().Msgf("Error parsing git context: %s", err)
		gitContext <- GitContext{
			IsGitRepo: false,
		}
		return
	}
	localBranch := getLocalBranch(output)
	upstreamBranch, remote, hasUpstream := getUpstreamBranchAndRemote(output)
	ahead, behind := getAheadBehind(output)

	// Return the git context
	log.Trace().Msgf("Found git repo")
	log.Trace().Msgf("Git remote: %s", remote)
	log.Trace().Msgf("Git branch: %s", localBranch)
	log.Trace().Msgf("Git upstream: %s", upstreamBranch)
	log.Trace().Msgf("Git ahead: %d, behind: %d, unstaged changes: %d, staged changes: %d, untracked files: %d", ahead, behind, changes.UnstagedChanges, changes.StagedChanges, changes.UntrackedFiles)

	isDirty := ahead > 0 || behind > 0 || changes.StagedChanges > 0 || changes.UnstagedChanges > 0 || changes.UntrackedFiles > 0

	result := GitContext{
		IsGitRepo:       true,
		IsDirty:         isDirty,
		IsDetachedHead:  false,
		HasUpstream:     hasUpstream,
		LocalBranch:     localBranch,
		UpstreamBranch:  upstreamBranch,
		Remote:          remote,
		Ahead:           ahead,
		Behind:          behind,
		UnstagedChanges: changes.UnstagedChanges,
		StagedChanges:   changes.StagedChanges,
		UntrackedFiles:  changes.UntrackedFiles,
		gitRoot:         utils.NewCachedData[string](getGitRoot, "git root"),
	}

	result.GitRoot = func() string { return result.gitRoot.GetContent() }

	gitContext <- result
}

func getchanges(gitStatus []byte) (changes, error) {
	stagedChanges, unstagedChanges, untrackedFiles := 0, 0, 0
	var changes changes

	//Parse git status string
	// We are parsing the output of git status --porcelain=v2
	for _, line := range strings.Split(string(gitStatus), "\x00") {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#") {
			// Header lines are parsed in another function
			continue
		}
		if strings.HasPrefix(line, "?") {
			untrackedFiles++
			continue
		}
		if strings.HasPrefix(line, "1") || strings.HasPrefix(line, "2") {
			// This means that there is a change in the file
			// Now we need to parse the second part of the string to know whether the change is staged or not

			XY := strings.Split(line, " ")[1]
			if !strings.HasSuffix(XY, ".") {
				unstagedChanges++
			}
			if !strings.HasPrefix(XY, ".") {
				stagedChanges++
			}
			continue
		}
	}
	changes.StagedChanges = stagedChanges
	changes.UnstagedChanges = unstagedChanges
	changes.UntrackedFiles = untrackedFiles
	return changes, nil
}

func getAheadBehind(gitStatus []byte) (int, int) {
	ahead, behind := 0, 0
	for _, line := range strings.Split(string(gitStatus), "\x00") {
		if strings.HasPrefix(line, "# branch.ab") {
			for _, word := range strings.Split(line, " ") {
				if strings.HasPrefix(word, "+") {
					aheadString, _ := strings.CutPrefix(word, "+")
					ahead, _ = strconv.Atoi(aheadString)
				}
				if strings.HasPrefix(word, "-") {
					behindString, _ := strings.CutPrefix(word, "-")
					behind, _ = strconv.Atoi(behindString)
				}
			}
			break
		}
	}
	return ahead, behind
}

func getUpstreamBranchAndRemote(gitStatus []byte) (string, string, bool) {
	upstreamBranch := ""
	remote := ""
	hasUpstream := false
	for _, line := range strings.Split(string(gitStatus), "\x00") {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "# branch.upstream") {
			// Header lines are parsed in another function
			upstream := strings.TrimPrefix(line, "# branch.upstream ")
			remote = strings.Split(upstream, "/")[0]
			upstreamBranch = strings.Split(upstream, "/")[1]
			hasUpstream = true
			break
		}
	}
	return upstreamBranch, remote, hasUpstream
}

func getLocalBranch(gitStatus []byte) string {
	localBranch := ""
	for _, line := range strings.Split(string(gitStatus), "\x00") {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "# branch.head ") {
			// Header lines are parsed in another function
			localBranch = strings.TrimPrefix(line, "# branch.head ")
			break
		}
	}
	return localBranch
}

func getGitRoot(result chan string) {
	gitRoot, _ := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	result <- strings.TrimSuffix(string(gitRoot), "\n")
}
