package config

import (
	"fmt"
	"os"
	"promptorium/internal/pkg/confpkg/context"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

var SPACER_PROMPT_ELEMENT = "---"

/*
 * Config parsing
 */

func ParseConfig(rawConfig RawConfig) (Config, error) {
	conf := Config{}
	conf.Version = rawConfig.Version
	conf.Context = rawConfig.Context

	conf.Modules = loadModules()

	conf.Theme = parseTheme(rawConfig.Theme)
	conf.Components = parseComponents(rawConfig.Components, conf.Theme, rawConfig.Context)
	conf.Prompt = parsePrompt(rawConfig.Prompt, conf.Theme, rawConfig.Context, conf.Components, conf.Modules)
	conf.Options = parseOptions(rawConfig.Options)
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

func parseComponents(components []RawComponent, theme Theme, context *context.ApplicationContext) map[string]Component {
	log.Trace().Msg("Parsing components")

	// Initialize components to an empty slice
	var resultComponents = getDefaultComponents()

	for _, component := range components {
		log.Trace().Msgf("Parsing component: %v", component.Name)

		var resultComponent Component
		// Add the $ symbol to the component name to differentiate user defined components from default components
		resultComponent.Name = "$" + component.Name

		// Parse the component style first, then pass the result to the content and icon style parsers
		// This is because the icon style depends on the component style

		resultComponent.Style = parseComponentStyle(component.Style, theme, context)
		resultComponent.Content = component.Content
		resultComponent.Icon = string(component.Style.Icon)
		resultComponent.Type = parseComponentType(component.Type, theme, context)

		// Return an error if a component with the same name already exists
		if _, ok := resultComponents[resultComponent.Name]; ok {
			fmt.Println("promptorium: error: component " + component.Name + " already exists")
		}

		resultComponents[resultComponent.Name] = resultComponent
	}

	return resultComponents
}

func getDefaultComponents() map[string]Component {

	return map[string]Component{
		"user": {
			Name:    "user",
			Type:    "module",
			Content: "user",
			Style: ComponentStyle{
				MarginLeft: 1,
			},
		},
		"cwd": {
			Name:    "cwd",
			Type:    "module",
			Content: "cwd",
			Icon:    "@",
			Style: ComponentStyle{
				ForegroundColor: Colors["white"],
				MarginLeft:      1,
				IconPadding:     1,
				IconPosition:    "left",
			},
		},
		"git_status": {
			Name:    "git_status",
			Type:    "module",
			Content: "git_status",
			Style: ComponentStyle{
				MarginLeft: 1,
			},
		},
		"git_branch": {
			Name:    "git_branch",
			Type:    "module",
			Content: "git_branch",
			Style: ComponentStyle{
				MarginLeft: 1,
			},
		},
		"os_icon": {
			Name:    "os_icon",
			Type:    "module",
			Content: "os_icon",
			Style: ComponentStyle{
				MarginLeft: 1,
			},
		},
		"time": {
			Name:    "time",
			Type:    "module",
			Content: "time",
			Style: ComponentStyle{
				MarginLeft: 1,
			},
		},
		"hostname": {
			Name:    "hostname",
			Type:    "module",
			Content: "hostname",
			Style: ComponentStyle{
				MarginLeft: 1,
			},
		},
		SPACER_PROMPT_ELEMENT: {
			Name:    "spacer",
			Type:    "spacer",
			Content: "",
			Style: ComponentStyle{
				MarginLeft: 1,
			},
		},
	}
}

// / Parses and validates the prompt array. "$" at the beginning of the prompt elements are here to indicate that the prompt element is a component
func parsePrompt(prompt [][]string, theme Theme, context *context.ApplicationContext, components map[string]Component, modules map[string]ModuleEntry) [][]string {
	log.Trace().Msgf("Parsing prompt: %v", prompt)
	resultPrompt := [][]string{}
	for _, promptLine := range prompt {
		resultPromptLine := []string{}
		for _, promptElement := range promptLine {

			trimmedPromptElement := strings.Trim(promptElement, " ")

			_, ok := components[trimmedPromptElement]
			if ok {
				resultPromptLine = append(resultPromptLine, trimmedPromptElement)
			} else {
				fmt.Fprintf(os.Stderr, "promptorium: component %v not found \n", trimmedPromptElement)
			}

		}

		resultPrompt = append(resultPrompt, resultPromptLine)

	}

	return resultPrompt
}

func parseOptions(options RawOptions) ConfigOptions {
	// TODO: Improve this
	log.Trace().Msgf("Parsing options: %v", options)
	resultOptions := ConfigOptions{}

	resultOptions.CWD.HighlightGitRoot = options.CWD.HighlightGitRoot

	return resultOptions
}

func parseComponentStyle(componentStyle RawComponentStyle, theme Theme, context *context.ApplicationContext) ComponentStyle {
	resultComponentStyle := ComponentStyle{}

	resultComponentStyle.BackgroundColor = parseColor(componentStyle.BackgroundColor, theme, "component background", theme.PrimaryColor, context)
	resultComponentStyle.ForegroundColor = parseColor(componentStyle.ForegroundColor, theme, "component foreground", theme.ForegroundColor, context)
	resultComponentStyle.MarginLeft, resultComponentStyle.MarginRight = parseMargin(componentStyle.Margin)
	resultComponentStyle.PaddingLeft, resultComponentStyle.PaddingRight = parsePadding(componentStyle.Padding)
	resultComponentStyle.IconPosition = IconPosition(componentStyle.IconPosition)
	resultComponentStyle.IconPadding = parseIconPadding(componentStyle.IconPadding)
	resultComponentStyle.IconSeparator = componentStyle.IconSeparator
	resultComponentStyle.IconForegroundColor = parseColor(componentStyle.IconForegroundColor, theme, "icon foreground", resultComponentStyle.ForegroundColor, context)
	resultComponentStyle.IconBackgroundColor = parseColor(componentStyle.IconBackgroundColor, theme, "icon background", resultComponentStyle.BackgroundColor, context)
	resultComponentStyle.StartDivider = parseStartDivider(componentStyle.StartDivider, resultComponentStyle, theme)
	resultComponentStyle.EndDivider = parseEndDivider(componentStyle.EndDivider, resultComponentStyle, theme)
	return resultComponentStyle
}

func parseComponentType(componentType RawComponentType, theme Theme, context *context.ApplicationContext) ComponentType {
	if componentType == "" {
		return ComponentType("text")
	}

	if _, ok := ComponentTypes[strings.ToLower(string(componentType))]; !ok {
		log.Warn().Msgf("Unknown component type: %v", componentType)
		return ComponentType("text")
	}

	return ComponentType(componentType)
}

func parseStartDivider(divider string, componentStyle ComponentStyle, theme Theme) string {
	if componentStyle.BackgroundColor.Name == "transparent" {
		return ""
	}
	if divider == "$default" {
		return theme.ComponentStartDivider
	}
	return divider
}

func parseEndDivider(divider string, componentStyle ComponentStyle, theme Theme) string {
	if componentStyle.BackgroundColor.Name == "transparent" {
		return ""
	}
	if divider == "$default" {
		return theme.ComponentEndDivider
	}
	return divider
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
		log.Trace().Msg("Setting git status color to no repository")
		return theme.GitStatusColorNoRepository
	}
	if !gitState.HasUpstream {
		log.Trace().Msg("Setting git status color to no upstream branch")
		return theme.GitStatusColorNoUpstream
	}
	if gitState.IsDirty {
		log.Trace().Msg("Setting git status color to dirty")
		return theme.GitStatusColorDirty
	}
	if gitState.LocalBranch == "" {
		log.Trace().Msg("Setting git status color to no repository")
		return theme.GitStatusColorNoRepository
	}

	log.Trace().Msg("Setting git status color to clean")
	return theme.GitStatusColorClean

}

func getExitCodeColor(theme Theme, context *context.ApplicationContext) Color {
	if context.ExitCode.GetContent() == 0 {
		return theme.ExitCodeColorOk
	}
	return theme.ExitCodeColorError
}
