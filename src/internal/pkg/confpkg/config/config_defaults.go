package config

import "gopkg.in/yaml.v3"

func getDefaultTheme() Theme {
	return Theme{
		ComponentStartDivider:      "",
		ComponentEndDivider:        "",
		Spacer:                     "─",
		PrimaryColor:               Colors["blue"],
		SecondaryColor:             Colors["green"],
		TertiaryColor:              Colors["magenta"],
		QuaternaryColor:            Colors["cyan"],
		SuccessColor:               Colors["green"],
		WarningColor:               Colors["yellow"],
		ErrorColor:                 Colors["red"],
		BackgroundColor:            Colors["transparent"],
		ForegroundColor:            Colors["white"],
		GitStatusColorClean:        Colors["green"],
		GitStatusColorDirty:        Colors["yellow"],
		GitStatusColorNoRepository: Colors["blue"],
		GitStatusColorNoUpstream:   Colors["yellow"],
		ExitCodeColorOk:            Colors["green"],
		ExitCodeColorError:         Colors["red"],
	}
}

func getColors() map[string]Color {

	colors := map[string]Color{}

	colors["black"] = Color{
		ForegroundCode: "30",
		BackgroudCode:  "40",
		Name:           "black",
	}
	colors["red"] = Color{
		ForegroundCode: "31",
		BackgroudCode:  "41",
		Name:           "red",
	}
	colors["green"] = Color{
		ForegroundCode: "32",
		BackgroudCode:  "42",
		Name:           "green",
	}
	colors["yellow"] = Color{
		ForegroundCode: "33",
		BackgroudCode:  "43",
		Name:           "yellow",
	}
	colors["blue"] = Color{
		ForegroundCode: "34",
		BackgroudCode:  "44",
		Name:           "blue",
	}
	colors["magenta"] = Color{
		ForegroundCode: "35",
		BackgroudCode:  "45",
		Name:           "magenta",
	}
	colors["cyan"] = Color{
		ForegroundCode: "36",
		BackgroudCode:  "46",
		Name:           "cyan",
	}
	colors["white"] = Color{
		ForegroundCode: "37",
		BackgroudCode:  "47",
		Name:           "white",
	}
	colors["none"] = Color{
		ForegroundCode: "0",
		BackgroudCode:  "0",
		Name:           "none",
	}
	colors["transparent"] = Color{
		ForegroundCode: "39",
		BackgroudCode:  "49",
		Name:           "transparent",
	}

	return colors
}

// Default values

func (c *RawComponent) UnmarshalYAML(data *yaml.Node) error {
	type xcomponent RawComponent
	var x xcomponent
	x.Style = getDefaultRawComponentStyle()

	err := data.Decode(&x)
	if err != nil {
		return err
	}
	*c = RawComponent(x)
	return nil
}

func (t *RawTheme) UnmarshalYAML(data *yaml.Node) error {
	type xtheme RawTheme
	x := xtheme(getDefaultRawTheme())

	err := data.Decode(&x)
	if err != nil {
		return err
	}
	*t = RawTheme(x)
	return nil
}

func getDefaultRawComponentStyle() RawComponentStyle {

	//Colors are not set here, because they depend on the theme
	//Default values are set when parsing the config (which happens after parsing the theme)
	return RawComponentStyle{
		BackgroundColor:     "$default",
		ForegroundColor:     "$default",
		StartDivider:        "$default",
		EndDivider:          "$default",
		Margin:              "1 0",
		Padding:             "1 1",
		IconPosition:        RawIconPosition("left"),
		IconPadding:         "1",
		IconSeparator:       " ",
		IconForegroundColor: "$default",
		IconBackgroundColor: "$default",
	}
}

func getDefaultRawTheme() RawTheme {
	return RawTheme{
		ComponentStartDivider:      "",
		ComponentEndDivider:        "",
		Spacer:                     " ",
		PrimaryColor:               "blue",
		SecondaryColor:             "green",
		TertiaryColor:              "magenta",
		QuaternaryColor:            "cyan",
		SuccessColor:               "green",
		WarningColor:               "yellow",
		ErrorColor:                 "red",
		BackgroundColor:            "transparent",
		ForegroundColor:            "white",
		GitStatusColorClean:        "green",
		GitStatusColorDirty:        "yellow",
		GitStatusColorNoRepository: "blue",
		GitStatusColorNoUpstream:   "yellow",
		ExitCodeColorOk:            "green",
		ExitCodeColorError:         "red",
	}
}

func getDefaultRawComponents() []RawComponent {
	return []RawComponent{
		{
			Name:    "user",
			Type:    "module",
			Content: "user",
			Style: RawComponentStyle{
				Padding:         "0",
				Margin:          "0 1",
				BackgroundColor: "transparent",
				ForegroundColor: "white",
			},
		},
		{
			Name:    "cwd",
			Type:    "module",
			Content: "cwd",
			Style: RawComponentStyle{
				Icon:            "@",
				Padding:         "0",
				Margin:          "0 1",
				BackgroundColor: "transparent",
				ForegroundColor: "white",
			},
		},
	}

}
func getDefaultPrompt() [][]string {
	return [][]string{
		[]string{
			"user",
			"cwd",
			"---",
			"time",
		},
	}
}

func getDefaultRawConfig() RawConfig {
	return RawConfig{
		Components: getDefaultRawComponents(),
		Theme:      getDefaultRawTheme(),
		Prompt:     getDefaultPrompt(),
	}
}
