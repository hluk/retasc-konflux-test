package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type message struct {
    Message string `json:"message"`
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, message{Message: "Hello"})
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/hello", hello)
	return router
}

func main() {
	router := setupRouter()
	router.Run("localhost:8081")
}
