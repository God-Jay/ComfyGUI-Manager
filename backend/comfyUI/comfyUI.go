package comfyUI

import (
	"comfygui-manager/backend/store"
	"comfygui-manager/backend/util"
	"context"
	"fmt"
	"log"
	"net"
	"os/exec"
	"time"
)

type ComfyUI struct {
	ctx context.Context
	cmd *exec.Cmd
}

func NewComfyUI() *ComfyUI {
	return &ComfyUI{}
}

func (c *ComfyUI) StartUp(ctx context.Context) {
	c.ctx = ctx
}

func (c *ComfyUI) Shutdown() {
	if c.cmd != nil {
		stopCmd(c.cmd)
	}
}

func (c *ComfyUI) CheckIsServerRunning() bool {
	return isPortInUse(8188)
}

func (c *ComfyUI) StartServer() bool {
	// TODO change args
	cmd := exec.Command("python", "main.py")

	switch util.GetGpu() {
	case util.APPLE:
		cmd.Args = append(cmd.Args, "--force-upcast-attention")
	case util.AMD:
		cmd.Args = append(cmd.Args, "--directml")
	}

	cmd.Dir = store.ComfyUIPath

	hideWindow(cmd)

	log.Println(cmd.String())
	err := cmd.Start()
	if err != nil {
		log.Println("start server err", err)
		return false
	}
	for i := 0; i < 120; i++ {
		time.Sleep(time.Second)
		if isPortInUse(8188) {
			c.cmd = cmd
			return true
		}
	}
	stopCmd(cmd)
	log.Println("start server timeout after 120s")
	return false
}

func (c *ComfyUI) StopServer() bool {
	if c.cmd == nil {
		log.Println("stop server failed, server is not running by ComfyGui")
		return false
	}

	err := stopCmd(c.cmd)
	if err != nil {
		log.Println("stop server err", err)
		return false
	}
	for i := 0; i < 30; i++ {
		time.Sleep(time.Second)
		if !isPortInUse(8188) {
			c.cmd = nil
			return true
		}
	}
	log.Println("stop server failed after 30s")
	return false
}

func isPortInUse(port int) bool {
	// TODO
	address := fmt.Sprintf("127.0.0.1:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return true
	}
	defer listener.Close()
	return false
}
