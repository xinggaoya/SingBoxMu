package model

/**
  @author: XingGao
  @date: 2024/8/7
**/

type SingBoxConfig struct {
	Log struct {
		Level     string `json:"level"`
		Timestamp bool   `json:"timestamp"`
	} `json:"log"`
	Dns struct {
		Servers []struct {
			Tag             string `json:"tag"`
			Address         string `json:"address"`
			AddressResolver string `json:"address_resolver,omitempty"`
			Strategy        string `json:"strategy,omitempty"`
			Detour          string `json:"detour"`
		} `json:"servers"`
		Rules []struct {
			Outbound string `json:"outbound,omitempty"`
			Server   string `json:"server"`
			RuleSet  string `json:"rule_set,omitempty"`
		} `json:"rules"`
		Final string `json:"final"`
	} `json:"dns"`
	Route struct {
		RuleSet []struct {
			Tag            string `json:"tag"`
			Type           string `json:"type"`
			Format         string `json:"format"`
			Url            string `json:"url"`
			DownloadDetour string `json:"download_detour"`
		} `json:"rule_set"`
		Rules []struct {
			Protocol string      `json:"protocol,omitempty"`
			Outbound string      `json:"outbound"`
			Port     interface{} `json:"port,omitempty"`
			Network  string      `json:"network,omitempty"`
			Type     string      `json:"type,omitempty"`
			Mode     string      `json:"mode,omitempty"`
			Rules    []struct {
				RuleSet string `json:"rule_set"`
				Invert  bool   `json:"invert,omitempty"`
			} `json:"rules,omitempty"`
			RuleSet     string `json:"rule_set,omitempty"`
			IpIsPrivate bool   `json:"ip_is_private,omitempty"`
		} `json:"rules"`
		Final               string `json:"final"`
		AutoDetectInterface bool   `json:"auto_detect_interface"`
	} `json:"route"`
	Inbounds  []Inbounds `json:"inbounds"`
	Outbounds []struct {
		Type       string `json:"type"`
		Tag        string `json:"tag"`
		Server     string `json:"server,omitempty"`
		ServerPort int    `json:"server_port,omitempty"`
		Uuid       string `json:"uuid,omitempty"`
		Transport  struct {
			Path    string `json:"path,omitempty"`
			Type    string `json:"type,omitempty"`
			Headers struct {
				Host string `json:"Host,omitempty"`
			} `json:"headers,omitempty"`
		} `json:"transport,omitempty"`
		Tls struct {
			Enabled  bool `json:"enabled,omitempty"`
			Insecure bool `json:"insecure,omitempty"`
		} `json:"tls,omitempty"`
		Network     string `json:"network,omitempty"`
		TcpFastOpen bool   `json:"tcp_fast_open,omitempty"`
	} `json:"outbounds"`
	Experimental struct {
		CacheFile struct {
			Enabled bool   `json:"enabled"`
			Path    string `json:"path"`
		} `json:"cache_file"`
	} `json:"experimental"`
}

type Inbounds struct {
	Type                     string   `json:"type"`
	Tag                      string   `json:"tag,omitempty"`
	Address                  []string `json:"address,omitempty"`
	Mtu                      int      `json:"mtu,omitempty"`
	Inet4Address             []string `json:"inet4_address,omitempty"`
	Inet6Address             []string `json:"inet6_address,omitempty"`
	AutoRoute                bool     `json:"auto_route,omitempty"`
	StrictRoute              bool     `json:"strict_route,omitempty"`
	Stack                    string   `json:"stack,omitempty"`
	Sniff                    bool     `json:"sniff,omitempty"`
	SniffOverrideDestination bool     `json:"sniff_override_destination,omitempty"`
	Listen                   string   `json:"listen,omitempty"`
	ListenPort               int      `json:"listen_port,omitempty"`
	SetSystemProxy           bool     `json:"set_system_proxy,omitempty"`
}
