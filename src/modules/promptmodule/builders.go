package promptmodule

import (
	"fmt"
	"os"
	"promptorium/modules/confmodule"
	"promptorium/modules/utils"
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
	bottomDecoration, _ := utils.GetBottomDecoration(b.Config)
	spacer := utils.GetSpacer(b.Config, leftPart.Len+rightPart.Len, b.TerminalWidth)
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
	log.Debug().Msgf("Building component %s", b.Component.Name)

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
		log.Debug().Msgf("Component icon is empty, skipping icon")
		return b
	}
	if b.ComponentLen == 0 {
		log.Debug().Msgf("Component is empty, skipping icon")
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
	colorizedIconString := utils.Colorize(iconString, iconForegroundColor, iconBackgroundColor, false, b.Config.State.Shell)

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
		log.Debug().Msgf("Component is empty, skipping margins")
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
		log.Debug().Msgf("Component is empty, skipping padding")
		return b
	}

	b.ComponentStr = utils.Colorize(paddingStringLeft, b.Component.Style.ForegroundColor, b.Component.Style.BackgroundColor, false, b.Config.State.Shell) + b.ComponentStr + utils.Colorize(paddingStringRight, b.Component.Style.ForegroundColor, b.Component.Style.BackgroundColor, false, b.Config.State.Shell)
	b.ComponentLen += utf8.RuneCountInString(paddingStringLeft) + utf8.RuneCountInString(paddingStringRight)
	return b
}

func (b *ComponentBuilder) addDividers() *ComponentBuilder {

	if b.ComponentLen == 0 {
		log.Debug().Msgf("Component is empty, skipping dividers")
		return b
	}
	leftDivider := b.Component.Style.StartDivider
	rightDivider := b.Component.Style.EndDivider
	if utf8.RuneCountInString(leftDivider) != 0 {
		colorizedLeftDivider := utils.Colorize(leftDivider, b.Component.Style.BackgroundColor, b.Config.Theme.BackgroundColor, false, b.Config.State.Shell)
		b.ComponentStr = colorizedLeftDivider + b.ComponentStr
		b.ComponentLen += 1
	}

	if utf8.RuneCountInString(rightDivider) != 0 {
		colorizedRightDivider := utils.Colorize(rightDivider, b.Component.Style.BackgroundColor, b.Config.Theme.BackgroundColor, false, b.Config.State.Shell)
		b.ComponentStr = b.ComponentStr + colorizedRightDivider
		b.ComponentLen += 1
	}

	return b
}

func (b *ComponentBuilder) addContent() *ComponentBuilder {
	switch b.Component.Content.Module {

	case "time":
		b.buildTimeModuleContent()

	case "hostname":
		b.buildHostnameModuleContent()

	case "cwd":
		b.buildCwdModuleContent()

	case "git_branch":
		b.buildGitBranchModuleContent()

	case "exit_status":
		b.buildExitStatusModuleContent()

	case "user":
		b.buildUserModuleContent()

	case "os_icon":
		b.buildOsIconModuleContent()

	case "git_status":
		b.buildGitStatusModuleContent()

	default:
		fmt.Fprintln(os.Stderr, "Module not implemented:", b.Component.Content.Module)
	}

	b.ComponentStr = utils.Colorize(b.ComponentStr, b.Component.Style.BackgroundColor, b.Component.Style.ForegroundColor, true, b.Config.State.Shell)

	return b
}
