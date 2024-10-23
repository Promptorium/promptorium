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
	GitStatus       string
	GitRemote       string
	GitBranch       string
	GitRemoteBranch string
	ExitCode        int
	OS              string
	Shell           string
}

func getApplicationState(shell string, exitCode int) ApplicationState {
	state := ApplicationState{}
	state.GitRemote = getGitRemote()
	state.GitRemoteBranch = getGitRemoteBranch(state.GitRemote)
	state.GitBranch = getGitBranch()
	state.GitStatus = getGitStatus(state.GitRemoteBranch)
	state.ExitCode = getExitCode(exitCode)
	state.OS = getOS()
	state.Shell = getShell(shell)
	return state
}

func getGitStatus(gitRemoteBranch string) string {
	stdErr := bytes.Buffer{}
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Stderr = &stdErr

	output, _ := cmd.Output()

	status := strings.Replace(string(output), "\n", "", -1)

	if strings.Contains(stdErr.String(), "fatal:") {
		return "no_branch"
	}
	if gitRemoteBranch == "" {
		return "no_remote"
	}

	if status == "" {
		return "clean"
	}

	return "dirty"
}

func getGitBranch() string {
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

func getGitRemote() string {
	output, err := exec.Command("git", "remote").Output()
	if err != nil {
		return ""
	}
	remote := strings.Replace(string(output), "\n", "", -1)
	return remote
}

func getGitRemoteBranch(remote string) string {
	output, err := exec.Command("git", "status", "-sb").Output()
	if err != nil {
		return ""
	}
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
func getOS() string {
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

func getExitCode(exitCode int) int {
	// Get exit code from last command
	if exitCode != 0 {
		return 1
	}
	return 0
}

func getShell(shell string) string {
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
