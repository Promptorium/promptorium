package gitcontext

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// GitContext

type GitContext struct {
	IsGitRepo       bool
	IsDirty         bool
	IsDetachedHead  bool
	LocalBranch     string
	UpstreamBranch  string
	Remote          string
	Ahead           int
	Behind          int
	UnstagedChanges int
	StagedChanges   int
	UntrackedFiles  int
}

type changes struct {
	StagedChanges   int
	UnstagedChanges int
	UntrackedFiles  int
}

func GetGitState() GitContext {
	// Get the git context by shelling out to git
	cmd := exec.Command("git", "status", "--porcelain=v2", "--branch", "-z")
	output, err := cmd.Output()
	if err != nil {
		log.Trace().Msgf("Error getting git context: %s", err)
		return GitContext{
			IsGitRepo: false,
		}
	}
	changes, err := getchanges(output)
	if err != nil {
		log.Trace().Msgf("Error parsing git context: %s", err)
		return GitContext{
			IsGitRepo: false,
		}
	}
	localBranch := getLocalBranch(output)
	upstreamBranch, remote := getUpstreamBranchAndRemote(output)
	ahead, behind := getAheadBehind(output)

	// Return the git context
	log.Trace().Msgf("Found git repo")
	log.Trace().Msgf("Git remote: %s", remote)
	log.Trace().Msgf("Git branch: %s", localBranch)
	log.Trace().Msgf("Git upstream: %s", upstreamBranch)
	log.Trace().Msgf("Git ahead: %d, behind: %d, unstaged changes: %d, staged changes: %d, untracked files: %d", ahead, behind, changes.UnstagedChanges, changes.StagedChanges, changes.UntrackedFiles)

	isDirty := ahead > 0 || behind > 0 || changes.StagedChanges > 0 || changes.UnstagedChanges > 0 || changes.UntrackedFiles > 0

	return GitContext{
		IsGitRepo:       true,
		IsDirty:         isDirty,
		IsDetachedHead:  false,
		LocalBranch:     localBranch,
		UpstreamBranch:  upstreamBranch,
		Remote:          remote,
		Ahead:           ahead,
		Behind:          behind,
		UnstagedChanges: changes.UnstagedChanges,
		StagedChanges:   changes.StagedChanges,
		UntrackedFiles:  changes.UntrackedFiles,
	}
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

func getUpstreamBranchAndRemote(gitStatus []byte) (string, string) {
	upstreamBranch := ""
	remote := ""
	for _, line := range strings.Split(string(gitStatus), "\x00") {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "# branch.upstream") {
			// Header lines are parsed in another function
			upstream := strings.TrimPrefix(line, "# branch.upstream ")
			remote = strings.Split(upstream, "/")[0]
			upstreamBranch = strings.Split(upstream, "/")[1]
			break
		}
	}
	return upstreamBranch, remote
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
