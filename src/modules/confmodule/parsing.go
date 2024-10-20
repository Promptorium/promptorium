package confmodule

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

/*
 * Config parsing
 */

func parseConfig(rawTheme RawTheme, rawConfig RawConfig, state ApplicationState) (Config, error) {
	conf := Config{}
	log.Debug().Msgf("Parsing config")
	log.Debug().Msgf("Raw config: %v", rawConfig)
	log.Debug().Msgf("Raw theme: %v", rawTheme)

	// Copy state from ConfigLoader to Config
	conf.State = state

	// Theme initialization and parsing
	conf.Theme = parseTheme(rawTheme)

	// Components initialization and parsing
	conf.Components = parseComponents(rawConfig.Components, conf.Theme, state)
	return conf, nil
}

// Parses the theme from the raw theme
func parseTheme(theme RawTheme) Theme {
	var resultTheme Theme
	defaultTheme := getDefaultTheme()

	resultTheme.ComponentStartDivider = theme.ComponentStartDivider
	resultTheme.ComponentEndDivider = theme.ComponentEndDivider
	resultTheme.Spacer = theme.Spacer

	// Colors
	resultTheme.PrimaryColor = parseThemeColor(theme.PrimaryColor, "primary", defaultTheme.PrimaryColor)
	resultTheme.SecondaryColor = parseThemeColor(theme.SecondaryColor, "secondary", defaultTheme.SecondaryColor)
	resultTheme.TertiaryColor = parseThemeColor(theme.TertiaryColor, "tertiary", defaultTheme.TertiaryColor)
	resultTheme.QuaternaryColor = parseThemeColor(theme.QuaternaryColor, "quaternary", defaultTheme.QuaternaryColor)
	resultTheme.SuccessColor = parseThemeColor(theme.SuccessColor, "success", defaultTheme.SuccessColor)
	resultTheme.WarningColor = parseThemeColor(theme.WarningColor, "warning", defaultTheme.WarningColor)
	resultTheme.ErrorColor = parseThemeColor(theme.ErrorColor, "error", defaultTheme.ErrorColor)
	resultTheme.BackgroundColor = parseThemeColor(theme.BackgroundColor, "background", defaultTheme.BackgroundColor)
	resultTheme.ForegroundColor = parseThemeColor(theme.ForegroundColor, "foreground", defaultTheme.ForegroundColor)
	resultTheme.GitStatusColorClean = parseThemeColor(theme.GitStatusColorClean, "git_status_clean", defaultTheme.GitStatusColorClean)
	resultTheme.GitStatusColorDirty = parseThemeColor(theme.GitStatusColorDirty, "git_status_dirty", defaultTheme.GitStatusColorDirty)
	resultTheme.GitStatusColorNoBranch = parseThemeColor(theme.GitStatusColorNoBranch, "git_status_no_branch", defaultTheme.GitStatusColorNoBranch)
	resultTheme.GitStatusColorNoRemote = parseThemeColor(theme.GitStatusColorNoRemote, "git_status_no_remote", defaultTheme.GitStatusColorNoRemote)
	resultTheme.ExitCodeColorOk = parseThemeColor(theme.ExitCodeColorOk, "exit_code_ok", defaultTheme.ExitCodeColorOk)
	resultTheme.ExitCodeColorError = parseThemeColor(theme.ExitCodeColorError, "exit_code_error", defaultTheme.ExitCodeColorError)
	return resultTheme
}

func parseComponents(components []RawComponent, theme Theme, state ApplicationState) []Component {
	log.Debug().Msgf("Parsing components")

	// Initialize components to an empty slice
	var resultComponents []Component

	for _, component := range components {
		log.Debug().Msgf("Parsing component: %v", component)

		var resultComponent Component
		resultComponent.Name = component.Name

		// Parse the component style first, then pass the result to the content and icon style parsers
		// This is because the icon style depends on the component style

		resultComponent.Style = parseComponentStyle(component.Style, theme, state)
		resultComponent.Content = parseContent(component.Content, theme, state, resultComponent.Style)

		resultComponents = append(resultComponents, resultComponent)
	}
	return resultComponents
}

func parseComponentStyle(componentStyle RawComponentStyle, theme Theme, state ApplicationState) ComponentStyle {
	resultComponentStyle := ComponentStyle{}

	resultComponentStyle.BackgroundColor = parseColor(componentStyle.BackgroundColor, theme, "component background", theme.PrimaryColor, state)
	resultComponentStyle.ForegroundColor = parseColor(componentStyle.ForegroundColor, theme, "component foreground", theme.ForegroundColor, state)
	resultComponentStyle.StartDivider = componentStyle.StartDivider
	resultComponentStyle.EndDivider = componentStyle.EndDivider
	resultComponentStyle.MarginLeft, resultComponentStyle.MarginRight = parseMargin(componentStyle.Margin)
	resultComponentStyle.PaddingLeft, resultComponentStyle.PaddingRight = parsePadding(componentStyle.Padding)
	resultComponentStyle.Align = componentStyle.Align
	return resultComponentStyle
}

func parseContent(content RawContent, theme Theme, state ApplicationState, componentStyle ComponentStyle) Content {
	var resultContent Content

	log.Debug().Msgf("Parsing content")
	resultContent.Module = content.Module
	resultContent.Icon = content.Icon
	resultContent.IconStyle = parseIconStyle(content.IconStyle, theme, state, componentStyle)

	return resultContent
}

func parseIconStyle(iconStyle RawIconStyle, theme Theme, state ApplicationState, componentStyle ComponentStyle) IconStyle {
	var resultIconStyle IconStyle
	// By default, the icon style is the same as the component style

	log.Debug().Msgf("Parsing icon style")
	log.Debug().Msgf("Raw Icon style: %v", iconStyle)
	log.Debug().Msgf("Component style: %v", componentStyle)

	resultIconStyle.BackgroundColor = parseColor(iconStyle.BackgroundColor, theme, "icon background", componentStyle.BackgroundColor, state)
	resultIconStyle.ForegroundColor = parseColor(iconStyle.ForegroundColor, theme, "icon foreground", componentStyle.ForegroundColor, state)
	resultIconStyle.Padding = parseIconPadding(iconStyle.Padding)
	resultIconStyle.Separator = iconStyle.Separator

	resultIconStyle.IconPosition = iconStyle.IconPosition

	log.Debug().Msgf("Icon style: %v", resultIconStyle)

	return resultIconStyle
}

func parsePadding(rawPadding string) (int, int) {
	paddingLeft := 0
	paddingRight := 0

	if rawPadding == "" {
		return paddingLeft, paddingRight
	}
	// Parse the string and split it into an array at each whitespace
	// If the array has only one element, it means that both paddings are the same
	// If the array has two elements, the first element is the left padding and the second element is the right padding
	// If the array has more than two elements, it means that the array is invalid and the default padding is used
	paddings := strings.Split(rawPadding, " ")
	if len(paddings) == 1 {
		padding, _ := strconv.Atoi(paddings[0])
		return padding, padding
	}
	if len(paddings) == 2 {
		paddingLeft, _ := strconv.Atoi(paddings[0])
		paddingRight, _ := strconv.Atoi(paddings[1])
		return paddingLeft, paddingRight
	}
	return paddingLeft, paddingRight
}

func parseIconPadding(rawPadding string) int {
	padding := 0
	if rawPadding == "" {
		return padding
	}
	padding, err := strconv.Atoi(rawPadding)
	if err != nil {
		fmt.Fprintln(os.Stderr, "promptorium: Error parsing icon padding", rawPadding, ", using default padding instead", "(", padding, ")")
		return padding
	}
	return padding
}

func parseMargin(rawMargin string) (int, int) {
	marginLeft := 0
	marginRight := 0

	if rawMargin == "" {
		return marginLeft, marginRight
	}
	// Parse the string and split it into an array at each whitespace
	// If the array has only one element, it means that both margins are the same
	// If the array has two elements, the first element is the left margin and the second element is the right margin
	// If the array has more than two elements, it means that the array is invalid and the default margin is used
	margins := strings.Split(rawMargin, " ")
	if len(margins) == 1 {
		margin, _ := strconv.Atoi(margins[0])
		return margin, margin
	}
	if len(margins) == 2 {
		marginLeft, _ := strconv.Atoi(margins[0])
		marginRight, _ := strconv.Atoi(margins[1])
		return marginLeft, marginRight
	}
	return marginLeft, marginRight
}
func parseColor(rawColor RawColorName, theme Theme, colorName string, defaultColor Color, state ApplicationState) Color {
	var color Color

	if rawColor == "$default" {
		return defaultColor
	}

	if rawColor == "" {
		return defaultColor
	}

	switch rawColor {
	case "black":
		color = Colors["black"]
	case "red":
		color = Colors["red"]
	case "green":
		color = Colors["green"]
	case "yellow":
		color = Colors["yellow"]
	case "blue":
		color = Colors["blue"]
	case "magenta":
		color = Colors["magenta"]
	case "cyan":
		color = Colors["cyan"]
	case "white":
		color = Colors["white"]
	case "none":
		color = Colors["transparent"]
	case "transparent":
		color = Colors["transparent"]
	case "$primary_color":
		color = theme.PrimaryColor
	case "$secondary_color":
		color = theme.SecondaryColor
	case "$tertiary_color":
		color = theme.TertiaryColor
	case "$quaternary_color":
		color = theme.QuaternaryColor
	case "$success_color":
		color = theme.SuccessColor
	case "$warning_color":
		color = theme.WarningColor
	case "$error_color":
		color = theme.ErrorColor
	case "$git_status_color":
		color = getGitStatusColor(theme, state)
	case "$exit_code_color":
		color = getExitCodeColor(theme, state)
	default:
		fmt.Fprintln(os.Stderr, "promptorium: Error parsing color", rawColor, ", using default", colorName, `color instead`, "(", defaultColor.Name, ")")
		color = defaultColor
	}
	return color
}

func getGitStatusColor(theme Theme, state ApplicationState) Color {
	if state.GitStatus == "clean" {
		return theme.GitStatusColorClean
	}
	if state.GitStatus == "dirty" {
		return theme.GitStatusColorDirty
	}
	if state.GitStatus == "no_branch" {
		return theme.GitStatusColorNoBranch
	}
	if state.GitStatus == "no_remote" {
		return theme.GitStatusColorNoRemote
	}
	return theme.GitStatusColorClean

}

func getExitCodeColor(theme Theme, state ApplicationState) Color {
	if state.ExitCode == 0 {
		return theme.ExitCodeColorOk
	}
	return theme.ExitCodeColorError
}

// Parses the color from the raw color name and sets the corresponding values in the Color struct
func parseThemeColor(rawColor RawColorName, colorName string, defaultColor Color) Color {
	var resultColor Color

	if rawColor == "" {
		return defaultColor
	}

	if rawColor == "$default" {
		return defaultColor
	}

	switch rawColor {
	case "black":
		resultColor = Colors["black"]
	case "red":
		resultColor = Colors["red"]
	case "green":
		resultColor = Colors["green"]
	case "yellow":
		resultColor = Colors["yellow"]
	case "blue":
		resultColor = Colors["blue"]
	case "magenta":
		resultColor = Colors["magenta"]
	case "cyan":
		resultColor = Colors["cyan"]
	case "white":
		resultColor = Colors["white"]
	case "none":
		resultColor = Colors["none"]
	case "transparent":
		resultColor = Colors["transparent"]
	default:
		fmt.Fprintln(os.Stderr, "promptorium: Error parsing color", rawColor, ", using default", colorName, `color instead`, "(", defaultColor.Name, ")")
		resultColor = defaultColor
	}
	return resultColor
}
