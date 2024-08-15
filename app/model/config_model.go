package model

import _ "embed"

/**
  @author: XingGao
  @date: 2024/8/7
**/

//go:embed config.json
var SingBoxConfigTemplate []byte

type SingBoxConfig struct {
	Log struct {
		Level     string `json:"level"`
		Timestamp bool   `json:"timestamp"`
	} `json:"log"`
	Experimental struct {
		ClashApi struct {
			ExternalController    string `json:"external_controller"`
			ExternalUi            string `json:"external_ui"`
			Secret                string `json:"secret"`
			ExternalUiDownloadUrl string `json:"external_ui_download_url"`
			DefaultMode           string `json:"default_mode"`
		} `json:"clash_api"`
		CacheFile struct {
			Enabled     bool `json:"enabled"`
			StoreFakeip bool `json:"store_fakeip"`
			StoreRdrc   bool `json:"store_rdrc"`
		} `json:"cache_file"`
	} `json:"experimental"`
	Inbounds  []Inbounds         `json:"inbounds"`
	Outbounds []SingBoxOutbounds `json:"outbounds"`
	Route     struct {
		RuleSet []struct {
			Tag            string `json:"tag"`
			Type           string `json:"type"`
			Format         string `json:"format"`
			Url            string `json:"url"`
			DownloadDetour string `json:"download_detour"`
		} `json:"rule_set"`
		Rules []struct {
			Protocol    string `json:"protocol,omitempty"`
			Port        int    `json:"port,omitempty"`
			Outbound    string `json:"outbound"`
			ClashMode   string `json:"clash_mode,omitempty"`
			Network     string `json:"network,omitempty"`
			RuleSet     string `json:"rule_set,omitempty"`
			IpIsPrivate bool   `json:"ip_is_private,omitempty"`
		} `json:"rules"`
		Final               string `json:"final"`
		AutoDetectInterface bool   `json:"auto_detect_interface"`
	} `json:"route"`
	Dns struct {
		Servers []struct {
			Tag             string `json:"tag"`
			Address         string `json:"address"`
			AddressResolver string `json:"address_resolver,omitempty"`
			Detour          string `json:"detour,omitempty"`
		} `json:"servers"`
		DisableCache     bool `json:"disable_cache"`
		DisableExpire    bool `json:"disable_expire"`
		IndependentCache bool `json:"independent_cache"`
		Rules            []struct {
			Outbound     string `json:"outbound,omitempty"`
			DisableCache bool   `json:"disable_cache,omitempty"`
			Server       string `json:"server"`
			ClashMode    string `json:"clash_mode,omitempty"`
			RuleSet      string `json:"rule_set,omitempty"`
		} `json:"rules"`
		Fakeip struct {
			Enabled    bool   `json:"enabled"`
			Inet4Range string `json:"inet4_range"`
			Inet6Range string `json:"inet6_range"`
		} `json:"fakeip"`
		Final    string `json:"final"`
		Strategy string `json:"strategy"`
	} `json:"dns"`
}

type DownloadSingBoxConfig struct {
	Outbounds []SingBoxOutbounds `json:"outbounds"`
}

type SingBoxOutbounds struct {
	Tag        string   `json:"tag"`
	Type       string   `json:"type"`
	Outbounds  []string `json:"outbounds,omitempty"`
	Url        string   `json:"url,omitempty"`
	Password   string   `json:"password,omitempty"`
	Interval   string   `json:"interval,omitempty"`
	Tolerance  int      `json:"tolerance,omitempty"`
	Server     string   `json:"server,omitempty"`
	ServerPort int      `json:"server_port,omitempty"`
	Uuid       string   `json:"uuid,omitempty"`
	Transport  struct {
		Path    string `json:"path,omitempty"`
		Type    string `json:"type,omitempty"`
		Headers struct {
			Host string `json:"Host,omitempty"`
		} `json:"headers,omitempty"`
	} `json:"transport,omitempty"`
	Tls         SingBoxInboundsTls `json:"tls,omitempty"`
	Network     string             `json:"network,omitempty"`
	TcpFastOpen bool               `json:"tcp_fast_open,omitempty"`
}

type SingBoxInboundsTls struct {
	Enabled    bool   `json:"enabled,omitempty"`
	ServerName string `json:"server_name,omitempty"`
	Insecure   any    `json:"insecure,omitempty"`
}

type Inbounds struct {
	Type                     string   `json:"type"`
	Tag                      string   `json:"tag,omitempty"`
	Address                  []string `json:"address,omitempty"`
	Mtu                      int      `json:"mtu,omitempty"`
	Inet4Address             string   `json:"inet4_address,omitempty"`
	Inet6Address             string   `json:"inet6_address,omitempty"`
	AutoRoute                bool     `json:"auto_route,omitempty"`
	StrictRoute              bool     `json:"strict_route,omitempty"`
	Stack                    string   `json:"stack,omitempty"`
	Sniff                    bool     `json:"sniff,omitempty"`
	SniffOverrideDestination bool     `json:"sniff_override_destination,omitempty"`
	Listen                   string   `json:"listen,omitempty"`
	ListenPort               int      `json:"listen_port,omitempty"`
	SetSystemProxy           bool     `json:"set_system_proxy,omitempty"`
}
