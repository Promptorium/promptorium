package context

import (
	"os"
	"path/filepath"
	"promptorium/internal/pkg/confpkg/context/gitcontext"
	"promptorium/internal/pkg/confpkg/context/oscontext"
	"promptorium/internal/utils"

	"github.com/rs/zerolog/log"
	"golang.org/x/term"
)

type ApplicationContext struct {
	ExitCode      utils.CachedData[int]
	CWD           utils.CachedData[string]
	GitContext    utils.CachedData[gitcontext.GitContext]
	OS            utils.CachedData[oscontext.OS]
	Shell         utils.CachedData[ShellType]
	TerminalWidth utils.CachedData[int]
}

type ShellType int

const (
	ShellBash ShellType = iota
	ShellZsh
	ShellOther
)

func GetApplicationContext(shell string, exitCode int) *ApplicationContext {
	context := ApplicationContext{}

	context.GitContext = utils.NewCachedData(gitcontext.GetGitState, "git repo")

	context.ExitCode = utils.NewCachedData(func(result chan int) { result <- exitCode }, "exit code")

	context.OS = utils.NewCachedData(oscontext.GetOS, "os")
	context.TerminalWidth = utils.NewCachedData(context.getTerminalWidth, "terminal width")

	context.CWD = utils.NewCachedData(context.getCWD, "cwd")

	context.Shell = utils.NewCachedData(func(shellType chan ShellType) { shellType <- context.getShell(shell) }, "shell")

	return &context
}

/*
 * Context cached data getters
 */

func (context *ApplicationContext) getShell(shell string) ShellType {
	shell = filepath.Base(shell)
	switch shell {
	case "bash":
		return ShellBash
	case "zsh":
		return ShellZsh
	default:
		return ShellOther
	}
}

func (context *ApplicationContext) getTerminalWidth(result chan int) {
	terminalWidth, _, error := term.GetSize(0)
	if error != nil {
		log.Warn().Msg("Error getting terminal width")
		terminalWidth = 0
	}
	result <- terminalWidth
}

func (context *ApplicationContext) getCWD(result chan string) {
	cwd, error := os.Getwd()
	if error != nil {
		log.Warn().Msg("Error getting CWD")
		cwd = ""
	}
	result <- cwd
}
