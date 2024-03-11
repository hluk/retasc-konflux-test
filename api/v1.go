package api_v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type message struct {
	Message string `json:"message"`
}

// @Summary		Greeting message
// @Description	Greets
// @Accept			json
// @Produce		json
// @Success		200	{object} message
// @Router			/api/v1/hello [get]
func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, message{Message: "Hello"})
}
