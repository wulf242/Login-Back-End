package main

import (
	"fmt"
	"my-back/authForm"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	request := authForm.New()

	api := r.Group("/api")
	{
		api.POST("/validate", func(c *gin.Context) {
			c.BindJSON(&request)
			c.Status(http.StatusNoContent)
		})
		api.GET("/validate", func(c *gin.Context) {
			hours, minutes, _ := time.Now().Clock()
			code := fmt.Sprintf("%d%02d", hours, minutes)
			if request.Email == "c137@onecause.com" && request.Password == "#th@nH@rm#y#r!$100%D0p#" && request.Token == code {
				c.String(http.StatusOK, code)
			} else {
				c.JSON(http.StatusOK, "nope")
			}
		})
	}

	r.Run()
}
