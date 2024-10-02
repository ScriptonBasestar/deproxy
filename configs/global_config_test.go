package configs

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"os"
	"proxynd/helpers"
	"testing"
)

func TestReadConfig_GlobalConfig(t *testing.T) {
	var err error
	err = os.Setenv("CONFIG_DIR", "../sample-conf/")
	if err != nil {
		return
	}
	err = os.Setenv("STORAGE_DIR", "../sample-conf/")
	if err != nil {
		return
	}
	cfg := GlobalConfig{}
	cfg.ReadConfig()
	assert.Equal(t, cfg.Cache.TTL, 3600)
	//assert.Equal(t, cfg.ConfigDir, "~/tmp/config")
	fmt.Println(helpers.YamlToString(cfg))
}
