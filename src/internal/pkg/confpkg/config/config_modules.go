package config

import (
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/rs/zerolog/log"
)

func loadModules() map[string]ModuleEntry {
	log.Trace().Msgf("Loading modules")
	modules := make(map[string]ModuleEntry)
	// Load modules
	modules["git_branch"] = ModuleEntry{Get: getGitBranchModuleContent}
	modules["hostname"] = ModuleEntry{Get: getHostnameModuleContent}
	modules["time"] = ModuleEntry{Get: getTimeModuleContent}
	modules["cwd"] = ModuleEntry{Get: getCwdModuleContent}
	modules["user"] = ModuleEntry{Get: getUserModuleContent}
	modules["os_icon"] = ModuleEntry{Get: getOsIconModuleContent}
	modules["git_status"] = ModuleEntry{Get: getGitStatusModuleContent}
	modules["exit_status"] = ModuleEntry{Get: getExitStatusModuleContent}
	modules["git_upstream"] = ModuleEntry{Get: getGitUpstreamModuleContent}
	modules["git_remote"] = ModuleEntry{Get: getGitRemoteModuleContent}
	return modules
}

/*
 * Default Modules
 */

func getGitBranchModuleContent(config *Config, component *Component) []ComponentContent {
	result := []ComponentContent{}
	if !config.Context.GitContext.GetContent().IsGitRepo {
		return result
	}

	localBranch := config.Context.GitContext.GetContent().LocalBranch
	len := utf8.RuneCountInString(localBranch)

	result = append(result, NewComponentContent(component, localBranch, len))

	return result
}

func getHostnameModuleContent(config *Config, component *Component) []ComponentContent {
	result := []ComponentContent{}
	hostname, err := os.Hostname()
	if err != nil {
		return result
	}
	len := utf8.RuneCountInString(hostname)
	result = append(result, NewComponentContent(component, hostname, len))
	return result
}

func getOsIconModuleContent(config *Config, component *Component) []ComponentContent {
	result := []ComponentContent{}
	icon := GetOSIcon(config)
	len := 1
	result = append(result, NewComponentContent(component, icon, len))
	return result
}

func getTimeModuleContent(config *Config, component *Component) []ComponentContent {
	result := []ComponentContent{}
	time := time.Now().Format("15:04:05")
	len := utf8.RuneCountInString(time)
	result = append(result, NewComponentContent(component, time, len))
	return result
}

func getUserModuleContent(config *Config, component *Component) []ComponentContent {
	result := []ComponentContent{}
	user := os.Getenv("USER")
	len := utf8.RuneCountInString(user)
	result = append(result, NewComponentContent(component, user, len))
	return result
}

func getGitStatusModuleContent(config *Config, component *Component) []ComponentContent {
	result := []ComponentContent{}
	if !config.Context.GitContext.GetContent().IsGitRepo {
		return result
	}
	var componentLen int
	gitState := config.Context.GitContext.GetContent()
	if !gitState.IsGitRepo {
		return result
	}

	gitUpstreamBranch := gitState.UpstreamBranch
	var gitStatusString string
	isAheadIndicator := ""
	isBehindIndicator := ""
	stagingAreaStatus := ""
	numberOfChanges := ""

	if gitState.LocalBranch == "" {
		return result
	}

	// Set staging area status
	if gitState.UnstagedChanges > 0 || gitState.UntrackedFiles > 0 {
		stagingAreaStatus += config.ColorizeString("", component.Style.ForegroundColor, component.Style.BackgroundColor)
		componentLen += 1
	} else if gitState.StagedChanges > 0 {
		stagingAreaStatus += config.ColorizeString("", component.Style.ForegroundColor, component.Style.BackgroundColor)
		componentLen += 1
	} else {
		stagingAreaStatus += config.ColorizeString("", config.Theme.SuccessColor, component.Style.BackgroundColor)
		componentLen += 1
	}

	// Set ahead/behind indicators
	if gitState.Behind > 0 {
		isBehindIndicator += config.ColorizeString(" ", config.Theme.ErrorColor, component.Style.BackgroundColor)
		isBehindIndicator += config.ColorizeString("↓", config.Theme.ErrorColor, component.Style.BackgroundColor)
		componentLen += 2
	}
	if gitState.Ahead > 0 || gitUpstreamBranch == "" {
		isAheadIndicator += config.ColorizeString(" ", config.Theme.ErrorColor, component.Style.BackgroundColor)
		isAheadIndicator += config.ColorizeString("↑", config.Theme.GitStatusColorDirty, component.Style.BackgroundColor)
		componentLen += 2
	}

	// Assemble git status string
	gitStatusString = stagingAreaStatus + numberOfChanges + isAheadIndicator + isBehindIndicator

	result = append(result, NewComponentContent(component, gitStatusString, componentLen))
	return result
}

func getCwdModuleContent(config *Config, component *Component) []ComponentContent {
	result := []ComponentContent{}
	cwd := config.Context.CWD.GetContent()
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return result
	}

	// Replace home directory with "~"
	cwd = strings.ReplaceAll(cwd, homeDir, "~")
	stringCwd := cwd
	cwdLen := utf8.RuneCountInString(cwd)
	componentContent := NewComponentContent(component, stringCwd, cwdLen)
	result = append(result, componentContent)

	if config.Options.CWD.HighlightGitRoot && config.Context.GitContext.GetContent().IsGitRepo {
		result = []ComponentContent{}
		gitRoot := strings.Split(strings.ReplaceAll(config.Context.GitContext.GetContent().GitRoot(), homeDir, "~"), "/")

		for i, part := range strings.Split(cwd, "/") {
			partLen := utf8.RuneCountInString(part)
			partStr := part
			if i > 0 {
				result = append(result, NewComponentContent(component, "/", 1))
			}
			resultPart := NewComponentContent(component, partStr, partLen)
			if i == len(gitRoot)-1 {
				resultPart.Bold = true
				resultPart.Underline = true
			}
			result = append(result, resultPart)
		}
	}

	return result
}

func getExitStatusModuleContent(config *Config, component *Component) []ComponentContent {
	result := []ComponentContent{}
	// if exit code is 0, return a checkmark

	status := strconv.Itoa(config.Context.ExitCode.GetContent())
	if config.Context.ExitCode.GetContent() == 0 {
		status = "✓"
	}

	len := utf8.RuneCountInString(status)

	switch len {
	case 1:
		status = " " + status + " "
		len = 3
	case 2:
		status = " " + status
		len = 3
	}

	result = append(result, NewComponentContent(component, status, len))
	return result
}

func getGitUpstreamModuleContent(config *Config, component *Component) []ComponentContent {
	result := []ComponentContent{}
	gitContext := config.Context.GitContext.GetContent()
	if !gitContext.IsGitRepo || !gitContext.HasUpstream {
		return result
	}
	upstream := gitContext.UpstreamBranch
	len := utf8.RuneCountInString(upstream)
	result = append(result, NewComponentContent(component, upstream, len))
	return result
}

func getGitRemoteModuleContent(config *Config, component *Component) []ComponentContent {
	result := []ComponentContent{}
	gitContext := config.Context.GitContext.GetContent()
	if !gitContext.IsGitRepo || !gitContext.HasUpstream {
		return result
	}
	remote := gitContext.Remote
	len := utf8.RuneCountInString(remote)
	result = append(result, NewComponentContent(component, remote, len))
	return result
}

type ComponentContent struct {
	Len             int
	Str             string
	BackgroundColor Color
	ForegroundColor Color
	Underline       bool
	Bold            bool
}

func NewComponentContent(component *Component, str string, len int) ComponentContent {
	bgcolor := component.Style.BackgroundColor
	fgcolor := component.Style.ForegroundColor

	return ComponentContent{
		Str:             str,
		Len:             len,
		BackgroundColor: bgcolor,
		ForegroundColor: fgcolor,
		Underline:       false,
		Bold:            false,
	}
}

func (c *ComponentContent) Render(config *Config) string {
	foregroundColor := c.ForegroundColor
	backgroundColor := c.BackgroundColor
	if foregroundColor == (Color{}) {
		foregroundColor = config.Theme.PrimaryColor
	}
	if backgroundColor == (Color{}) {
		backgroundColor = config.Theme.BackgroundColor
	}
	underline := c.Underline
	bold := c.Bold

	return addColor(c.Str, foregroundColor.ForegroundCode, backgroundColor.BackgroudCode, bold, underline, config.Context.Shell.GetContent())
}
