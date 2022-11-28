package di

import (
	"github.com/gin-gonic/gin"
)

func GetHttpHandler(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	cakes := r.Group("/cakes")
	cakes.GET("/", GetCakeHandler().GetCakes)
	cakes.GET("/:id", GetCakeHandler().GetCakeByID)
	cakes.POST("/", GetCakeHandler().Store)
	cakes.PUT("/:id", GetCakeHandler().Update)
	cakes.DELETE("/:id", GetCakeHandler().Delete)
	cakes.POST("/upload", GetCakeHandler().UploadImage)
}
