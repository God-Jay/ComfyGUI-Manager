//go:build windows
// +build windows

package comfyUI

import (
	"os/exec"
	"syscall"
)

func hideWindow(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
}

// TODO ctrl+c ?
func stopCmd(cmd *exec.Cmd) error {
	return cmd.Process.Kill()
}
