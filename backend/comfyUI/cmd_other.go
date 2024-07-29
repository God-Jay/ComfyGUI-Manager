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
	var err error
	err = cmd.Process.Signal(os.Interrupt)
	if err != nil {
		return err
	}
	return cmd.Wait()
}
