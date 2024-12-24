package config

import (
	"fmt"
	"os"
	"path/filepath"
	"promptorium/internal/pkg/confpkg/context"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

func GetRawConfig(configPath string, context *context.ApplicationContext, version string) RawConfig {
	// Load raw config
	path := getConfigPath(configPath)

	log.Trace().Msgf("Loading config")

	rawConfig := loadRawComponents(path)
	rawTheme := loadRawTheme(path)
	rawOptions := loadRawOptions(path)
	rawPrompt := loadRawPrompt(path)
	return RawConfig{
		Version:    version,
		Context:    context,
		Components: rawConfig,
		Theme:      rawTheme,
		Options:    rawOptions,
		Prompt:     rawPrompt,
	}
}

// Loads config from the configPath file in the raw format
func loadRawComponents(configPath string) []RawComponent {
	type RawConfigComponents struct {
		Components []RawComponent `yaml:"components"`
	}
	rawConf := RawConfigComponents{}
	type rawConfigComponentsString struct {
		Components string `yaml:"components"`
	}
	var componentsPath rawConfigComponentsString

	config_file, err := os.ReadFile(configPath)
	if err != nil {
		log.Trace().Msg("Could not read config file, using default config")
		return getDefaultRawComponents()
	}
	err = yaml.Unmarshal(config_file, &componentsPath)
	if componentsPath.Components != "" {
		if !filepath.IsAbs(componentsPath.Components) {
			componentsPath.Components = filepath.Join(filepath.Dir(configPath), componentsPath.Components)
		}
		log.Info().Msgf("Loading config from %s", componentsPath.Components)
		config_file, err = os.ReadFile(componentsPath.Components)
		if err != nil {
			fmt.Fprintln(os.Stderr, "promptorium: Could not read config file, using default config")
			return getDefaultRawComponents()
		}
	}
	err = yaml.Unmarshal(config_file, &rawConf)
	if err != nil {
		fmt.Fprintln(os.Stderr, "promptorium: Could not parse config file, using default config")
		return getDefaultRawComponents()
	}

	return rawConf.Components
}

// Loads theme from the themePath file in the raw format
func loadRawTheme(themePath string) RawTheme {
	type rawConfigTheme struct {
		Theme RawTheme `yaml:"theme"`
	}

	type RawConfigThemeString struct {
		Theme string `yaml:"theme"`
	}
	rawConfigThemeString := RawConfigThemeString{}
	theme := rawConfigTheme{}

	if themePath == "" {
		log.Trace().Msg("Theme path is empty, using default theme path")
		themePath, _ = findFile(DEFAULT_CONFIG_PATH, []string{"config.yaml", "config.yml", "config.json", "conf.yaml", "conf.yml", "conf.json"})
	}
	// Look for conf.yaml, conf.yml or conf.json in the theme path
	themeFile, err := os.ReadFile(themePath)
	if err != nil {
		log.Trace().Msg("Could not read theme file, using default theme")
		return getDefaultRawTheme()
	}
	err = yaml.Unmarshal(themeFile, &rawConfigThemeString)
	if rawConfigThemeString.Theme != "" {
		if !filepath.IsAbs(rawConfigThemeString.Theme) {
			rawConfigThemeString.Theme = filepath.Join(filepath.Dir(themePath), rawConfigThemeString.Theme)
		}
		log.Info().Msgf("Loading theme from %s", rawConfigThemeString.Theme)
		themeFile, err = os.ReadFile(rawConfigThemeString.Theme)
		if err != nil {
			log.Trace().Msg("Could not read theme file, using default theme")
			return getDefaultRawTheme()
		}
	}

	err = yaml.Unmarshal(themeFile, &theme)
	if err != nil {
		log.Trace().Msg("Could not unmarshal theme file, using default theme")
		return getDefaultRawTheme()
	}
	return theme.Theme
}

func loadRawOptions(optionsPath string) RawOptions {
	type RawConfigOptions struct {
		Options RawOptions `yaml:"options"`
	}

	type RawConfigOptionsString struct {
		Options string `yaml:"options"`
	}
	rawConfigOptionsString := RawConfigOptionsString{}

	rawConfigOptions := RawConfigOptions{}
	if optionsPath == "" {
		log.Trace().Msg("Options path is empty, using default options path")
		optionsPath, _ = findFile(DEFAULT_CONFIG_PATH, []string{"config.yaml", "config.yml", "config.json, conf.yaml", "conf.yml", "conf.json"})
	}
	// Look for conf.yaml, conf.yml or conf.json in the options path
	optionsFile, err := os.ReadFile(optionsPath)
	if err != nil {
		log.Trace().Msg("Could not read options file, using default options")
		return rawConfigOptions.Options
	}

	err = yaml.Unmarshal(optionsFile, &rawConfigOptionsString)
	if rawConfigOptionsString.Options != "" {
		if !filepath.IsAbs(rawConfigOptionsString.Options) {
			rawConfigOptionsString.Options = filepath.Join(filepath.Dir(optionsPath), rawConfigOptionsString.Options)
		}
		log.Info().Msgf("Loading options from %s", rawConfigOptionsString.Options)
		optionsFile, err = os.ReadFile(rawConfigOptionsString.Options)
		if err != nil {
			log.Trace().Msg("Could not read options file, using default options")
			return rawConfigOptions.Options
		}
	}

	err = yaml.Unmarshal(optionsFile, &rawConfigOptions)
	if err != nil {
		log.Trace().Msg("Could not unmarshal options file, using default options")
		return rawConfigOptions.Options
	}
	return rawConfigOptions.Options
}

func loadRawPrompt(promptPath string) [][]string {
	type RawConfigPrompt struct {
		Prompt []string `yaml:"prompt"`
	}
	type RawConfigMultilinePrompt struct {
		Prompt [][]string `yaml:"prompt"`
	}
	type RawConfigPromptString struct {
		Prompt string `yaml:"prompt"`
	}
	defaultPrompt := getDefaultRawConfig().Prompt
	rawConfigPrompt := RawConfigPrompt{}
	rawConfigMultilinePrompt := RawConfigMultilinePrompt{}
	rawConfigPromptString := RawConfigPromptString{}

	if promptPath == "" {
		log.Trace().Msg("Prompt path is empty, using default prompt path")
		promptPath, _ = findFile(DEFAULT_CONFIG_PATH, []string{"config.yaml", "config.yml", "config.json, conf.yaml", "conf.yml", "conf.json"})
	}
	// Look for conf.yaml, conf.yml or conf.json in the prompt path
	promptFile, err := os.ReadFile(promptPath)
	if err != nil {
		log.Trace().Msg("Could not read prompt file, using default prompt")
		return defaultPrompt
	}

	// Try to unmarshal prompt field as a string, if it fails, try to unmarshal as a struct

	err = yaml.Unmarshal(promptFile, &rawConfigPromptString)
	if rawConfigPromptString.Prompt != "" {
		// If the prompt field is a string, look for the file in the specified path
		if !filepath.IsAbs(rawConfigPromptString.Prompt) {
			rawConfigPromptString.Prompt = filepath.Join(filepath.Dir(promptPath), rawConfigPromptString.Prompt)
		}
		log.Info().Msgf("Loading prompt from %s", rawConfigPromptString.Prompt)
		promptFile, err = os.ReadFile(rawConfigPromptString.Prompt)
		if err != nil {
			log.Trace().Msg("Could not read prompt file, using default prompt")
			return defaultPrompt
		}
		return defaultPrompt
	}

	// Try to unmarshal the prompt as an array of strings
	err = yaml.Unmarshal(promptFile, &rawConfigMultilinePrompt)
	if err != nil {
		err = yaml.Unmarshal(promptFile, &rawConfigPrompt)
		if err != nil {
			log.Trace().Msg("Could not unmarshal prompt file, using default prompt")
			return defaultPrompt
		}
		return [][]string{rawConfigPrompt.Prompt}
	}

	return rawConfigMultilinePrompt.Prompt

}

/*
 * Helper functions
 */

// Function that looks for specified files in given directory and returns the first found file
func findFile(dir string, files []string) (string, error) {
	log.Trace().Msgf("Looking for %s in %s", files, dir)
	for _, file := range files {
		path := filepath.Join(dir, file)
		if _, err := os.Stat(path); err == nil {
			log.Trace().Msgf("Found %s in %s", file, dir)
			return path, nil
		}
	}
	return "", fmt.Errorf("promptorium: Could not find file")
}

func getConfigPath(configPath string) string {
	if configPath == "" {
		log.Trace().Msg("Config path is empty, using default config path")
		configPath, _ = findFile(DEFAULT_CONFIG_PATH, []string{"config.yaml", "config.yml", "config.json"})
	}

	type rawConfigPreset struct {
		Preset string `yaml:"preset"`
	}

	rawConf := rawConfigPreset{}

	config_file, err := os.ReadFile(configPath)
	if err != nil {
		log.Trace().Msg("Could not read config file")
		return configPath
	}
	err = yaml.Unmarshal(config_file, &rawConf)
	if err != nil {
		log.Trace().Msg("No preset found in config file")
		return configPath
	}

	if rawConf.Preset == "" {
		log.Trace().Msg("No preset found in config file")
		return configPath
	}

	log.Trace().Msgf("Found preset %s in config file", rawConf.Preset)

	_, err = os.Stat(filepath.Join(DEFAULT_PRESET_PATH, rawConf.Preset))
	if err != nil {
		log.Trace().Msgf("Could not find preset directory %s", filepath.Join(DEFAULT_PRESET_PATH, rawConf.Preset))
		fmt.Fprintln(os.Stderr, "promptorium: Could not find preset directory", filepath.Join(DEFAULT_PRESET_PATH, rawConf.Preset))
		return configPath
	}

	configPath, err = findFile(filepath.Join(DEFAULT_PRESET_PATH, rawConf.Preset), []string{"config.yaml", "config.yml", "config.json"})
	if err != nil {
		log.Trace().Msgf("Could not find preset config file in directory %s", filepath.Join(DEFAULT_PRESET_PATH, rawConf.Preset))
		fmt.Fprintln(os.Stderr, "promptorium: Could not find preset config file in directory %s", filepath.Join(DEFAULT_PRESET_PATH, rawConf.Preset))
		return configPath
	}
	log.Trace().Msgf("Using preset config file %s", configPath)
	return configPath
}

/*
 * Types
 */

type RawConfig struct {
	Version    string
	Context    *context.ApplicationContext
	Components []RawComponent
	Theme      RawTheme
	Options    RawOptions
	Prompt     [][]string
}

type RawColorName string

type RawComponent struct {
	Name    string            `yaml:"name"`
	Type    RawComponentType  `yaml:"type"`
	Content string            `yaml:"content"`
	Style   RawComponentStyle `yaml:"style"`
}

type RawIcon string

type RawComponentType string

var RawComponentTypes = map[string]RawComponentType{
	"module": RawComponentType("module"),
	"plugin": RawComponentType("plugin"),
	"text":   RawComponentType("text"),
}

type RawComponentStyle struct {
	BackgroundColor     RawColorName    `yaml:"background_color,omitempty"`
	ForegroundColor     RawColorName    `yaml:"foreground_color,omitempty"`
	StartDivider        string          `yaml:"start_divider,omitempty"`
	EndDivider          string          `yaml:"end_divider,omitempty"`
	Margin              string          `yaml:"margin,omitempty"`
	Padding             string          `yaml:"padding,omitempty"`
	Icon                RawIcon         `yaml:"icon,omitempty"`
	IconPosition        RawIconPosition `yaml:"icon_position,omitempty"`
	IconPadding         string          `yaml:"icon_padding,omitempty"`
	IconSeparator       string          `yaml:"icon_separator,omitempty"`
	IconForegroundColor RawColorName    `yaml:"icon_foreground_color,omitempty"`
	IconBackgroundColor RawColorName    `yaml:"icon_background_color,omitempty"`
}

type RawAlign string

var RawIconPositions = map[string]RawIconPosition{
	"left":  RawIconPosition("left"),
	"right": RawIconPosition("right"),
}

type RawIconPosition string

type RawTheme struct {
	ComponentStartDivider      string       `yaml:"component_start_divider,omitempty"`
	ComponentEndDivider        string       `yaml:"component_end_divider,omitempty"`
	Spacer                     string       `yaml:"component_spacer,omitempty"`
	PrimaryColor               RawColorName `yaml:"primary_color,omitempty"`
	SecondaryColor             RawColorName `yaml:"secondary_color,omitempty"`
	TertiaryColor              RawColorName `yaml:"tertiary_color,omitempty"`
	QuaternaryColor            RawColorName `yaml:"quaternary_color,omitempty"`
	SuccessColor               RawColorName `yaml:"success_color,omitempty"`
	WarningColor               RawColorName `yaml:"warning_color,omitempty"`
	ErrorColor                 RawColorName `yaml:"error_color,omitempty"`
	BackgroundColor            RawColorName `yaml:"background_color,omitempty"`
	ForegroundColor            RawColorName `yaml:"foreground_color,omitempty"`
	GitStatusColorClean        RawColorName `yaml:"git_status_clean,omitempty"`
	GitStatusColorDirty        RawColorName `yaml:"git_status_dirty,omitempty"`
	GitStatusColorNoRepository RawColorName `yaml:"git_status_no_repository,omitempty"`
	GitStatusColorNoUpstream   RawColorName `yaml:"git_status_no_upstream,omitempty"`
	ExitCodeColorOk            RawColorName `yaml:"exit_code_ok,omitempty"`
	ExitCodeColorError         RawColorName `yaml:"exit_code_error,omitempty"`
}

type RawOptions struct {
	CWD RawCwdOptions `yaml:"cwd"`
}

type RawCwdOptions struct {
	HighlightGitRoot bool `yaml:"highlight_git_root"`
}
