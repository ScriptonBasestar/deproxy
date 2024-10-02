package proxy

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"proxynd/configs"
	"proxynd/helpers"
)

func DockerProxy(c *gin.Context) {
	log.Printf("Access proxy docker\n")

	imageName := c.Param("imageName")
	tag := c.Param("tag")

	// nolint:godox fixme di??
	storageDir := helpers.GetStorageDir()
	globalConfig := configs.GlobalConfig{}
	globalConfig.ReadConfig()
	config := configs.DockerProxyConfig{}
	config.ReadConfig()

	// Create the file path
	filefullpath := path.Join(storageDir, config.Path, imageName, tag)
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
			//log.Fatal(err)
			//return
		}
		defer func() {
			if err := out.Close(); err != nil {
				log.Printf("Error closing file: %v", err)
			}
		}()

		// Get the data
		//proxy := config.Proxies["docker"]
		for s, server := range config.Proxies {
			log.Printf("Trying server %d: %s\n", s, server.Url)
			result, err := url.JoinPath(server.Url, imageName, tag)
			if err != nil {
				log.Fatal(err)
				return
			}
			resp, err := http.Get(result)
			if err != nil {
				log.Fatal(err)
				return
			}
			if resp.StatusCode != http.StatusOK {
				log.Printf("Server %s returned status %d\n", server.Url, resp.StatusCode)
				continue
			}
			// Write the body to file
			_, err = io.Copy(out, resp.Body)
			if err != nil {
				log.Fatal(err)
				return
			}
			err = resp.Body.Close()
			if err != nil {
				log.Fatal(err)
				return
			}
			break
		}
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.File(filefullpath)
}

/*
get docker image layer
GET /v2/<repository>/blobs/<digest> HTTP/1.1
Host: <registry>
User-Agent: docker/20.10.7
Authorization: Bearer <token>
Accept: application/vnd.docker.distribution.manifest.v2+json
*/
