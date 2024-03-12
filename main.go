package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/hluk/retasc-konflux-test/api"
	_ "github.com/hluk/retasc-konflux-test/docs"
)

type indexData struct {
	SwaggerUI  string `json:"swagger_ui"`
	Version    string `json:"version"`
	Revision   string `json:"revision"`
	LastCommit string `json:"last_commit"`
	DirtyBuild bool   `json:"dirty_build"`
}

// @Summary		Index
// @Description	Index
// @Accept			json
// @Produce		json
// @Success		200	{object}	indexData
// @Router			/ [get]
func index(c *gin.Context) {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	swaggerUrl := fmt.Sprintf("%s://%s/swagger/index.html", scheme, c.Request.Host)

	index := indexData{
		SwaggerUI:  swaggerUrl,
		Version:    "unknown",
		Revision:   "unknown",
		LastCommit: "unknown",
		DirtyBuild: true,
	}

	info, ok := debug.ReadBuildInfo()
	if ok {
		index.Version = info.Main.Version
		for _, kv := range info.Settings {
			switch kv.Key {
			case "vcs.revision":
				index.Revision = kv.Value
			case "vcs.time":
				index.LastCommit = kv.Value
			case "vcs.modified":
				index.DirtyBuild = kv.Value == "true"
			}
		}
	}

	c.JSON(http.StatusOK, index)
}

func setupRouter() *gin.Engine {
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

//	@title			ReTaSC Konflux Test
//	@version		1.0
//	@description	Proof of concept

//	@contact.name	API Support
//	@contact.url	https://github.com/hluk/retasc-konflux-test
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	router := setupRouter()
	err := router.Run("0.0.0.0:8081")
	if err != nil {
		log.Fatal(err)
	}
}
