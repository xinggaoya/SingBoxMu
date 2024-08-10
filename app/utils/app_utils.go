package utils

import (
	"golang.org/x/sys/windows/registry"
	"os"
	"path/filepath"
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
