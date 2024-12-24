package config

//This package contains the logic for parsing the config file and theme file into a Config object.

import (
	"os"
	"path/filepath"
	"promptorium/internal/pkg/confpkg/context"

	"github.com/rs/zerolog/log"
)

var Colors map[string]Color = getColors()

var DEFAULT_PLUGIN_PATH = filepath.Join(os.Getenv("HOME"), ".config", "promptorium", "plugins")
var DEFAULT_CONFIG_PATH = filepath.Join(os.Getenv("HOME"), ".config", "promptorium")
var DEFAULT_PRESET_PATH = filepath.Join(os.Getenv("HOME"), ".config", "promptorium", "presets")

// GetConfig reads the config file and theme file from the paths specified in
// the passed arguments, and returns a parsed Config object.
// If the configPath or themePath arguments are empty, it uses the default paths.
func GetConfig(configPath string, shell string, exitCode int, version string) Config {

	context := context.GetApplicationContext(shell, exitCode)

	conf, err := ParseConfig(GetRawConfig(configPath, context, version))
	if err != nil {
		log.Trace().Msg("Error parsing config")
		return conf
	}

	return conf
}

type Config struct {
	Version    string
	Prompt     [][]string
	Theme      Theme
	Components map[string]Component
	Context    *context.ApplicationContext
	Options    ConfigOptions
	Modules    map[string]ModuleEntry
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

type ModuleEntry struct {
	Get func(config *Config, component *Component) []ComponentContent
}

type Component struct {
	Name    string
	Type    ComponentType
	Style   ComponentStyle
	Content string
	Icon    string
}

type ComponentType string

type ComponentStyle struct {
	BackgroundColor     Color
	ForegroundColor     Color
	StartDivider        string
	EndDivider          string
	MarginLeft          int
	MarginRight         int
	PaddingLeft         int
	PaddingRight        int
	Align               Align
	IconPosition        IconPosition
	IconPadding         int
	IconSeparator       string
	IconForegroundColor Color
	IconBackgroundColor Color
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

var ComponentTypes = map[string]ComponentType{
	"module": ComponentType("module"),
	"plugin": ComponentType("plugin"),
	"text":   ComponentType("text"),
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

// Options
type ConfigOptions struct {
	CWD CwdOptions
}

type CwdOptions struct {
	HighlightGitRoot bool
}
