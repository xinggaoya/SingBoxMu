package utils

import (
	"changeme/app/consts"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

/**
  @author: XingGao
  @date: 2024/8/10
**/

type AppUtils struct{}

func NewAppUtils() *AppUtils {
	return &AppUtils{}
}

// GetAppDir 获取程序目录
func (g *AppUtils) GetAppDir(name ...string) (string, error) {
	executable, err := os.Executable()
	if err != nil {
		return "", err
	}
	// 只保留到目录
	path := filepath.Dir(executable)
	// 拼接
	for _, s := range name {
		path = filepath.Join(path, s)
	}
	return path, nil
}

// SetSystemProxy 设置windows系统代理
func (g *AppUtils) SetSystemProxy(proxyAddress string, enable bool) error {
	key, _, err := registry.CreateKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer key.Close()

	err = key.SetStringValue("ProxyServer", proxyAddress)
	if err != nil {
		return err
	}
	num := uint32(0)
	if !enable {
		num = 0
	} else {
		num = 1
	}

	err = key.SetDWordValue("ProxyEnable", num)
	if err != nil {
		return err
	}

	return nil
}

// IsRunningAsAdmin 检查是否具有管理员权限
func (g *AppUtils) IsRunningAsAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}

// RunAsAdmin 以管理员权限重新启动程序
func (g *AppUtils) RunAsAdmin() error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	cmd := exec.Command("powershell", "Start-Process", exe, "-Verb", "runAs")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	err = cmd.Start()
	if err != nil {
		return err
	}
	cmd.Wait()

	os.Exit(0)
	return nil
}

// RegisterStartup 注册开机自启
func (g *AppUtils) RegisterStartup() error {
	key, _, err := registry.CreateKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer key.Close()

	exe, err := os.Executable()

	exe = fmt.Sprintf("%s %s", exe, "-hide")
	if err != nil {
		return err
	}

	err = key.SetStringValue(consts.TaskName, exe)
	if err != nil {
		return err
	}
	return nil
}

// UnregisterStartup 取消开机自启
func (g *AppUtils) UnregisterStartup() error {
	key, _, err := registry.CreateKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer key.Close()

	err = key.DeleteValue(consts.TaskName)
	if err != nil {
		return err
	}
	return nil
}
