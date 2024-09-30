package configs

import (
	"path"
	"proxynd/helpers"
)

type DockerProxyServer struct {
	Id          string `yaml:"id,omitempty"`
	Name        string `yaml:"name"`
	Url         string `yaml:"url,omitempty"`
	Description string `yaml:"description"`
}

type DockerProxyConfig struct {
	Path    string              `yaml:"path,omitempty"`
	Proxies []DockerProxyServer `yaml:"proxies"`
}

func (cfg *DockerProxyConfig) ConfigExists() bool {
	confDir := helpers.GetConfigDir()
	return helpers.FileExists(path.Join(confDir, "docker-proxy.yaml"))
}

func (cfg *DockerProxyConfig) ReadConfig() {
	confDir := helpers.GetConfigDir()
	helpers.ReadYaml(path.Join(confDir, "docker-proxy.yaml"), cfg)
}
