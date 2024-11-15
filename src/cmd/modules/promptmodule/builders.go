package promptmodule

import (
	"fmt"
	"os"
	"promptorium/cmd/modules/confmodule"
	"strings"
	"unicode/utf8"

	"github.com/rs/zerolog/log"
)

type PromptBuilder struct {
	Config        confmodule.Config
	Str           string
	Len           int
	TerminalWidth int
}

type PromptPartBuilder struct {
	Config        confmodule.Config
	Components    []confmodule.Component
	PromptPartStr string
	PromptPartLen int
}

type PromptPart struct {
	Str string
	Len int
}

type ComponentBuilder struct {
	Component    confmodule.Component
	Config       confmodule.Config
	ComponentStr string
	ComponentLen int
}

type Component struct {
	Str string
	Len int
}

func NewPromptBuilder(config confmodule.Config, terminalWidth int) *PromptBuilder {
	return &PromptBuilder{
		Config:        config,
		Str:           "",
		Len:           0,
		TerminalWidth: terminalWidth,
	}
}
func (b *PromptBuilder) NewPromptPartBuilder(components []confmodule.Component) *PromptPartBuilder {
	return &PromptPartBuilder{
		Config:        b.Config,
		Components:    components,
		PromptPartStr: "",
		PromptPartLen: 0,
	}
}

func (b *PromptPartBuilder) NewComponentBuilder(component confmodule.Component) *ComponentBuilder {
	return &ComponentBuilder{
		Component:    component,
		Config:       b.Config,
		ComponentStr: "",
		ComponentLen: 0,
	}

}
func (b *PromptBuilder) BuildPrompt() string {

	leftComponents, rightComponents := b.splitComponents()

	leftPart := b.NewPromptPartBuilder(leftComponents).buildPart()
	rightPart := b.NewPromptPartBuilder(rightComponents).buildPart()

	return b.joinParts(leftPart, rightPart)
}

func (b *PromptBuilder) splitComponents() ([]confmodule.Component, []confmodule.Component) {
	leftComponents := []confmodule.Component{}
	rightComponents := []confmodule.Component{}
	log.Debug().Msgf("Splitting components")

	for _, component := range b.Config.Components {
		if component.Style.Align == "right" {
			rightComponents = append(rightComponents, component)
		} else {
			leftComponents = append(leftComponents, component)
		}
	}
	return leftComponents, rightComponents
}

func (b *PromptBuilder) joinParts(leftPart PromptPart, rightPart PromptPart) string {
	//TODO: Better handling of the arrow decoration
	bottomDecoration, _ := b.Config.GetBottomDecoration()
	if rightPart.Len == 0 {
		return leftPart.Str + bottomDecoration
	}
	spacer := b.Config.GetSpacer(leftPart.Len+rightPart.Len, b.TerminalWidth)
	topRow := leftPart.Str + spacer + rightPart.Str + "\n"
	bottomRow := bottomDecoration
	return topRow + bottomRow
}

func (b *PromptPartBuilder) buildPart() PromptPart {

	for _, component := range b.Components {
		componentBuilder := b.NewComponentBuilder(component).buildComponent()
		b.PromptPartStr += componentBuilder.Str
		b.PromptPartLen += componentBuilder.Len
	}
	return PromptPart{
		Str: b.PromptPartStr,
		Len: b.PromptPartLen,
	}
}

func (b *ComponentBuilder) buildComponent() Component {
	log.Debug().Msgf("[ComponentBuilder] Building component %s", b.Component.Name)

	b.addContent().addIcon().addPadding().addDividers().addMargin()
	return Component{
		Str: b.ComponentStr,
		Len: b.ComponentLen,
	}
}

func (b *ComponentBuilder) addIcon() *ComponentBuilder {

	// TODO: Handle icon separators

	icon := b.Component.Content.Icon
	var isRight bool
	padding := b.Component.Content.IconStyle.Padding

	isRight = b.Component.Content.IconStyle.IconPosition == "right"
	if icon == "" {
		log.Debug().Msgf("[ComponentBuilder] Component %s icon is empty, skipping icon", b.Component.Name)
		return b
	}
	if b.ComponentLen == 0 {
		log.Debug().Msgf("[ComponentBuilder] Component %s is empty, skipping icon", b.Component.Name)
		return b
	}

	iconBackgroundColor := b.Component.Content.IconStyle.BackgroundColor
	iconForegroundColor := b.Component.Content.IconStyle.ForegroundColor

	var iconString string
	if isRight {
		iconString = getPaddingString(" ", padding) + icon
	} else {
		iconString = icon + getPaddingString(" ", padding)
	}
	colorizedIconString := b.Config.ColorizeString(iconString, iconForegroundColor, iconBackgroundColor)

	if isRight {
		b.ComponentStr = b.ComponentStr + colorizedIconString
		b.ComponentLen += 2
	} else {
		b.ComponentStr = colorizedIconString + b.ComponentStr
		b.ComponentLen += 2
	}
	return b
}

func getPaddingString(paddingChar string, len int) string {
	if len <= 0 {
		return ""
	}
	paddingString := strings.Repeat(paddingChar, len)
	return paddingString
}

func (b *ComponentBuilder) addMargin() *ComponentBuilder {
	marginLeft := b.Component.Style.MarginLeft
	marginRight := b.Component.Style.MarginRight

	marginStringLeft := getMarginString(" ", marginLeft, 0)
	marginStringRight := getMarginString(" ", marginRight, 0)

	if b.ComponentLen == 0 {
		log.Debug().Msgf("[ComponentBuilder] Component %s is empty, skipping margins", b.Component.Name)
		return b
	}

	b.ComponentStr = marginStringLeft + b.ComponentStr + marginStringRight
	b.ComponentLen += utf8.RuneCountInString(marginStringLeft) + utf8.RuneCountInString(marginStringRight)
	return b
}

func getMarginString(marginChar string, marginLen int, previousMarginLen int) string {
	if marginLen <= 0 {
		return ""
	}
	if marginLen-previousMarginLen <= 0 {
		return ""
	}
	marginString := strings.Repeat(marginChar, marginLen-previousMarginLen)
	return marginString
}

func (b *ComponentBuilder) addPadding() *ComponentBuilder {
	paddingLeft := b.Component.Style.PaddingLeft
	paddingRight := b.Component.Style.PaddingRight
	paddingStringLeft := getPaddingString(" ", paddingLeft)
	paddingStringRight := getPaddingString(" ", paddingRight)
	if b.ComponentLen == 0 {
		log.Debug().Msgf("[ComponentBuilder] Component %s is empty, skipping padding", b.Component.Name)
		return b
	}
	colorizedLeftPadding := b.Config.ColorizeString(paddingStringLeft, b.Component.Style.ForegroundColor, b.Component.Style.BackgroundColor)
	colorizedRightPadding := b.Config.ColorizeString(paddingStringRight, b.Component.Style.ForegroundColor, b.Component.Style.BackgroundColor)
	b.ComponentStr = colorizedLeftPadding + b.ComponentStr + colorizedRightPadding
	b.ComponentLen += utf8.RuneCountInString(paddingStringLeft) + utf8.RuneCountInString(paddingStringRight)
	return b
}

func (b *ComponentBuilder) addDividers() *ComponentBuilder {

	if b.ComponentLen == 0 {
		log.Debug().Msgf("[ComponentBuilder] Component %s is empty, skipping dividers", b.Component.Name)
		return b
	}
	leftDivider := b.Component.Style.StartDivider
	rightDivider := b.Component.Style.EndDivider

	// Check if dividers are empty. If so, use the theme dividers
	if leftDivider == "$default" {
		leftDivider = b.Config.Theme.ComponentStartDivider
	}

	if rightDivider == "$default" {
		rightDivider = b.Config.Theme.ComponentEndDivider
	}

	if utf8.RuneCountInString(leftDivider) != 0 {
		colorizedLeftDivider := b.Config.ColorizeString(leftDivider, b.Component.Style.BackgroundColor, b.Config.Theme.BackgroundColor)
		b.ComponentStr = colorizedLeftDivider + b.ComponentStr
		b.ComponentLen += 1
	}

	if utf8.RuneCountInString(rightDivider) != 0 {
		colorizedRightDivider := b.Config.ColorizeString(rightDivider, b.Component.Style.BackgroundColor, b.Config.Theme.BackgroundColor)
		b.ComponentStr = b.ComponentStr + colorizedRightDivider
		b.ComponentLen += 1
	}

	return b
}

func (b *ComponentBuilder) addContent() *ComponentBuilder {

	moduleEntry, ok := b.Config.Modules[b.Component.Content.Module]
	if !ok {
		fmt.Fprintln(os.Stderr, "Module not found:", b.Component.Content.Module)
		return b
	}
	str, len := moduleEntry.Get(&b.Config, &b.Component)

	b.ComponentStr += str
	b.ComponentLen += len

	b.ComponentStr = b.Config.ColorizeString(b.ComponentStr, b.Component.Style.ForegroundColor, b.Component.Style.BackgroundColor)

	return b
}
