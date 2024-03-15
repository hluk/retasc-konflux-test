package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var body map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &body)
	assert.Nil(t, err)

	value, exists := body["swagger_ui"]
	assert.True(t, exists)
	assert.Equal(t, "http:///swagger/index.html", value.(string))
}

func TestIndexData(t *testing.T) {
	index := NewIndexData("/swagger/index.html")
	assert.Equal(t, "/swagger/index.html", index.SwaggerUI)
	assert.Equal(t, "unknown", index.Version)
	assert.Equal(t, "unknown", index.Revision)
	assert.Equal(t, "unknown", index.LastCommit)
	assert.Equal(t, true, index.DirtyBuild)
}

func TestIndexDataFromBuildInfo(t *testing.T) {
	info := debug.BuildInfo{}
	info.Main.Version = "(debug)"
	info.Settings = []debug.BuildSetting{
		{Key: "vcs.revision", Value: "REVISION"},
		{Key: "vcs.time", Value: "TIME"},
		{Key: "vcs.modified", Value: "false"},
	}
	index := NewIndexDataFromBuildInfo("/swagger/index.html", &info)
	assert.Equal(t, "/swagger/index.html", index.SwaggerUI)
	assert.Equal(t, "(debug)", index.Version)
	assert.Equal(t, "REVISION", index.Revision)
	assert.Equal(t, "TIME", index.LastCommit)
	assert.Equal(t, false, index.DirtyBuild)
}
