package utils

import (
	"changeme/app/model"
	"encoding/json"
	"os"
)

/**
  @author: XingGao
  @date: 2024/8/7
**/

// updateConfig 更新配置文件
func updateConfig(config *model.SingBoxConfig) error {
	backupFile, err := os.OpenFile("./sing-box/config.bak", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer backupFile.Close()

	// 创建备份
	err = json.NewEncoder(backupFile).Encode(config)
	if err != nil {
		return err
	}

	// 写入新的配置
	file, err := os.OpenFile("./sing-box/config.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(config)
	if err != nil {
		return err
	}

	return nil
}

// SetProxy 系统代理
func SetProxy() error {
	var config model.SingBoxConfig

	// 读取文件 ./sing-box/config.json
	file, err := os.Open("./sing-box/config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return err
	}

	// 修改内容
	config.Inbounds = []model.Inbounds{
		{
			Type:           "mixed",
			Listen:         "0.0.0.0",
			ListenPort:     1081,
			SetSystemProxy: true,
		},
	}

	return updateConfig(&config)
}

// SetTun Tun虚拟代理
func SetTun() error {
	var config model.SingBoxConfig

	// 读取文件 ./sing-box/config.json
	file, err := os.Open("./sing-box/config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return err
	}

	config.Inbounds = []model.Inbounds{
		{
			Type:                     "tun",
			Tag:                      "tun-in",
			Inet4Address:             []string{"172.18.0.1/30"},
			Inet6Address:             []string{"fdfe:dcba:9876::1/126"},
			Mtu:                      1492,
			AutoRoute:                true,
			StrictRoute:              true,
			Stack:                    "system",
			Sniff:                    true,
			SniffOverrideDestination: false,
		},
	}

	return updateConfig(&config)
}
