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

	content, err := os.ReadFile("./sing-box/config.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &config)
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

	content, err := os.ReadFile("./sing-box/config.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		return err
	}

	config.Inbounds = []model.Inbounds{
		{
			Type:                     "tun",
			Tag:                      "tun-in",
			Inet4Address:             "172.18.0.1/30",
			Inet6Address:             "fdfe:dcba:9876::1/126",
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
