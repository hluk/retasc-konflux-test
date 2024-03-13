package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
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
