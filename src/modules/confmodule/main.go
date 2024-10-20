package confmodule

//This package contains the logic for parsing the config file and theme file into a Config object.

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

var Colors map[string]Color = getColors()

const config_file = "conf.json"

const theme_file = "theme.json"

var config_path = filepath.Join(os.Getenv("HOME"), ".config", "promptorium", config_file)

var theme_path = filepath.Join(os.Getenv("HOME"), ".config", "promptorium", theme_file)

// GetConfig reads the config file and theme file from the paths specified in
// the passed arguments, and returns a parsed Config object.
// If the configPath or themePath arguments are empty, it uses the default paths.
func GetConfig(configPath string, themePath string, shell string, exitCode int) Config {
	log.Debug().Msg("Loading config")

	if configPath == "" {
		log.Debug().Msg("Config path is empty, using default config path")
		configPath = config_path
	}
	if themePath == "" {
		log.Debug().Msg("Theme path is empty, using default theme path")
		themePath = theme_path
	}
	// Load raw config
	rawConfig := loadRawConfig(configPath)
	rawTheme := loadRawTheme(themePath)
	// Check if the raw config contains a preset
	if rawConfig.Preset != "" {
		// Load the preset
		presetConfigPath := filepath.Join(os.Getenv("HOME"), ".config", "promptorium", "presets", rawConfig.Preset, "conf.json")
		presetThemePath := filepath.Join(os.Getenv("HOME"), ".config", "promptorium", "presets", rawConfig.Preset, "theme.json")
		log.Debug().Msgf("Loading preset %s", rawConfig.Preset)
		log.Debug().Msgf("Preset config path: %s", presetConfigPath)
		log.Debug().Msgf("Preset theme path: %s", presetThemePath)
		rawConfig = loadRawConfig(presetConfigPath)
		rawTheme = loadRawTheme(presetThemePath)
	}

	conf, err := parseConfig(rawTheme, rawConfig, getApplicationState(shell, exitCode))

	if err != nil {
		log.Debug().Msg("Error parsing config")
		return conf
	}

	return conf
}

// Loads config from the configPath file in the raw format
func loadRawConfig(configPath string) RawConfig {
	conf := RawConfig{}
	config_file, err := os.ReadFile(configPath)
	if err != nil {
		log.Debug().Msg("Could not read config file, using default config")
		conf = getDefaultRawConfig()
		return conf
	}
	err = json.Unmarshal(config_file, &conf)
	if err != nil {
		log.Debug().Msg("Error unmarshalling config file")
		return conf
	}
	return conf
}

// Loads theme from the themePath file in the raw format
func loadRawTheme(themePath string) RawTheme {
	theme := RawTheme{}
	themeFile, err := os.ReadFile(themePath)
	if err != nil {
		log.Debug().Msg("Could not read theme file, using default theme")
		theme = getDefaultRawTheme()
		return theme
	}
	err = json.Unmarshal(themeFile, &theme)
	if err != nil {
		log.Debug().Msg("Error unmarshalling theme file")
		return theme
	}
	return theme
}
