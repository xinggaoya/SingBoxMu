package service

import (
	"changeme/app/model/response"
	"changeme/app/utils"
)

/**
  @author: XingGao
  @date: 2024/8/15
**/

type ClashService struct{}

// GetProxies 获取所有代理
func (s *ClashService) GetProxies() response.ResInfo {
	clash := utils.NewClashClient()
	proxies, err := clash.GetProxies()
	if err != nil {
		return response.Error("获取代理失败")
	}
	return response.Success(proxies)
}

// SwitchProxy 更换代理
func (s *ClashService) SwitchProxy(group, name string) response.ResInfo {
	clash := utils.NewClashClient()
	err := clash.SwitchProxy(group, name)
	if err != nil {
		return response.Error("更换代理失败")
	}
	return response.Success("更换代理成功")
}

// ReloadConfig 重载配置
func (s *ClashService) ReloadConfig(force bool) response.ResInfo {
	clash := utils.NewClashClient()
	err := clash.ReloadConfig(force)
	if err != nil {
		return response.Error("重载配置失败")
	}
	return response.Success("重载配置成功")
}

// GetVersion 获取内核版本
func (s *ClashService) GetVersion() response.ResInfo {
	clash := utils.NewClashClient()
	version, err := clash.GetVersion()
	if err != nil {
		return response.Error("内核未启动")
	}
	return response.Success(version)
}
