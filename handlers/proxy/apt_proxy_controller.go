package proxy

import (
	"deproxy/configs"
	"deproxy/helpers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/go-homedir"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func Apt(c *gin.Context) {
	log.Printf("Access proxy apt\n")
	requestPath := c.Param("path")
	config := configs.AptProxyConfig{}
	config.ReadConfig("config/proxy-apt.yaml")

	basedir := os.Getenv("BASE_DIR")
	if basedir == "" {
		var err error
		basedir, err = homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}
	}

	// Create the file
	filefullpath := path.Join(basedir, os.Getenv("CACHE_DIR")+requestPath)
	filename := filepath.Base(filefullpath)
	if _, err := os.Stat(filefullpath); os.IsNotExist(err) {
		dirpath := filepath.Dir(filefullpath)
		os.MkdirAll(dirpath, os.ModePerm)
		out, err := os.Create(filefullpath)
		if err != nil {
			panic(err)
			return
		}
		defer out.Close()

		// Get the data
		for s, server := range config.Proxies.Ubuntu {
			fmt.Printf("for moon %s\n", s)
			resp, err := http.Get(helpers.JoinURL(server.Url, requestPath))
			if err != nil {
				log.Fatal(err)
				return
			}
			//fmt.Println(resp.Header)
			fmt.Println(resp.StatusCode)
			// Writer the body to file
			_, err = io.Copy(out, resp.Body)
			if err != nil {
				log.Fatal(err)
				return
			}
			resp.Body.Close()
			break
		}
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.File(filefullpath)
}
