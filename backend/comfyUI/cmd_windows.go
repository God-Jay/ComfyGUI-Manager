//go:build windows
// +build windows

package comfyUI

import (
	"errors"
	"log"
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
	var err error
	err = cmd.Process.Kill()
	if err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) && exitErr.ExitCode() != 0 {
			log.Println("Process was killed and exited with status:", exitErr.ExitCode())
			return nil
		}
		return err
	}

	return nil
}
