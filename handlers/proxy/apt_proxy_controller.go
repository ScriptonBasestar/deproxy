package proxy

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"proxynd/configs"
	"proxynd/helpers"
)

func AptProxy(c *gin.Context) {
	log.Printf("Access proxy apt\n")

	pathOs := c.Param("osType")
	requestPath := c.Param("requestPath")

	// nolint:godox fixme di??
	storageDir := helpers.GetStorageDir()
	globalConfig := configs.GlobalConfig{}
	globalConfig.ReadConfig()
	config := configs.AptProxyConfig{}
	config.ReadConfig()

	//basedir := os.Getenv("STORAGE_DIR")
	//if basedir == "" {
	//	var err error
	//	basedir, err = homedir.Dir()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	// Create the file
	filefullpath := path.Join(storageDir, config.Path, requestPath)
	filename := filepath.Base(filefullpath)
	if _, err := os.Stat(filefullpath); os.IsNotExist(err) {
		dirpath := filepath.Dir(filefullpath)
		err = os.MkdirAll(dirpath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
			return
		}
		out, err := os.Create(filefullpath)
		if err != nil {
			panic(err)
			//return
		}
		defer out.Close()

		// Get the data
		proxy := config.Proxies[pathOs]
		for s, server := range proxy {
			fmt.Printf("for moon %d\n", s)
			resp, err := http.Get(helpers.JoinURL(server.URL, requestPath))
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
			//break
		}
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.File(filefullpath)
}
