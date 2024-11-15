package confmodule

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

type ApplicationState struct {
	gitStatus       CachedState[string]
	GitStatus       CachedState[string]
	GitRemote       CachedState[string]
	GitBranch       CachedState[string]
	GitRemoteBranch CachedState[string]
	ExitCode        CachedState[int]
	OS              CachedState[string]
	Shell           CachedState[string]
}

func getApplicationState(shell string, exitCode int) ApplicationState {
	state := ApplicationState{}
	state.GitRemote = NewCachedState[string](state.getGitRemote)
	state.GitRemoteBranch = NewCachedState[string](state.getGitRemoteBranch)
	state.GitBranch = NewCachedState[string](state.getGitBranch)
	state.GitStatus = NewCachedState[string](state.getGitStatus)
	state.ExitCode = NewCachedState[int](func() int { return exitCode })
	state.OS = NewCachedState[string](state.getOS)
	state.Shell = NewCachedState[string](func() string { return shell })
	return state
}

func (state *ApplicationState) getGitStatus() string {
	stdErr := bytes.Buffer{}
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Stderr = &stdErr

	output, _ := cmd.Output()

	status := strings.Replace(string(output), "\n", "", -1)

	if strings.Contains(stdErr.String(), "fatal:") {
		return "no_branch"
	}
	if state.GitRemoteBranch.GetContent() == "" {
		return "no_remote"
	}

	if status == "" {
		return "clean"
	}

	return "dirty"
}

func (state *ApplicationState) getGitBranch() string {
	output, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		return ""
	}
	branch := strings.Replace(string(output), "\n", "", -1)
	if branch == "" {
		return "master"
	}
	return branch
}

func (state *ApplicationState) getGitRemote() string {
	output, err := exec.Command("git", "remote").Output()
	if err != nil {
		return ""
	}
	remote := strings.Replace(string(output), "\n", "", -1)
	return remote
}

func (state *ApplicationState) getGitRemoteBranch() string {
	output, err := exec.Command("git", "status", "-sb").Output()
	if err != nil {
		return ""
	}
	remote := state.GitRemote.GetContent()
	remoteBranch := ""
	for _, line := range strings.Split(string(output), "\n") {
		if strings.HasPrefix(line, "## ") {
			for _, remoteBranch := range strings.Split(line, "...") {
				if strings.HasPrefix(remoteBranch, remote+"/") {
					log.Debug().Msgf("Remote branch: %s", remoteBranch)
					remoteBranch = strings.TrimPrefix(remoteBranch, remote+"/")
					return remoteBranch
				}
			}
		}
	}
	log.Debug().Msgf("No remote branch found")
	return remoteBranch
}
func (state *ApplicationState) getOS() string {
	output, err := exec.Command("uname", "-s").Output()
	if err != nil {
		return ""
	}
	os := strings.Replace(string(output), "\n", "", -1)
	if os == "" {
		return "unknown"
	}
	// Return directly if not Linux
	if os != "Linux" {
		return os
	}
	// Get Linux distribution
	output, err = exec.Command("cat", "/etc/os-release").Output()
	if err != nil {
		fmt.Println("Error getting Linux distribution:", err)
		return "unknown"
	}
	entries := strings.Split(string(output), "\n")
	for _, entry := range entries {
		if strings.HasPrefix(entry, "ID=") {
			os = strings.Split(entry, "=")[1]
			break
		}
	}
	// TODO: Check if the distribution is supported

	return os
}

func (state *ApplicationState) getExitCode(exitCode int) int {
	// Get exit code from last command
	if exitCode != 0 {
		return 1
	}
	return 0
}

func (state *ApplicationState) getShell(shell string) string {
	shell = filepath.Base(shell)
	switch shell {
	case "bash":
		return "bash"
	case "zsh":
		return "zsh"
	default:
		return "unknown"
	}
}

type CachedState[T any] struct {
	content  T
	isCached bool
	refresh  func() T
}

func NewCachedState[T any](refresh func() T) CachedState[T] {
	return CachedState[T]{
		isCached: false,
		refresh:  refresh,
	}
}

func (c *CachedState[T]) GetContent() T {
	if !c.isCached {
		c.content = c.refresh()
		c.isCached = true
	}
	return c.content
}
