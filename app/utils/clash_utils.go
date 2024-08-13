package utils

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
	"io"
	"log/slog"
	"net/http"
)

/**
  @author: XingGao
  @date: 2024/8/8
**/

// ClashClient 客户端
type ClashClient struct {
	client  *http.Client
	baseURL string
}

// NewClashClient 创建一个新的 HTTP 客户端实例
func NewClashClient() *ClashClient {
	return &ClashClient{
		client:  &http.Client{},
		baseURL: "http://127.0.0.1:20123",
	}
}

// InitRequestHeaders 初始化请求头
func (c *ClashClient) InitRequestHeaders(req *http.Request) {
	req.Header.Set("Authorization", "Bearer ID_k2am62qu")
}

// DoRequest 执行HTTP请求
func (c *ClashClient) DoRequest(method, path string, body interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, path)
	reqBody, _ := json.Marshal(body)

	req, err := http.NewRequestWithContext(context.Background(), method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	c.InitRequestHeaders(req)

	return c.client.Do(req)
}

// GetLogs 获取日志信息
func (c *ClashClient) GetLogs() {
	request, err := c.DoRequest("GET", "/logs", nil)
	if err != nil {
		slog.Error("获取日志信息失败", "err", err)
		return
	}
	defer request.Body.Close()

	// 检查响应状态码是否为成功
	if request.StatusCode != http.StatusOK {
		slog.Error("获取日志信息失败", "status", request.Status)
		return
	}

	// 逐行读取日志
	reader := bufio.NewReader(request.Body)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break // 读取完毕
		} else if err != nil {
			slog.Error("读取日志时发生错误", "err", err)
			break
		}

		application.Get().Events.Emit(&application.WailsEvent{
			Name: "logs",
			Data: line,
		})
	}

}

// GetTraffic 获取流量信息
func (c *ClashClient) GetTraffic() {
	request, err := c.DoRequest("GET", "/traffic", nil)
	if err != nil {
		slog.Error("获取流量信息失败", "err", err)
		return
	}
	defer request.Body.Close()
	if request.StatusCode != http.StatusOK {
		slog.Error("获取流量信息失败", "status", request.Status)
		return
	}

	reader := bufio.NewReader(request.Body)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break // 读取完毕
		} else if err != nil {
			slog.Error("读取流量信息时发生错误", "err", err)
			break
		}
		application.Get().Events.Emit(&application.WailsEvent{
			Name: "traffic",
			Data: line,
		})

	}
}

// GetMemory 获取使用内存
func (c *ClashClient) GetMemory() {
	request, err := c.DoRequest("GET", "", nil)
	if err != nil {
		slog.Error("获取内存信息失败", "err", err)
		return
	}
	defer request.Body.Close()
	if request.StatusCode != http.StatusOK {
		slog.Error("获取内存信息失败", request.Status)
		return
	}

	reader := bufio.NewReader(request.Body)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break // 读取完毕
		} else if err != nil {
			slog.Error("读取内存信息时发生错误", "err", err)
			break
		}
		application.Get().Events.Emit(&application.WailsEvent{
			Name: "memory",
			Data: line,
		})
	}
}

// GetVersion 获取版本信息
func (c *ClashClient) GetVersion() (string, error) {
	request, err := c.DoRequest("GET", "/version", nil)
	if err != nil {
		slog.Error("获取版本信息失败", "err", err)
		return "", err
	}
	defer request.Body.Close()
	if request.StatusCode != http.StatusOK {
		slog.Error("获取版本信息失败", "status", request.Status)
		return "", err
	}

	body, err := io.ReadAll(request.Body)

	if err != nil {
		slog.Error("读取版本信息时发生错误", "err", err)
		return "", err
	}
	return string(body), nil
}
