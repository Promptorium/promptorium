package promptmodule

import (
	"fmt"
	"os"
	"os/exec"
	"promptorium/modules/confmodule"
	"promptorium/modules/utils"
	"strings"
	"time"
	"unicode/utf8"
)

func (b *ComponentBuilder) buildCwdModuleContent() *ComponentBuilder {
	cwd, err := os.Getwd()
	if err != nil {
		return b
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return b
	}

	cwd = strings.ReplaceAll(cwd, homeDir, "~")
	b.ComponentStr += cwd
	b.ComponentLen += utf8.RuneCountInString(cwd)
	return b
}

func (b *ComponentBuilder) buildExitStatusModuleContent() *ComponentBuilder {
	b.ComponentStr += "[" + string(b.Component.Content.Module) + "]"
	b.ComponentLen += utf8.RuneCountInString(string(b.Component.Content.Module)) + 4
	return b
}

func (b *ComponentBuilder) buildGitBranchModuleContent() *ComponentBuilder {

	gitBranch := b.Config.State.GitBranch

	b.ComponentLen += utf8.RuneCountInString(string(gitBranch.GetContent()))

	b.ComponentStr += gitBranch.GetContent()

	return b

}

func (b *ComponentBuilder) buildHostnameModuleContent() *ComponentBuilder {
	hostname, err := os.Hostname()
	if err != nil {
		return b
	}
	b.ComponentLen += utf8.RuneCountInString(hostname)
	b.ComponentStr += hostname
	return b
}

func (b *ComponentBuilder) buildOsIconModuleContent() *ComponentBuilder {
	// TODO: Implement os icon module
	b.ComponentStr += getOSIcon(b.Config)
	b.ComponentLen += 1
	return b
}

func getOSIcon(config confmodule.Config) string {

	switch strings.ToLower(config.State.OS.GetContent()) {
	case "linux":
		return "󰌽"
	case "macos":
		return ""
	case "fedora":
		return ""
	case "ubuntu":
		return ""
	case "arch":
		return "󰣇"
	case "debian":
		return ""
	default:
		return ""
	}
}

func (b *ComponentBuilder) buildTimeModuleContent() *ComponentBuilder {
	time := time.Now().Format("15:04:05")
	b.ComponentLen += utf8.RuneCountInString(time)
	b.ComponentStr += time
	return b
}

func (b *ComponentBuilder) buildUserModuleContent() *ComponentBuilder {
	user := os.Getenv("USER")
	b.ComponentLen += utf8.RuneCountInString(user)
	b.ComponentStr += user
	return b
}

func (b *ComponentBuilder) buildGitStatusModuleContent() *ComponentBuilder {
	gitStatus := getGitStatus()
	gitRemoteBranch := b.Config.State.GitRemoteBranch
	var gitStatusString string
	isAheadIndicator := ""
	isBehindIndicator := ""
	stagingAreaStatus := ""
	numberOfChanges := ""

	if b.Config.State.GitBranch.GetContent() == "" {
		return b
	}

	// Set staging area status
	if gitStatus.HasUnstagedChanges {
		stagingAreaStatus += utils.Colorize("", b.Component.Style.ForegroundColor, b.Component.Style.BackgroundColor, false, b.Config.State.Shell.GetContent())
		b.ComponentLen += 1
	} else if gitStatus.HasStagedChanges {
		stagingAreaStatus += utils.Colorize("", b.Component.Style.ForegroundColor, b.Component.Style.BackgroundColor, false, b.Config.State.Shell.GetContent())
		b.ComponentLen += 1
	} else {
		stagingAreaStatus += utils.Colorize("", b.Config.Theme.SuccessColor, b.Component.Style.BackgroundColor, false, b.Config.State.Shell.GetContent())
		b.ComponentLen += 1
	}

	// Set ahead/behind indicators
	if gitStatus.IsBehind {
		isBehindIndicator += utils.Colorize(" ", b.Config.Theme.ErrorColor, b.Component.Style.BackgroundColor, false, b.Config.State.Shell.GetContent())
		isBehindIndicator += utils.Colorize("↓", b.Config.Theme.ErrorColor, b.Component.Style.BackgroundColor, false, b.Config.State.Shell.GetContent())
		b.ComponentLen += 2
	}
	if gitStatus.IsAhead || gitRemoteBranch.GetContent() == "" {
		isAheadIndicator += utils.Colorize(" ", b.Config.Theme.ErrorColor, b.Component.Style.BackgroundColor, false, b.Config.State.Shell.GetContent())
		isAheadIndicator += utils.Colorize("↑", b.Config.Theme.GitStatusColorDirty, b.Component.Style.BackgroundColor, false, b.Config.State.Shell.GetContent())
		b.ComponentLen += 2
	}

	// Set number of changes
	if gitStatus.NumUnstagedChanges > 0 {
		numberOfChanges = utils.Colorize(fmt.Sprintf("%d", gitStatus.NumUnstagedChanges), b.Config.Theme.ForegroundColor, b.Config.Theme.GitStatusColorDirty, false, b.Config.State.Shell.GetContent())
		b.ComponentLen += utf8.RuneCountInString(fmt.Sprintf("%d", gitStatus.NumUnstagedChanges))
	}
	// Assemble git status string

	gitStatusString = stagingAreaStatus + numberOfChanges + isAheadIndicator + isBehindIndicator
	b.ComponentStr += gitStatusString
	return b
}

func getGitStatus() GitStatus {
	gitStatus := GitStatus{}
	output, err := exec.Command("git", "status").Output()
	if err != nil {
		return gitStatus
	}
	status := string(output)
	if status == "" {
		return gitStatus
	}
	for _, line := range strings.Split(status, "\n") {
		if strings.HasPrefix(line, "Your branch is ahead of") {
			gitStatus.IsAhead = true
			continue
		}
		if strings.HasPrefix(line, "Your branch is behind") {
			gitStatus.IsBehind = true
			continue
		}
		if strings.HasPrefix(line, "Changes not staged for commit:") {
			gitStatus.HasUnstagedChanges = true
			continue
		}
		if strings.HasPrefix(line, "Changes to be committed:") {
			gitStatus.HasStagedChanges = true
			continue
		}
		if strings.HasPrefix(line, "Untracked files:") {
			gitStatus.HasUnstagedChanges = true
			continue
		}

		if strings.Contains(line, "Your branch and 'origin/") || strings.Contains(line, "have diverged,") {
			gitStatus.IsBehind = true
			gitStatus.IsAhead = true
			continue
		}
		//TODO: Handle more cases

	}

	return gitStatus
}

type GitStatus struct {
	IsBehind           bool
	IsAhead            bool
	HasUnstagedChanges bool
	HasStagedChanges   bool
	NumUnstagedChanges int
	NumStagedChanges   int
}
