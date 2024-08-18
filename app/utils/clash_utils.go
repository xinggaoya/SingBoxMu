package utils

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
	"io"
	"log"
	"log/slog"
	"net/http"
	"time"
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
	req.Header.Set("Authorization", "Bearer CDBEA571-B1FC-4AEF-A8BE-1E5B3D11B9C2")
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
		// 重试
		time.Sleep(1 * time.Second)
		c.GetLogs()
		return
	}
	defer request.Body.Close()

	fmt.Printf("获取日志信息成功")

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
		time.Sleep(1 * time.Second)
		c.GetTraffic()
		return
	}
	defer request.Body.Close()

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
	request, err := c.DoRequest("GET", "/memory", nil)
	if err != nil {
		slog.Error("获取内存信息失败", "err", err)
		time.Sleep(1 * time.Second)
		c.GetMemory()
		return
	}
	defer request.Body.Close()

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

// ReloadConfig 重载配置
func (c *ClashClient) ReloadConfig(force bool) error {
	body := map[string]interface{}{
		"force": force,
	}
	request, err := c.DoRequest("PUT", "/configs/reload", body)
	if err != nil {
		slog.Error("重载配置失败", "err", err)
		return err
	}
	defer request.Body.Close()

	return nil
}

// GetProxies 获取所有代理
func (c *ClashClient) GetProxies() (string, error) {
	request, err := c.DoRequest("GET", "/proxies", nil)
	if err != nil {
		fmt.Printf("获取代理失败")
		return "", err
	}

	defer request.Body.Close()

	// Define a buffer size for reading chunks of data.
	bufferSize := 4096

	var buffer []byte
	// Read the response body chunk by chunk and process each chunk.
	for {
		chunk := make([]byte, bufferSize)
		reader := io.LimitReader(request.Body, int64(bufferSize)) // Re-create LimitReader for each iteration
		n, err := reader.Read(chunk)
		if n > 0 {
			buffer = append(buffer, chunk[:n]...)
		}
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error reading from stream: %v", err)
		}
	}
	return string(buffer), nil
}

// SwitchProxy 切换代理
func (c *ClashClient) SwitchProxy(group, name string) error {

	request, err := c.DoRequest("PUT", "/proxies/"+group, struct {
		Name string `json:"name"`
	}{
		Name: name,
	})
	if err != nil {
		slog.Error("切换代理失败", "err", err)
		return err
	}
	defer request.Body.Close()

	if request.StatusCode != 204 {
		slog.Error("切换代理失败", "status", request.Status)
		return fmt.Errorf("切换代理失败")
	}

	return nil
}

// ReloadConfigFile 重新加载配置文件
func (c *ClashClient) ReloadConfigFile() error {
	request, err := c.DoRequest("PUT", "/configs/reload", nil)
	if err != nil {
		slog.Error("重新加载配置文件失败", "err", err)
		return err
	}
	defer request.Body.Close()

	if request.StatusCode != 204 {
		slog.Error("重新加载配置文件失败", "status", request.Status)
		return fmt.Errorf("重新加载配置文件失败")
	}
	return nil
}
