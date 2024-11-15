package confmodule

import (
	"encoding/json"
	"promptorium/cmd/modules/confmodule/context"
)

type Config struct {
	Version    string
	Theme      Theme
	Components []Component
	Context    *context.ApplicationContext
	Modules    map[string]ModuleEntry
}

type ModuleEntry struct {
	Name string
	Get  func(config *Config, component *Component) (string, int)
}
type Theme struct {
	ComponentStartDivider      string
	ComponentEndDivider        string
	Spacer                     string
	PrimaryColor               Color
	SecondaryColor             Color
	TertiaryColor              Color
	QuaternaryColor            Color
	SuccessColor               Color
	WarningColor               Color
	ErrorColor                 Color
	BackgroundColor            Color
	ForegroundColor            Color
	GitStatusColorClean        Color
	GitStatusColorDirty        Color
	GitStatusColorNoRepository Color
	GitStatusColorNoUpstream   Color
	ExitCodeColorOk            Color
	ExitCodeColorError         Color
}

type Component struct {
	Name    string
	Style   ComponentStyle
	Content Content
}

type Content struct {
	Module    string
	Icon      string
	IconStyle IconStyle
}

type ComponentStyle struct {
	BackgroundColor Color
	ForegroundColor Color
	StartDivider    string
	EndDivider      string
	MarginLeft      int
	MarginRight     int
	PaddingLeft     int
	PaddingRight    int
	Align           Align
}

type Align string

type IconPosition string

var IconPositions = map[string]IconPosition{
	"left":  IconPosition("left"),
	"right": IconPosition("right"),
}

var Alignments = map[string]Align{
	"left":  Align("left"),
	"right": Align("right"),
}

type IconStyle struct {
	BackgroundColor Color
	ForegroundColor Color
	IconPosition    IconPosition
	Padding         int
	Separator       string
}

type ModuleStyle struct {
	BackgroundColor Color
	ForegroundColor Color
	MarginLeft      int
	MarginRight     int
	Separator       string
}

type Color struct {
	BackgroudCode  string
	ForegroundCode string
	Name           string
}

type ColorName string

// Raw config

type RawConfig struct {
	Version    string         `json:"version,omitempty"`
	Preset     string         `json:"preset,omitempty"`
	Components []RawComponent `json:"components,omitempty"`
}

type RawColorName string

type RawComponent struct {
	Name    string            `json:"name,omitempty"`
	Content RawContent        `json:"content,omitempty"`
	Style   RawComponentStyle `json:"style,omitempty"`
}

type RawContent struct {
	Module    string       `json:"module,omitempty"`
	Icon      string       `json:"icon,omitempty"`
	IconStyle RawIconStyle `json:"icon_style,omitempty"`
}

type RawComponentStyle struct {
	BackgroundColor RawColorName `json:"background_color,omitempty"`
	ForegroundColor RawColorName `json:"foreground_color,omitempty"`
	StartDivider    string       `json:"start_divider,omitempty"`
	EndDivider      string       `json:"end_divider,omitempty"`
	Margin          string       `json:"margin,omitempty"`
	Padding         string       `json:"padding,omitempty"`
	Align           Align        `json:"align,omitempty"`
}

type RawIconStyle struct {
	BackgroundColor RawColorName `json:"background_color,omitempty"`
	ForegroundColor RawColorName `json:"foreground_color,omitempty"`
	Padding         string       `json:"padding,omitempty"`
	IconPosition    IconPosition `json:"icon_position,omitempty"`
	Separator       string       `json:"separator,omitempty"`
}

type RawTheme struct {
	ComponentStartDivider      string       `json:"component_start_divider,omitempty"`
	ComponentEndDivider        string       `json:"component_end_divider,omitempty"`
	Spacer                     string       `json:"spacer,omitempty"`
	PrimaryColor               RawColorName `json:"primary_color,omitempty"`
	SecondaryColor             RawColorName `json:"secondary_color,omitempty"`
	TertiaryColor              RawColorName `json:"tertiary_color,omitempty"`
	QuaternaryColor            RawColorName `json:"quaternary_color,omitempty"`
	SuccessColor               RawColorName `json:"success_color,omitempty"`
	WarningColor               RawColorName `json:"warning_color,omitempty"`
	ErrorColor                 RawColorName `json:"error_color,omitempty"`
	BackgroundColor            RawColorName `json:"background_color,omitempty"`
	ForegroundColor            RawColorName `json:"foreground_color,omitempty"`
	GitStatusColorClean        RawColorName `json:"git_status_clean,omitempty"`
	GitStatusColorDirty        RawColorName `json:"git_status_dirty,omitempty"`
	GitStatusColorNoRepository RawColorName `json:"git_status_no_repository,omitempty"`
	GitStatusColorNoUpstream   RawColorName `json:"git_status_no_upstream,omitempty"`
	ExitCodeColorOk            RawColorName `json:"exit_code_ok,omitempty"`
	ExitCodeColorError         RawColorName `json:"exit_code_error,omitempty"`
}

// TODO: move this to a separate file

/*
 * Default values
 */

func (a *Align) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	_, ok := Alignments[s]
	if !ok {
		*a = Align("left")
		return err
	}
	*a = Align(s)
	return nil
}

func (c *RawComponent) UnmarshalJSON(data []byte) error {
	type xcomponent RawComponent
	var x xcomponent
	x.Style = getDefaultRawComponentStyle()

	err := json.Unmarshal(data, &x)
	if err != nil {
		return err
	}
	*c = RawComponent(x)
	return nil
}

func (t *RawTheme) UnmarshalJSON(data []byte) error {
	type xtheme RawTheme
	x := xtheme(getDefaultRawTheme())

	err := json.Unmarshal(data, &x)
	if err != nil {
		return err
	}
	*t = RawTheme(x)
	return nil
}

func (c *RawContent) UnmarshalJSON(data []byte) error {
	type xcontent RawContent
	var x xcontent
	x.IconStyle = getDefaultRawIconStyle()
	err := json.Unmarshal(data, &x)
	if err != nil {
		return err
	}
	*c = RawContent(x)
	return nil
}

func (i *RawIconStyle) UnmarshalJSON(data []byte) error {
	type xiconstyle RawIconStyle
	var x xiconstyle = xiconstyle(getDefaultRawIconStyle())
	err := json.Unmarshal(data, &x)
	if err != nil {
		return err
	}
	*i = RawIconStyle(x)
	return nil
}
