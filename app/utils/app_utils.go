package utils

import (
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
