package config

import (
	"promptorium/internal/pkg/confpkg/context"
	"promptorium/internal/pkg/confpkg/context/oscontext"
	"strings"
	"unicode/utf8"
)

func (c *Config) ColorizeString(text string, fgcolor Color, bgcolor Color) string {
	return addColor(text, fgcolor.ForegroundCode, bgcolor.BackgroudCode, false, false, c.Context.Shell.GetContent())
}
func (c *Config) ColorizeStringBold(text string, fgcolor Color, bgcolor Color) string {
	return addColor(text, fgcolor.ForegroundCode, bgcolor.BackgroudCode, true, false, c.Context.Shell.GetContent())
}

func (c *Config) ColorizeStringUnderline(text string, fgcolor Color, bgcolor Color) string {
	return addColor(text, fgcolor.ForegroundCode, bgcolor.BackgroudCode, false, true, c.Context.Shell.GetContent())
}

func (c *Config) GetSpacer(promptLen int, terminalWidth int) string {
	// Check if prompt is wider than terminal
	if promptLen > terminalWidth {
		return ""
	}

	// Check if prompt  + spacerChar is wider than terminal
	if promptLen+utf8.RuneCountInString(c.Theme.Spacer) > terminalWidth {
		return ""
	}

	// Calculate margin
	margin := terminalWidth - promptLen
	spacerChar := c.Theme.Spacer
	if spacerChar == "" {
		spacerChar = " "
	}

	spacer := strings.Repeat(spacerChar, margin)
	return c.ColorizeString(spacer, c.Theme.ForegroundColor, c.Theme.BackgroundColor)
}

func GetOSIcon(config *Config) string {

	switch config.Context.OS.GetContent() {
	case oscontext.OSLinux:
		return "󰌽"
	case oscontext.OSMac:
		return ""
	case oscontext.OSFedora:
		return ""
	case oscontext.OSUbuntu:
		return ""
	case oscontext.OSArch:
		return "󰣇"
	case oscontext.OSDebian:
		return ""
	default:
		return ""
	}
}

func addColor(text string, fgcode string, bgcode string, bold bool, underline bool, shell context.ShellType) string {

	var resultString string

	// ANSI codes
	ansiReset := "\x1b[0m"
	ansiBold := "\x1b[1m"
	ansiUnderline := "\x1b[4m"
	// Escape codes for different shells
	escapeCodeStart := ""
	escapeCodeEnd := ""

	if shell == context.ShellBash {
		escapeCodeStart = "\\["
		escapeCodeEnd = "\\]"
	}
	if shell == context.ShellZsh {
		escapeCodeStart = "%{"
		escapeCodeEnd = "%}"
	}

	if bold {
		resultString += escapeCodeStart + ansiBold + escapeCodeEnd
	}
	if underline {
		resultString += escapeCodeStart + ansiUnderline + escapeCodeEnd
	}

	resultString += escapeCodeStart + "\x1b[" + fgcode + ";" + bgcode + "m" + escapeCodeEnd + text + escapeCodeStart + ansiReset + escapeCodeEnd
	return resultString
}
