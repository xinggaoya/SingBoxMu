package service

import (
	"changeme/app/model"
	"changeme/app/utils"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

type AppService struct{}

var singBox *exec.Cmd

func (g *AppService) Greet(name string) string {
	return "Hello " + name + "!"
}

// DownloadLatestKernel 下载最新内核
func (g *AppService) DownloadLatestKernel() {
	// 获取当前系统类型
	goos := runtime.GOOS
	// 获取系统架构
	goarch := runtime.GOARCH

	// 根据系统类型和架构下载对应的内核
	system := fmt.Sprintf("%s-%s.zip", goos, goarch)

	resp, err := http.Get("https://api.github.com/repos/SagerNet/sing-box/releases/latest")
	if err != nil {
		slog.Error("Failed to get latest kernel version", "error", err)
		return
	}
	defer resp.Body.Close()

	var releases model.SingBoxReleases

	err = json.NewDecoder(resp.Body).Decode(&releases)
	if err != nil {
		return
	}

	for _, info := range releases.Assets {
		if strings.Contains(info.Name, system) {
			fileName := fmt.Sprintf("./sing-box/%s", info.Name)
			err = utils.DownloadFile(info.BrowserDownloadUrl, fileName)
			if err != nil {
				slog.Error("Failed to download latest kernel version", "error", err)
				return
			}
			// 解压
			err = utils.Unzip(fileName, "./sing-box")
			if err != nil {
				slog.Error("Failed to unzip latest kernel version", "error", err)
				return
			}
			break
		}
	}

}

// ChangeProxyMode 设置代理
func (g *AppService) ChangeProxyMode(mode string) {
	if mode == "system" {
		err := utils.SetProxy()
		if err != nil {
			slog.Error("Failed to set proxy", "error", err)
			return
		}
		if singBox != nil && singBox.Process != nil {
			g.RestartCommand()
		}
	} else if mode == "tun" {
		err := utils.SetTun()
		if err != nil {
			slog.Error("Failed to set tun", "error", err)
			return
		}
		if singBox != nil && singBox.Process != nil {
			g.RestartCommand()
		}
	}
}

// StartCommand Function to start a command and return the Cmd struct
func (g *AppService) StartCommand() {
	cmd := exec.Command("./sing-box/sing-box", "run", "-D", "./sing-box")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()

	if err != nil {
		fmt.Printf("Error starting command: %v\n", err)
	}

	fmt.Printf("Started command with PID %d\n", cmd.Process.Pid)
	singBox = cmd
}

// StopCommand Function to stop a running command
func (g *AppService) StopCommand() {
	if singBox != nil && singBox.Process != nil {
		// 停止进程
		err := singBox.Process.Signal(syscall.SIGKILL)
		if err != nil {
			fmt.Printf("Error stopping command: %v\n", err)
			return
		}
		fmt.Printf("Stopped command with PID %d\n", singBox.Process.Pid)
	}
}

// RestartCommand Function to restart a command
func (g *AppService) RestartCommand() {
	g.StopCommand()
	g.StartCommand()
}
