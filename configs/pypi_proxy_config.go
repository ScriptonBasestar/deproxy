package configs

import (
	"path"
	"proxynd/helpers"
)

type PypiProxyServer struct {
	Id          string `yaml:"id,omitempty"`
	Name        string `yaml:"name"`
	Url         string `yaml:"url,omitempty"`
	Description string `yaml:"description"`
}

type PypiProxyConfig struct {
	Path    string            `yaml:"path,omitempty"`
	Proxies []PypiProxyServer `yaml:"proxies"`
}

func (cfg *PypiProxyConfig) ConfigExists() bool {
	confDir := helpers.GetConfigDir()
	return helpers.FileExists(path.Join(confDir, "pypi-proxy.yaml"))
}

func (cfg *PypiProxyConfig) ReadConfig() {
	confDir := helpers.GetConfigDir()
	helpers.ReadYaml(path.Join(confDir, "pypi-proxy.yaml"), cfg)
}
