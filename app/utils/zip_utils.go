package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

/**
  @author: XingGao
  @date: 2024/8/7
**/

// Unzip 解压ZIP文件到指定目录
func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %w", err)
	}
	defer func(r *zip.ReadCloser) {
		err = r.Close()
		if err != nil {
			log.Printf("关闭zip文件失败: %v", err)
		}
		// 删除zip文件
		err = os.Remove(src)
		if err != nil {
			log.Printf("failed to remove zip file: %v", err)
		}
	}(r)

	for _, f := range r.File {
		err = extractAndWriteFile(f, dest)
		if err != nil {
			return fmt.Errorf("failed to extract file %s: %w", f.Name, err)
		}
	}

	return nil
}

// extractAndWriteFile 从zip.Reader中提取单个文件并写入到磁盘
func extractAndWriteFile(file *zip.File, dest string) error {
	// 仅处理文件，忽略目录
	if file.FileInfo().IsDir() {
		return nil
	}

	// 构造文件的完整路径
	path := filepath.Join(dest, filepath.Base(file.Name))

	// 确保目标目录存在
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	srcFile, err := file.Open()
	if err != nil {
		return err
	}
	defer func(srcFile io.ReadCloser) {
		err = srcFile.Close()
		if err != nil {
			log.Printf("关闭文件失败: %v", err)
		}
	}(srcFile)

	outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, srcFile)
	return err
}
