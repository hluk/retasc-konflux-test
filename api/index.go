package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

type IndexData struct {
	SwaggerUI  string `json:"swagger_ui"`
	Version    string `json:"version"`
	Revision   string `json:"revision"`
	LastCommit string `json:"last_commit"`
	DirtyBuild bool   `json:"dirty_build"`
}

func NewIndexData(swaggerUi string) IndexData {
	return IndexData{
		SwaggerUI:  swaggerUi,
		Version:    "unknown",
		Revision:   "unknown",
		LastCommit: "unknown",
		DirtyBuild: true,
	}
}

func NewIndexDataFromBuildInfo(swaggerUi string, info *debug.BuildInfo) IndexData {
	index := NewIndexData(swaggerUi)
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
	return index
}

// @Summary		Index
// @Description	Index
// @Accept			json
// @Produce		json
// @Success		200	{object}	IndexData
// @Router			/ [get]
func index(c *gin.Context) {
	swaggerUrl := fmt.Sprintf(
		"%s://%s/swagger/index.html", c.Request.URL.Scheme, c.Request.Host)

	info, ok := debug.ReadBuildInfo()
	if ok {
		index := NewIndexDataFromBuildInfo(swaggerUrl, info)
		c.JSON(http.StatusOK, index)
	} else {
		c.JSON(http.StatusOK, NewIndexData(swaggerUrl))
	}
}
