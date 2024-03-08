package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello"})
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
