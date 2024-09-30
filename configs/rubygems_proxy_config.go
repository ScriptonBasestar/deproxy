package configs

import (
	"path"
	"proxynd/helpers"
)

type RubygemsProxyServer struct {
	Id          string `yaml:"id,omitempty"`
	Name        string `yaml:"name"`
	Url         string `yaml:"url,omitempty"`
	Description string `yaml:"description"`
}

type RubygemsProxyConfig struct {
	Path    string                `yaml:"path,omitempty"`
	Proxies []RubygemsProxyServer `yaml:"proxies"`
}

func (cfg *RubygemsProxyConfig) ConfigExists() bool {
	confDir := helpers.GetConfigDir()
	return helpers.FileExists(path.Join(confDir, "rubygems-proxy.yaml"))
}

func (cfg *RubygemsProxyConfig) ReadConfig() {
	confDir := helpers.GetConfigDir()
	helpers.ReadYaml(path.Join(confDir, "rubygems-proxy.yaml"), cfg)
}
