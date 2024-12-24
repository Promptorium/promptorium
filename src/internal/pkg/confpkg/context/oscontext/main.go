package oscontext

import (
	"fmt"
	"os/exec"
	"strings"
)

type OS int

const (
	OSLinux OS = iota
	OSMac
	OSFedora
	OSUbuntu
	OSDebian
	OSArch
	OSOther
)

func GetOS(result chan OS) {

	output, err := exec.Command("uname", "-s").Output()
	if err != nil {
		result <- OSOther
		return
	}
	os := strings.Replace(string(output), "\n", "", -1)
	if os == "" {
		result <- OSOther
		return
	}
	// Return directly if not Linux
	if os != "Linux" {
		result <- OSOther
		return
	}
	// Get Linux distribution
	output, err = exec.Command("cat", "/etc/os-release").Output()
	if err != nil {
		fmt.Println("Error getting Linux distribution:", err)
		result <- OSOther
		return
	}
	entries := strings.Split(string(output), "\n")
	for _, entry := range entries {
		if strings.HasPrefix(entry, "ID=") {
			os = strings.Split(entry, "=")[1]
			break
		}
	}
	switch os {
	case "arch":
		result <- OSArch
	case "debian":
		result <- OSDebian
	case "fedora":
		result <- OSFedora
	case "macos":
		result <- OSMac
	case "ubuntu":
		result <- OSUbuntu
	default:
		result <- OSLinux
	}
}
