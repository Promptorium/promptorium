package confmodule

import (
	"fmt"
	"os"
	"promptorium/cmd/modules/confmodule/context"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

/*
 * Config parsing
 */

func ParseConfig(rawTheme RawTheme, rawConfig RawConfig, context *context.ApplicationContext) (Config, error) {
	conf := Config{}
	// Copy context from ConfigLoader to Config
	conf.Context = context

	// Theme initialization and parsing
	conf.Theme = parseTheme(rawTheme)

	// Components initialization and parsing
	conf.Components = parseComponents(rawConfig.Components, conf.Theme, context)
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
	resultTheme.PrimaryColor = parseBaseColor(theme.PrimaryColor, "primary", defaultTheme.PrimaryColor)
	resultTheme.SecondaryColor = parseBaseColor(theme.SecondaryColor, "secondary", defaultTheme.SecondaryColor)
	resultTheme.TertiaryColor = parseBaseColor(theme.TertiaryColor, "tertiary", defaultTheme.TertiaryColor)
	resultTheme.QuaternaryColor = parseBaseColor(theme.QuaternaryColor, "quaternary", defaultTheme.QuaternaryColor)
	resultTheme.SuccessColor = parseBaseColor(theme.SuccessColor, "success", defaultTheme.SuccessColor)
	resultTheme.WarningColor = parseBaseColor(theme.WarningColor, "warning", defaultTheme.WarningColor)
	resultTheme.ErrorColor = parseBaseColor(theme.ErrorColor, "error", defaultTheme.ErrorColor)
	resultTheme.BackgroundColor = parseBaseColor(theme.BackgroundColor, "background", defaultTheme.BackgroundColor)
	resultTheme.ForegroundColor = parseBaseColor(theme.ForegroundColor, "foreground", defaultTheme.ForegroundColor)
	resultTheme.GitStatusColorClean = parseBaseColor(theme.GitStatusColorClean, "git_status_clean", defaultTheme.GitStatusColorClean)
	resultTheme.GitStatusColorDirty = parseBaseColor(theme.GitStatusColorDirty, "git_status_dirty", defaultTheme.GitStatusColorDirty)
	resultTheme.GitStatusColorNoRepository = parseBaseColor(theme.GitStatusColorNoRepository, "git_status_no_repository", defaultTheme.GitStatusColorNoRepository)
	resultTheme.GitStatusColorNoUpstream = parseBaseColor(theme.GitStatusColorNoUpstream, "git_status_no_upstream", defaultTheme.GitStatusColorNoUpstream)
	resultTheme.ExitCodeColorOk = parseBaseColor(theme.ExitCodeColorOk, "exit_code_ok", defaultTheme.ExitCodeColorOk)
	resultTheme.ExitCodeColorError = parseBaseColor(theme.ExitCodeColorError, "exit_code_error", defaultTheme.ExitCodeColorError)
	return resultTheme
}

func parseComponents(components []RawComponent, theme Theme, context *context.ApplicationContext) []Component {
	log.Trace().Msg("Parsing components")

	// Initialize components to an empty slice
	var resultComponents []Component

	for _, component := range components {
		log.Trace().Msgf("Parsing component: %v", component.Name)

		var resultComponent Component
		resultComponent.Name = component.Name

		// Parse the component style first, then pass the result to the content and icon style parsers
		// This is because the icon style depends on the component style

		resultComponent.Style = parseComponentStyle(component.Style, theme, context)
		resultComponent.Content = parseContent(component.Content, theme, context, resultComponent.Style)

		resultComponents = append(resultComponents, resultComponent)
	}
	return resultComponents
}

func parseComponentStyle(componentStyle RawComponentStyle, theme Theme, context *context.ApplicationContext) ComponentStyle {
	resultComponentStyle := ComponentStyle{}

	resultComponentStyle.BackgroundColor = parseColor(componentStyle.BackgroundColor, theme, "component background", theme.PrimaryColor, context)
	resultComponentStyle.ForegroundColor = parseColor(componentStyle.ForegroundColor, theme, "component foreground", theme.ForegroundColor, context)
	resultComponentStyle.StartDivider = componentStyle.StartDivider
	resultComponentStyle.EndDivider = componentStyle.EndDivider
	resultComponentStyle.MarginLeft, resultComponentStyle.MarginRight = parseMargin(componentStyle.Margin)
	resultComponentStyle.PaddingLeft, resultComponentStyle.PaddingRight = parsePadding(componentStyle.Padding)
	resultComponentStyle.Align = componentStyle.Align
	return resultComponentStyle
}

func parseContent(content RawContent, theme Theme, context *context.ApplicationContext, componentStyle ComponentStyle) Content {
	var resultContent Content

	resultContent.Module = content.Module
	resultContent.Icon = content.Icon
	resultContent.IconStyle = parseIconStyle(content.IconStyle, theme, context, componentStyle)

	return resultContent
}

func parseIconStyle(iconStyle RawIconStyle, theme Theme, context *context.ApplicationContext, componentStyle ComponentStyle) IconStyle {
	var resultIconStyle IconStyle
	// By default, the icon style is the same as the component style

	resultIconStyle.BackgroundColor = parseColor(iconStyle.BackgroundColor, theme, "icon background", componentStyle.BackgroundColor, context)
	resultIconStyle.ForegroundColor = parseColor(iconStyle.ForegroundColor, theme, "icon foreground", componentStyle.ForegroundColor, context)
	resultIconStyle.Padding = parseIconPadding(iconStyle.Padding)
	resultIconStyle.Separator = iconStyle.Separator

	resultIconStyle.IconPosition = iconStyle.IconPosition

	// TODO: Handle icon separators

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
func parseColor(rawColor RawColorName, theme Theme, colorName string, defaultColor Color, context *context.ApplicationContext) Color {
	var color Color

	if rawColor == "$default" {
		return defaultColor
	}

	if rawColor == "" {
		return defaultColor
	}

	color, ok := Colors[string(rawColor)]
	if ok {
		return color
	}

	switch rawColor {
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
		color = getGitStatusColor(theme, context)
	case "$exit_code_color":
		color = getExitCodeColor(theme, context)
	default:
		fmt.Fprintln(os.Stderr, "promptorium: Error parsing color", rawColor, ", using default", colorName, `color instead`, "(", defaultColor.Name, ")")
		color = defaultColor
	}
	return color
}

// Parses the color from the raw color name and sets the corresponding values in the Color struct
func parseBaseColor(rawColor RawColorName, colorName string, defaultColor Color) Color {

	if rawColor == "" {
		return defaultColor
	}
	if rawColor == "$default" {
		return defaultColor
	}
	color, ok := Colors[string(rawColor)]
	if !ok {
		fmt.Fprintln(os.Stderr, "promptorium: Error parsing color", rawColor, ", using default", colorName, `color instead`, "(", defaultColor.Name, ")")
		return defaultColor
	}
	return color
}

// Color Functions

func getGitStatusColor(theme Theme, context *context.ApplicationContext) Color {
	gitState := context.GitContext.GetContent()

	if !gitState.IsGitRepo {
		log.Debug().Msg("Setting git status color to no repository")
		return theme.GitStatusColorNoRepository
	}
	if gitState.IsDirty {
		log.Debug().Msg("Setting git status color to dirty")
		return theme.GitStatusColorDirty
	}
	if gitState.LocalBranch == "" {
		log.Debug().Msg("Setting git status color to no repository")
		return theme.GitStatusColorNoRepository
	}
	if gitState.UpstreamBranch == "" {
		log.Debug().Msg("Setting git status color to no upstream branch")
		return theme.GitStatusColorNoUpstream
	}
	log.Debug().Msg("Setting git status color to clean")
	return theme.GitStatusColorClean

}

func getExitCodeColor(theme Theme, context *context.ApplicationContext) Color {
	if context.ExitCode.GetContent() == 0 {
		return theme.ExitCodeColorOk
	}
	return theme.ExitCodeColorError
}
