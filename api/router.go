package api

import (
	"github.com/gin-gonic/gin"
	"log"

	swaggerfiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"github.com/hluk/retasc-konflux-test/api/v1"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", index)
	v1 := router.Group("/api/v1")
	router.ForwardedByClientIP = true
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatal(err)
	}
	{
		v1.GET("/hello", api_v1.Hello)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
