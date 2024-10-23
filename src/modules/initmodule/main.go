package initmodule

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func InitPromptorium() {
	// Initialize promptorium
	fmt.Println("Initializing promptorium")

	copyConfigFiles()
	copyThemeFiles()
	copyPresetFiles()
	// Add line to ~/.bashrc and/or ~/.zshrc to source promptorium shell
	addSourceLines()
}

func copyConfigFiles() {
	// Check if config files already exist
	_, err := os.Stat("/home/" + os.Getenv("USER") + "/.config/promptorium")
	if err == nil {
		fmt.Println("Found existing config files")
		return
	}
	// Copy config files
	fmt.Println("Copying config files")
	user := os.Getenv("USER")
	cmd := exec.Command("sudo", "cp", "-r", "/usr/share/promptorium/conf", "/home/"+user+"/.config/promptorium")
	cmd.Run()
}

func copyThemeFiles() {
	// Check if theme files already exist
	_, err := os.Stat("/home/" + os.Getenv("USER") + "/.config/promptorium")
	if err == nil {
		fmt.Println("Found existing theme files")
		return
	}
	// Copy theme files
	fmt.Println("Copying theme files")
	user := os.Getenv("USER")
	cmd := exec.Command("sudo", "cp", "-r", "/usr/share/promptorium/theme", "/home/"+user+"/.config/promptorium")
	cmd.Run()
}

func copyPresetFiles() {
	// Check if preset files already exist
	_, err := os.Stat("/home/" + os.Getenv("USER") + "/.config/promptorium")
	if err == nil {
		fmt.Println("Found existing preset files")
		return
	}
	// Copy preset files
	fmt.Println("Copying preset files")
	user := os.Getenv("USER")
	cmd := exec.Command("sudo", "cp", "-r", "/usr/share/promptorium/presets", "/home/"+user+"/.config/promptorium")
	cmd.Run()
}

func addSourceLines() {
	// Add line to ~/.bashrc and/or ~/.zshrc to source promptorium shell
	_, err := os.Stat("/home/" + os.Getenv("USER") + "/.zshrc")
	if err == nil {
		fmt.Println("Found .zshrc file")
		addSourceLine("/home/" + os.Getenv("USER") + "/.zshrc")
	}

	_, err = os.Stat("/home/" + os.Getenv("USER") + "/.bashrc")
	if err == nil {
		fmt.Println("Found .bashrc file")
		addSourceLine("/home/" + os.Getenv("USER") + "/.bashrc")
	}
}

func addSourceLine(file string) {
	// Add line to .bashrc or .zshrc to source promptorium shell
	var lineToAdd = "if [[ $(command -v promptorium) 2> /dev/null ]]; then source <(promptorium shell); fi"
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	// Check if line already exists
	file_contents, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	if strings.Contains(string(file_contents), lineToAdd) {
		fmt.Println("Line already exists in file")
		return
	}
	// Add line
	defer f.Close()
	fmt.Println("Adding line to " + file)
	if _, err = f.WriteString("\n" + lineToAdd + "\n"); err != nil {
		fmt.Println(err)
	}
}
