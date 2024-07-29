//go:build darwin
// +build darwin

package util

import "os/exec"

func gpuCmd() *exec.Cmd {
	return exec.Command("system_profiler", "SPDisplaysDataType")
}
