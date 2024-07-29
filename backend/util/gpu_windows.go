//go:build windows
// +build windows

package util

import "os/exec"

func gpuCmd() *exec.Cmd {
	return exec.Command("wmic", "path", "win32_VideoController", "get", "name")
}
