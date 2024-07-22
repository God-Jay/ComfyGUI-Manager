//go:build darwin
// +build darwin

package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func init() {
	bashCommands := `
        # Load Bash configuration files
        if [ -f ~/.bash_profile ]; then
            source ~/.bash_profile
        fi

        if [ -f ~/.bashrc ]; then
            source ~/.bashrc
        fi

        # Load Zsh configuration files
        if [ -f ~/.zshenv ]; then
            source ~/.zshenv
        fi

        if [ -f ~/.zprofile ]; then
            source ~/.zprofile
        fi

        if [ -f ~/.zshrc ]; then
            source ~/.zshrc
        fi

        # Print the final PATH
        echo $PATH
    `
	cmd := exec.Command("bash", "-c", bashCommands)
	output, err := cmd.Output()
	if err != nil {
		log.Println("set path err:", err)
	} else {
		path := strings.TrimSpace(string(output))
		os.Setenv("PATH", path)
		log.Println("set path:", path)
	}
}
