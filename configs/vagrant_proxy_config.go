package configs

import (
	"path"
	"proxynd/helpers"
)

type VagrantProxyServer struct {
	Id          string `yaml:"id,omitempty"`
	Name        string `yaml:"name"`
	Url         string `yaml:"url,omitempty"`
	Description string `yaml:"description"`
}

type VagrantProxyConfig struct {
	Path   string                        `yaml:"path,omitempty"`
	Server map[string]VagrantProxyServer `yaml:"server"`
}

func (cfg *VagrantProxyConfig) ConfigExists() bool {
	configDir := helpers.GetConfigDir()
	return helpers.FileExists(path.Join(configDir, "vagrant-proxy.yaml"))
}

func (cfg *VagrantProxyConfig) ReadConfig() {
	configDir := helpers.GetConfigDir()
	helpers.ReadYaml(path.Join(configDir, "vagrant-proxy.yaml"), cfg)
}
