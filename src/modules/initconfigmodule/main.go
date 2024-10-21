package initconfigmodule

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var config_path = filepath.Join(os.Getenv("HOME"), ".config", "promptorium", "conf.json")
var theme_path = filepath.Join(os.Getenv("HOME"), ".config", "promptorium", "theme.json")
var presets_path = filepath.Join(os.Getenv("HOME"), ".config", "promptorium", "presets")
var bash_script_path = filepath.Join(os.Getenv("HOME"), ".config", "promptorium", "shell", "promptorium.bash")
var zsh_script_path = filepath.Join(os.Getenv("HOME"), ".config", "promptorium", "shell", "promptorium.zsh")

var default_config_path = filepath.Join("usr", "share", "promptorium", "conf.json")
var default_theme_path = filepath.Join("usr", "share", "promptorium", "theme.json")
var default_presets_path = filepath.Join("usr", "share", "promptorium", "presets")
var default_bash_script_path = filepath.Join("usr", "share", "promptorium", "shell", "promptorium.bash")
var default_zsh_script_path = filepath.Join("usr", "share", "promptorium", "shell", "promptorium.zsh")

func InitConfig() {

	fmt.Println("Initializing promptorium config and theme files")
	// Check if config and theme files exist
	checkConfigFile()
	checkThemeFile()
	checkPresets()
	checkBashScriptFile()
	checkZshScriptFile()
	modifyRcFiles()
}

func checkConfigFile() {

	if _, err := os.Stat(config_path); os.IsNotExist(err) {
		// Copy default config file
		fmt.Println("Copying default config file:", default_config_path)
		copyFile(default_config_path, config_path)
	} else {
		fmt.Println("Config file already exists:", config_path)
	}
}

func checkThemeFile() {

	if _, err := os.Stat(theme_path); os.IsNotExist(err) {
		// Copy default theme file
		fmt.Println("Copying default theme file:", default_theme_path)
		copyFile(default_theme_path, theme_path)
	} else {
		fmt.Println("Theme file already exists:", theme_path)
	}
}

func checkPresets() {

	if _, err := os.Stat(presets_path); os.IsNotExist(err) {
		// Copy default presets
		fmt.Println("Copying default presets:", default_presets_path)
		copyFile(default_presets_path, presets_path)
	} else {
		fmt.Println("Preset folder already exists:", presets_path)
	}
}

func checkBashScriptFile() {

	if _, err := os.Stat(bash_script_path); os.IsNotExist(err) {
		// Copy default bash script
		fmt.Println("Copying default bash script:", default_bash_script_path)
		copyFile(default_bash_script_path, bash_script_path)
	} else {
		fmt.Println("Bash script already exists:", bash_script_path)
	}
}

func checkZshScriptFile() {

	if _, err := os.Stat(zsh_script_path); os.IsNotExist(err) {
		// Copy default zsh script
		fmt.Println("Copying default zsh script:", default_zsh_script_path)
		copyFile(default_zsh_script_path, zsh_script_path)
	} else {
		fmt.Println("Zsh script already exists:", zsh_script_path)
	}
}

func modifyRcFiles() {
	var zshrc = filepath.Join(os.Getenv("HOME"), ".zshrc")
	var bashrc = filepath.Join(os.Getenv("HOME"), ".bashrc")
	var rc_files = []string{bashrc, zshrc}
	for _, rc_file := range rc_files {
		if _, err := os.Stat(rc_file); err == nil {
			fmt.Println("Found file", filepath.Base(rc_file))
			lineToAppend := ""
			switch filepath.Base(rc_file) {
			case ".zshrc":
				lineToAppend = "source ~/.config/promptorium/shell/promptorium.zsh"
			case ".bashrc":
				lineToAppend = "source ~/.config/promptorium/shell/promptorium.bash"
			default:

			}
			checkIfLineExistsAndAppend(lineToAppend, rc_file)
		}
	}
}

func checkIfLineExistsAndAppend(stringToAppend string, fileName string) {
	if _, err := os.Stat(fileName); err == nil {
		output, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, "promptorium: Error reading file", fileName)
			os.Exit(1)
		}
		for _, line := range strings.Split(string(output), "\n") {
			if strings.EqualFold(line, stringToAppend) {
				fmt.Println("Line already exists in", filepath.Base(fileName))
				return
			}
		}
	}

	fmt.Println("Do you want promptorium to modify it? (y/n)")
	var answer string
	fmt.Scanln(&answer)

	if strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes" {
		appendLineToFile(stringToAppend, fileName)
		fmt.Println("Added line to", filepath.Base(fileName))
	}
}

func appendLineToFile(stringToAppend string, fileName string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "promptorium: Error opening rc file")
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.WriteString("\n" + stringToAppend)
	if err != nil {
		fmt.Fprintln(os.Stderr, "promptorium: Error writing to rc file")
		os.Exit(1)
	}
}

func copyFile(src string, dst string) {
	data, err := os.ReadFile(src)
	if err != nil {
		fmt.Fprintln(os.Stderr, "promptorium: Error reading file", src)
		os.Exit(1)
	}
	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "promptorium: Error writing file", dst)
		os.Exit(1)
	}
}
