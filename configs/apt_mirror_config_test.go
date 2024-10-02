package configs

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"os"
	"proxynd/helpers"
	"testing"
)

func TestAptMirrorConfig_ReadConfig(t *testing.T) {
	os.Setenv("CONFIG_DIR", "../sample-conf/")
	cfg := AptMirrorConfig{}
	cfg.ReadConfig()
	fmt.Println(cfg)
	assert.Equal(t, cfg.Path, "mirror/apt")
	assert.Equal(t, len(cfg.Mirrors.Ubuntu), 2)
	assert.Equal(t, len(cfg.Mirrors.Debian), 2)
	fmt.Println(helpers.YamlToString(cfg))
}
