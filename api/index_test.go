package api

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"runtime/debug"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIndexSwaggerUrlSimple(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var body map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &body)
	assert.Nil(t, err)

	value, exists := body["swagger_ui"]
	assert.True(t, exists)
	assert.Equal(t, "http:///swagger/index.html", value.(string))
}

func TestIndexSwaggerUrlFull(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	c.Request, _ = http.NewRequest(http.MethodGet, "https://test.example.com/", nil)
	c.Request.TLS = &tls.ConnectionState{}
	assert.True(t, nil != c.Request.TLS)
	index(c)

	assert.Equal(t, 200, w.Code)

	var body map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &body)
	assert.Nil(t, err)

	value, exists := body["swagger_ui"]
	assert.True(t, exists)
	assert.Equal(t, "https://test.example.com/swagger/index.html", value.(string))
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
