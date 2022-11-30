package di

import (
	"crypto/rand"
	"privy-backend-test/internal/helpers"
	"privy-backend-test/internal/middleware"

	"github.com/gin-gonic/gin"
)

func GetHttpHandler(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.POST("/login", GetUserHandler().Login)

	cakes := r.Group("/cakes")
	cakes.Use(middleware.AuthMiddleware())
	cakes.GET("/", GetCakeHandler().GetCakes)
	cakes.GET("/:id", GetCakeHandler().GetCakeByID)
	cakes.POST("/", GetCakeHandler().Store)
	cakes.PUT("/:id", GetCakeHandler().Update)
	cakes.DELETE("/:id", GetCakeHandler().Delete)
	cakes.POST("/upload", GetCakeHandler().UploadImage)

	r.GET("/generate-aes", func(ctx *gin.Context) {
		bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
		if _, err := rand.Read(bytes); err != nil {
			panic(err.Error())
		}

		key := helpers.ConvertToString(bytes)
		ctx.JSON(200, gin.H{
			"key": key,
		})
	})
}
