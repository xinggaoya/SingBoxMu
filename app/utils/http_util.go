package utils

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

/**
  @author: XingGao
  @date: 2024/8/7
**/

// DownloadFile 下载文件并保存到本地
func DownloadFile(url, savePath string) error {
	req, err := http.Get(url)

	if err != nil {
		return err
	}
	defer req.Body.Close()

	if err = os.MkdirAll(filepath.Dir(savePath), 0755); err != nil {
		return err
	}

	file, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, req.Body)

	return err
}
