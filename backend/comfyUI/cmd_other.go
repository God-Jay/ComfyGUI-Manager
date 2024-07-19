//go:build darwin || linux
// +build darwin linux

package comfyUI

import (
	"os"
	"os/exec"
)

func hideWindow(cmd *exec.Cmd) {
}

func stopCmd(cmd *exec.Cmd) error {
	return cmd.Process.Signal(os.Interrupt)
}
