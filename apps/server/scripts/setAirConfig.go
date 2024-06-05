package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	var configFile string

	// Determine the operating system
	if runtime.GOOS == "windows" {
		configFile = ".air.windows.toml"
	} else {
		configFile = ".air.unix.toml"
	}

	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	dirPath := filepath.Dir(exePath)

	// Construct the full path to the configuration file
	fullPath := filepath.Join(dirPath, configFile)

	// Build the command to execute Air with the correct configuration file
	cmd := exec.Command("air", "-c", fullPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Execute the command
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
