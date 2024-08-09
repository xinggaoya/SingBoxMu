package service

import (
	"changeme/app/model"
	"changeme/app/utils"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"time"
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
		slog.Error("Failed to decode latest kernel version", "error", err)
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

// DownloadSubscription 下载订阅
func (g *AppService) DownloadSubscription(url string) {
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("Failed to get subscription", "error", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed to read subscription", "error", err)
		return
	}

	var outInfo model.DownloadSingBoxConfig
	err = json.Unmarshal(body, &outInfo)
	if err != nil {
		slog.Error("Failed to unmarshal subscription", "error", err)
		return
	}

	// 获取模版
	var inInfo model.SingBoxConfig
	err = json.Unmarshal(model.SingBoxConfigTemplate, &inInfo)
	if err != nil {
		slog.Error("Failed to unmarshal template", "error", err)
		return
	}
	var newInfo model.DownloadSingBoxConfig
	for _, item := range outInfo.Outbounds {
		// 只需要url不为空
		if item.Server != "" {
			inInfo.Outbounds = append(inInfo.Outbounds, item)
			if newInfo.Outbounds == nil {
				newInfo.Outbounds = append(newInfo.Outbounds, item)
			}
		}
	}
	inInfo.Outbounds[0].Outbounds[1] = newInfo.Outbounds[0].Tag
	inInfo.Outbounds[1].Outbounds[0] = newInfo.Outbounds[0].Tag

	// json转字符
	data, err := json.Marshal(inInfo)
	err = os.WriteFile("./sing-box/config.json", data, 0644)
	if err != nil {
		slog.Error("Failed to write config", "error", err)
		return
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

// StartCommand 启动内核
func (g *AppService) StartCommand() {
	cmd := exec.Command("./sing-box/sing-box", "run", "-D", "./sing-box")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

	err := cmd.Start()

	if err != nil {
		fmt.Printf("Error starting command: %v\n", err)
	}

	slog.Info("Started command with PID %d\n", cmd.Process.Pid)
	singBox = cmd

	go ListenKernelInfo()
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
		// Wait for the process to exit gracefully
		if err = singBox.Wait(); err != nil {
			slog.Error("Error waiting for command to stop", "error", err)
			return
		}

		singBox = nil
		slog.Info("Stopped command")
	}
}

// GetVersion 获取内核版本
func (g *AppService) GetVersion() string {
	clash := utils.NewClashClient()
	return clash.GetVersion()
}

// ListenKernelInfo 监听日志
func ListenKernelInfo() {
	// 延迟10s
	time.Sleep(5 * time.Second)
	clash := utils.NewClashClient()
	go clash.GetLogs()
	go clash.GetMemory()
	go clash.GetTraffic()
}

// RestartCommand Function to restart a command
func (g *AppService) RestartCommand() {
	g.StopCommand()
	g.StartCommand()
}

// SetAutoStart 设置开机自启
func (g *AppService) SetAutoStart() {
	// 获取当前系统类型
	task := utils.NewTaskUtils()
	if err := task.CreateTask(); err != nil {
		slog.Error("Failed to create task", "error", err)
		return
	}
}

// RemoveAutoStart 移除开机自启
func (g *AppService) RemoveAutoStart() {
	// 获取当前系统类型
	task := utils.NewTaskUtils()
	if err := task.DeleteTask(); err != nil {
		slog.Error("Failed to delete task", "error", err)
		return
	}
}
