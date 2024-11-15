package context

import (
	"path/filepath"
	"promptorium/cmd/modules/confmodule/context/gitcontext"
	"promptorium/cmd/modules/confmodule/context/oscontext"

	"github.com/rs/zerolog/log"
)

type ApplicationContext struct {
	ExitCode   CachedData[int]
	CWD        CachedData[string]
	GitContext CachedData[gitcontext.GitContext]
	OS         CachedData[oscontext.OS]
	Shell      CachedData[Shell]
}

type Shell int

const (
	ShellBash Shell = iota
	ShellZsh
	ShellOther
)

func GetApplicationContext(shell string, exitCode int) *ApplicationContext {
	context := ApplicationContext{}
	context.GitContext = NewCachedData(gitcontext.GetGitState, "git repo")
	context.ExitCode = NewCachedData(func() int { return context.getExitCode(exitCode) }, "exit code")
	context.OS = NewCachedData(oscontext.GetOS, "os")
	context.Shell = NewCachedData(func() Shell { return context.getShell(shell) }, "shell")
	return &context
}

func (context *ApplicationContext) getExitCode(exitCode int) int {
	// Get exit code from last command
	if exitCode != 0 {
		return 1
	}
	return 0
}

func (context *ApplicationContext) getShell(shell string) Shell {
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

// Cached State

type CachedData[T any] struct {
	content  T
	name     string
	isCached bool
	refresh  func() T
}

func NewCachedData[T any](refresh func() T, name string) CachedData[T] {
	return CachedData[T]{
		name:     name,
		isCached: false,
		refresh:  refresh,
	}
}

func (c *CachedData[T]) GetContent() T {
	if !c.isCached {
		log.Debug().Msgf("[CACHED_DATA@context] Refreshing cached data %s", c.name)
		c.content = c.refresh()
		c.isCached = true
	}
	return c.content
}
