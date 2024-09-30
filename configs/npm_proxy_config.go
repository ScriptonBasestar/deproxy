package configs

import (
	"path"
	"proxynd/helpers"
)

type NpmProxyServer struct {
	Id          string `yaml:"id,omitempty"`
	Name        string `yaml:"name"`
	Url         string `yaml:"url,omitempty"`
	Description string `yaml:"description"`
}

type NpmProxyConfig struct {
	Path    string           `yaml:"path,omitempty"`
	Proxies []NpmProxyServer `yaml:"proxies"`
}

func (cfg *NpmProxyConfig) ConfigExists() bool {
	confDir := helpers.GetConfigDir()
	return helpers.FileExists(path.Join(confDir, "npm-proxy.yaml"))
}

func (cfg *NpmProxyConfig) ReadConfig() {
	confDir := helpers.GetConfigDir()
	helpers.ReadYaml(path.Join(confDir, "npm-proxy.yaml"), cfg)
}
