package utils

import (
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

/**
  @author: XingGao
  @date: 2024/8/7
**/

type ProgressBar struct {
	TotalSize int64 `json:"totalSize"`
	Progress  int64 `json:"progress"`
}

// DownloadFile 下载文件并保存到本地 记录进度条
func DownloadFile(url, savePath, uuid string) error {
	req, err := http.Get(url)

	if err != nil {
		return err
	}
	defer req.Body.Close()

	// 获取 Content-Length
	totalSize := req.ContentLength

	if err = os.MkdirAll(filepath.Dir(savePath), 0755); err != nil {
		return err
	}

	file, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 保存并记录进度
	var downloadedSize int64

	// 最后一次更新
	lastUpdate := time.Now()

	for {
		buffer := make([]byte, 1024)
		n, err := req.Body.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := file.Write(buffer[:n]); err != nil {
			return err
		}
		downloadedSize += int64(n)

		if time.Since(lastUpdate) > time.Second {
			Notify(uuid, totalSize, downloadedSize)
			lastUpdate = time.Now()
		}
		if downloadedSize >= totalSize {
			Notify(uuid, totalSize, downloadedSize)
			break
		}
	}

	// 清理
	application.Get().Events.Off(uuid)
	return err
}

// Notify 通知前端
func Notify(uuid string, totalSize, downloadedSize int64) {
	info := &ProgressBar{
		TotalSize: totalSize,
		Progress:  downloadedSize,
	}
	// 转json
	jsonStr, err := json.Marshal(info)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	application.Get().Events.Emit(&application.WailsEvent{
		Name: uuid,
		Data: string(jsonStr),
	})
}
