package confmodule

import (
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/rs/zerolog/log"
)

// Loads modules into the modules map
func loadModules() map[string]ModuleEntry {
	log.Debug().Msg("Loading modules")
	modules := make(map[string]ModuleEntry)
	// Load modules
	modules["git_branch"] = ModuleEntry{Name: "git_branch", Get: getGitBranchModuleContent}
	modules["hostname"] = ModuleEntry{Name: "hostname", Get: getHostnameModuleContent}
	modules["time"] = ModuleEntry{Name: "time", Get: getTimeModuleContent}
	modules["cwd"] = ModuleEntry{Name: "cwd", Get: getCwdModuleContent}
	modules["user"] = ModuleEntry{Name: "user", Get: getUserModuleContent}
	modules["os_icon"] = ModuleEntry{Name: "os_icon", Get: getOsIconModuleContent}
	modules["git_status"] = ModuleEntry{Name: "git_status", Get: getGitStatusModuleContent}

	return modules
}

/*
 * Default Modules
 */

func getGitBranchModuleContent(config *Config, component *Component) (string, int) {
	localBranch := config.Context.GitContext.GetContent().LocalBranch
	len := utf8.RuneCountInString(localBranch)
	return localBranch, len
}

func getHostnameModuleContent(config *Config, component *Component) (string, int) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", 0
	}
	len := utf8.RuneCountInString(hostname)
	return hostname, len
}

func getOsIconModuleContent(config *Config, component *Component) (string, int) {
	icon := GetOSIcon(config)
	len := 1
	return icon, len
}

func getTimeModuleContent(config *Config, component *Component) (string, int) {
	time := time.Now().Format("15:04:05")
	len := utf8.RuneCountInString(time)
	return time, len
}

func getUserModuleContent(config *Config, component *Component) (string, int) {
	user := os.Getenv("USER")
	len := utf8.RuneCountInString(user)
	return user, len
}

func getGitStatusModuleContent(config *Config, component *Component) (string, int) {
	var componentLen int
	gitState := config.Context.GitContext.GetContent()
	if !gitState.IsGitRepo {
		return "", 0
	}

	gitRemoteBranch := gitState.RemoteBranch
	var gitStatusString string
	isAheadIndicator := ""
	isBehindIndicator := ""
	stagingAreaStatus := ""
	numberOfChanges := ""

	if gitState.LocalBranch == "" {
		return "", 0
	}

	// Set staging area status
	if gitState.UnstagedChanges > 0 {
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
	if gitState.Ahead > 0 || gitRemoteBranch == "" {
		isAheadIndicator += config.ColorizeString(" ", config.Theme.ErrorColor, component.Style.BackgroundColor)
		isAheadIndicator += config.ColorizeString("↑", config.Theme.GitStatusColorDirty, component.Style.BackgroundColor)
		componentLen += 2
	}

	// Assemble git status string
	gitStatusString = stagingAreaStatus + numberOfChanges + isAheadIndicator + isBehindIndicator

	return gitStatusString, componentLen
}

func getCwdModuleContent(config *Config, component *Component) (string, int) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", 0
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", 0
	}

	cwd = strings.ReplaceAll(cwd, homeDir, "~")
	stringCwd := cwd
	cwdLen := utf8.RuneCountInString(cwd)
	return stringCwd, cwdLen
}
