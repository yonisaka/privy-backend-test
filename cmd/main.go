package main

import (
	"net/http"
	"path"
	"path/filepath"
	"privy-backend-test/internal/di"
	"privy-backend-test/internal/helpers"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	di.GetHttpHandler(r)

	getImageFromURL(r)
	r.Run(helpers.GoDotEnvVariable("BASE_URL"))
}

func getImageFromURL(r *gin.Engine) {
	r.GET("/image/:modul/*filename", func(c *gin.Context) {
		modul := c.Param("modul")
		filename := c.Param("filename")
		if strings.TrimPrefix(filename, "/") == "" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		fullname := filepath.Join(modul, filepath.FromSlash(path.Clean("/"+filename)))
		c.File("./storage/images/" + fullname)
	})
}
