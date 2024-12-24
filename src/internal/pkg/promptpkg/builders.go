package promptpkg

import (
	"promptorium/internal/pkg/confpkg/config"
	"strings"
	"unicode/utf8"

	"github.com/rs/zerolog/log"
)

var SPACER_PROMPT_ELEMENT = "---"

/*
 * Pre-render components
 * These are obtained by calling the Build() method of the respective builder struct.
 * They are organized in a hierarchical order (Prompt -> PromptLine -> PromptComponent -> ComponentContent)
 * All of them must provide a Render() method, which gives the final prompt string
 */

type Prompt struct {
	Config      config.Config
	PromptLines []PromptLine
}

type PromptLine struct {
	PromptComponents []PromptComponent
	Config           config.Config
}
type PromptComponent struct {
	Style     config.ComponentStyle
	Content   []config.ComponentContent
	Len       int
	Config    config.Config
	IsSpacer  bool
	Component config.Component
}

func (p Prompt) Render() string {
	result := ""

	for _, line := range p.PromptLines {
		result += line.Render() + "\n"
	}
	return result
}

func (p PromptLine) Render() string {
	result := ""

	leftPartComponents := []PromptComponent{}
	rightPartComponents := []PromptComponent{}

	leftPart := ""
	leftPartLen := 0
	rightPart := ""
	rightPartLen := 0

	spacer := ""
	spacerLen := 0

	// Check if the line has a spacer component
	foundSpacer := false
	for _, component := range p.PromptComponents {
		if component.IsSpacer {
			foundSpacer = true
			continue
		}
		if foundSpacer {
			rightPartComponents = append(rightPartComponents, component)
			rightPartLen += component.Len
		} else {
			leftPartComponents = append(leftPartComponents, component)
			leftPartLen += component.Len
		}
	}

	for _, component := range leftPartComponents {
		leftPart += component.Render()
	}

	for _, component := range rightPartComponents {
		rightPart += component.Render()
	}

	// Calculate spacer length
	if foundSpacer {
		spacerLen = p.Config.Context.TerminalWidth.GetContent() - leftPartLen - rightPartLen
		spacerChar := p.Config.Theme.Spacer
		if spacerChar == "" {
			spacerChar = " "
		}
		if spacerLen > 0 {
			spacer = strings.Repeat(spacerChar, spacerLen)

		}
	}

	result = leftPart
	if foundSpacer {
		result += spacer
	}
	result += rightPart
	return result
}

func (p PromptComponent) Render() string {
	var result = ""

	for _, content := range p.Content {
		componentContent := content.Render(&p.Config)
		result += componentContent
	}
	return result
}

/*
 * Builders
 * These are used to build the pre-render components
 */

type PromptBuilder struct {
	Config config.Config
}

type PromptLineBuilder struct {
	Config config.Config
	Line   []string
}

type PromptComponentBuilder struct {
	Component config.Component
	Config    config.Config
}

type ComponentContentBuilder struct {
	bgcolor   config.Color
	fgcolor   config.Color
	underline bool
	bold      bool
}

func NewPromptBuilder(config config.Config) PromptBuilder {
	return PromptBuilder{Config: config}
}

func (b PromptBuilder) BuildPrompt() Prompt {
	var promptLines []PromptLine
	for _, line := range b.Config.Prompt {
		promptLines = append(promptLines, b.NewPromptLineBuilder(line).BuildPromptLine())
	}
	return Prompt{Config: b.Config, PromptLines: promptLines}
}

func (b PromptBuilder) NewPromptLineBuilder(line []string) PromptLineBuilder {
	return PromptLineBuilder{
		Config: b.Config,
		Line:   line,
	}
}
func (b PromptLineBuilder) BuildPromptLine() PromptLine {
	result := []PromptComponent{}

	for _, component := range b.Line {
		result = append(result, PromptComponent(b.NewPromptComponentBuilder(component).BuildPromptComponent()))
	}

	return PromptLine{
		PromptComponents: result,
		Config:           b.Config,
	}

}
func (b PromptLineBuilder) NewPromptComponentBuilder(componentName string) PromptComponentBuilder {
	// Get the ComponentContent from the module with the given name
	component, ok := b.Config.Components[componentName]
	if !ok {
		log.Error().Msgf("Component %s not found", componentName)
		return PromptComponentBuilder{}
	}

	return PromptComponentBuilder{
		Component: component,
		Config:    b.Config,
	}
}

func (b PromptComponentBuilder) BuildPromptComponent() PromptComponent {
	componentContent := []config.ComponentContent{}
	result := PromptComponent{}
	contentLen := 0

	switch b.Component.Type {
	case "module":
		module, ok := b.Config.Modules[b.Component.Content]
		if !ok {
			log.Error().Msgf("Module %s not found", b.Component.Content)
			return PromptComponent{}
		}
		componentContent = addDecorationsContent(module.Get(&b.Config, &b.Component), b.Component, b.Component.Style, b.Config)

	case "text":
		componentContent = addDecorationsContent([]config.ComponentContent{config.NewComponentContent(&b.Component, b.Component.Content, utf8.RuneCountInString(b.Component.Content))}, b.Component, b.Component.Style, b.Config)
	case "spacer":
		componentContent = []config.ComponentContent{
			config.NewComponentContent(&b.Component, b.Component.Content, utf8.RuneCountInString(b.Component.Content)),
		}
		result.Len = 0
		result.Content = []config.ComponentContent{}
		result.IsSpacer = true
	}

	for _, content := range componentContent {
		contentLen += content.Len
	}

	result.Style = b.Component.Style
	result.Content = componentContent
	result.Config = b.Config
	result.Len = contentLen
	result.Component = b.Component

	return result
}

func addDecorationsContent(componentContent []config.ComponentContent, component config.Component, style config.ComponentStyle, conf config.Config) []config.ComponentContent {
	result := componentContent

	if len(componentContent) == 0 {
		return result
	}

	// The order in which the decorations are added is important
	// Icon -> Padding -> Dividers -> Margins

	// Icon
	if component.Icon != "" {
		var iconPadding = conf.ColorizeString(strings.Repeat(" ", style.IconPadding), style.IconForegroundColor, style.IconBackgroundColor)
		var icon = conf.ColorizeString(component.Icon, style.IconForegroundColor, style.BackgroundColor)
		if style.IconPosition == "left" {
			result = append([]config.ComponentContent{config.NewComponentContent(&component, icon+iconPadding, 1+style.IconPadding)}, result...)
		} else {
			result = append(result, config.NewComponentContent(&component, icon+iconPadding, 1+style.IconPadding))
		}
	}

	// Padding
	if style.PaddingLeft > 0 {
		result = append([]config.ComponentContent{config.NewComponentContent(&component, strings.Repeat(" ", style.PaddingLeft), style.PaddingLeft)}, result...)
	}
	if style.PaddingRight > 0 {
		result = append(result, config.NewComponentContent(&component, strings.Repeat(" ", style.PaddingRight), style.PaddingRight))
	}

	// Dividers
	if style.BackgroundColor.Name != "transparent" {
		if style.StartDivider != "" {
			var startDivider = conf.ColorizeString(style.StartDivider, style.BackgroundColor, conf.Theme.BackgroundColor)
			result = append([]config.ComponentContent{config.NewComponentContent(&component, startDivider, 1)}, result...)
		}
		if style.EndDivider != "" {
			var endDivider = conf.ColorizeString(style.EndDivider, style.BackgroundColor, conf.Theme.BackgroundColor)
			result = append(result, config.NewComponentContent(&component, endDivider, 1))
		}
	}

	// Margin
	if style.MarginLeft > 0 {
		var marginLeft = conf.ColorizeString(strings.Repeat(" ", style.MarginLeft), style.BackgroundColor, conf.Theme.BackgroundColor)
		result = append([]config.ComponentContent{config.NewComponentContent(&component, strings.Repeat(marginLeft, style.MarginLeft), style.MarginLeft)}, result...)
	}
	if style.MarginRight > 0 {
		var marginRight = conf.ColorizeString(strings.Repeat(" ", style.MarginRight), style.BackgroundColor, conf.Theme.BackgroundColor)
		result = append(result, config.NewComponentContent(&component, strings.Repeat(marginRight, style.MarginRight), style.MarginRight))
	}

	return result
}
