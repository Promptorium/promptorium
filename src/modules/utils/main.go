package utils

import (
	"promptorium-go/modules/confmodule"
	"strings"
	"unicode/utf8"
)

func GetBottomDecoration(config confmodule.Config) (string, int) {
	//TODO: Implement multiple styles

	decorationString := "  "
	decorationLen := 3
	decorationString = Colorize(decorationString, config.Theme.SecondaryColor, config.Theme.BackgroundColor, false, config.State.Shell)
	return decorationString, decorationLen
}

func GetSpacer(config confmodule.Config, promptLen int, terminalWidth int) string {
	// Check if prompt is wider than terminal
	if promptLen > terminalWidth {
		return ""
	}

	// Check if prompt  + spacerChar is wider than terminal
	if promptLen+utf8.RuneCountInString(config.Theme.Spacer) > terminalWidth {
		return ""
	}

	// Calculate margin
	margin := terminalWidth - promptLen
	spacerChar := config.Theme.Spacer
	if spacerChar == "" {
		spacerChar = " "
	}

	spacer := strings.Repeat(spacerChar, margin)
	return Colorize(spacer, config.Theme.ForegroundColor, config.Theme.BackgroundColor, false, config.State.Shell)
}

/*
 * Colorize
 */

// Colorize takes a string and applies the foreground and background colors to it.

// If the reverse parameter is true, the foreground and background colors are swapped.
func Colorize(text string, fgcolor confmodule.Color, bgcolor confmodule.Color, reverse bool, shell string) string {

	bgcode := bgcolor.BackgroudCode
	fgcode := fgcolor.ForegroundCode

	if reverse {
		bgcode = fgcolor.BackgroudCode
		fgcode = bgcolor.ForegroundCode
	}

	return addColor(text, fgcode, bgcode, false, false, shell)
}

func addColor(text string, fgcode string, bgcode string, bold bool, underline bool, shell string) string {

	var resultString string

	// ANSI codes
	ansiReset := "\x1b[0m"
	ansiBold := "\x1b[1m"
	ansiUnderline := "\x1b[4m"
	// Escape codes for different shells
	escapeCodeStart := ""
	escapeCodeEnd := ""

	if shell == "bash" {
		escapeCodeStart = "\\["
		escapeCodeEnd = "\\]"
	}
	if shell == "zsh" {
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
