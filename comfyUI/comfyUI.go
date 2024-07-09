package comfyUI

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

type ComfyUI struct {
	ctx         context.Context
	ComfyUIPath string
	cmd         *exec.Cmd
}

func NewComfyUI() *ComfyUI {
	return &ComfyUI{}
}

func (c *ComfyUI) StartUp(ctx context.Context) {
	c.ctx = ctx
}

func (c *ComfyUI) CheckIsServerRunning() bool {
	return isPortInUse(8188)
}

func (c *ComfyUI) StartServer(comfyUIPath string) bool {
	// TODO delete
	cmd := exec.Command("python", "main.py", "--force-upcast-attention")
	cmd.Dir = comfyUIPath

	log.Println(cmd.String())
	err := cmd.Start()
	if err != nil {
		log.Println("start server err", err)
		return false
	}
	for i := 0; i < 30; i++ {
		time.Sleep(time.Second)
		if isPortInUse(8188) {
			c.cmd = cmd
			return true
		}
	}
	cmd.Process.Kill()
	log.Println("start server timeout after 30s")
	return false
}

func (c *ComfyUI) StopServer() bool {
	if c.cmd == nil {
		log.Println("stop server failed, server is not running by ComfyGui")
		return false
	}

	err := c.cmd.Process.Signal(os.Interrupt)
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
