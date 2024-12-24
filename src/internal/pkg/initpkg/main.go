package initpkg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/rs/zerolog/log"
)

func InitPromptorium() {
	fmt.Println("Initializing promptorium...")

	createDirectory()
	copyConfigFiles()
	copyPresetFiles()
	giveFilePermissions()
	addSourceLines()

	fmt.Println("Done!")
}
func createDirectory() {
	// Check if ~/.config/promptorium directory already exists
	_, err := os.Stat("/home/" + os.Getenv("USER") + "/.config/promptorium")
	if err == nil {
		log.Trace().Msg("Found existing ~/.config/promptorium directory")
		return
	}
	// Create ~/.config/promptorium directory
	fmt.Println("Creating ~/.config/promptorium directory..")
	user := os.Getenv("USER")
	cmd := exec.Command("sudo", "mkdir", "-p", "/home/"+user+"/.config/promptorium")
	cmd.Run()
}

func copyConfigFiles() {
	// Check if config files already exist
	_, err := os.Stat("/home/" + os.Getenv("USER") + "/.config/promptorium/conf.json")
	if err == nil {
		log.Trace().Msg("Found existing config file")
		return
	}
	// Check if config files exist in /usr/share/promptorium/conf
	out, err := exec.Command("sudo", "ls", "/usr/share/promptorium/conf").Output()
	if err != nil {
		log.Trace().Msg("Error checking for config files in /usr/share/promptorium/conf")
		fmt.Println()
		return
	}

	if !strings.Contains(string(out), "conf.json") {
		log.Trace().Msg("Config files not found in /usr/share/promptorium/conf")
		return
	}

	// Copy config files
	fmt.Println("Copying config files...")
	user := os.Getenv("USER")
	cmd := exec.Command("sudo", "cp", "/usr/share/promptorium/conf/conf.json", "/home/"+user+"/.config/promptorium/conf.json")
	cmd.Run()
}

func giveFilePermissions() {
	// Give file permissions to user
	user := os.Getenv("USER")
	log.Trace().Msgf("Giving file permissions to user %s", user)
	cmd := exec.Command("sudo", "chown", "-R", user+":"+user, "/home/"+user+"/.config/promptorium")
	cmd.Run()
}

func copyPresetFiles() {
	// Check if preset files already exist
	_, err := os.Stat("/home/" + os.Getenv("USER") + "/.config/promptorium/presets")
	if err == nil {
		log.Trace().Msg("Found existing preset files")
		return
	}
	// Check if preset files exist in /usr/share/promptorium/conf
	out, err := exec.Command("sudo", "ls", "/usr/share/promptorium/conf").Output()
	if err != nil {
		log.Trace().Msg("Error checking for preset files in /usr/share/promptorium/conf")
		return
	}
	if strings.Contains(string(out), "presets") {
		log.Trace().Msg("Preset files not found in /usr/share/promptorium/conf")
		return
	}

	// Copy preset files
	log.Trace().Msg("Copying preset files")
	user := os.Getenv("USER")
	cmd := exec.Command("sudo", "cp", "-r", "/usr/share/promptorium/conf/presets", "/home/"+user+"/.config/promptorium/presets")
	cmd.Run()
}

func addSourceLines() {
	// Add line to ~/.bashrc and/or ~/.zshrc to source promptorium shell
	_, err := os.Stat("/home/" + os.Getenv("USER") + "/.zshrc")
	if err == nil {
		log.Trace().Msg("Found .zshrc file")
		addSourceLine("/home/"+os.Getenv("USER")+"/.zshrc", "zsh")
	}

	_, err = os.Stat("/home/" + os.Getenv("USER") + "/.bashrc")
	if err == nil {
		log.Trace().Msg("Found .bashrc file")
		addSourceLine("/home/"+os.Getenv("USER")+"/.bashrc", "bash")
	}
}

func addSourceLine(file string, shell string) {
	// Add line to .bashrc or .zshrc to source promptorium shell
	var lineToAdd = "if [[ -n $(command -v promptorium) ]]; then source <(promptorium shell --shell " + shell + "); fi"
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	// Check if line already exists
	file_contents, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if strings.Contains(string(file_contents), lineToAdd) {
		log.Trace().Msg("Line already exists in file")
		return
	}
	// Add line
	defer f.Close()
	log.Trace().Msgf("Adding line to %s", file)
	if _, err = f.WriteString("\n" + lineToAdd); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
