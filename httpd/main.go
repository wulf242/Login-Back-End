package main

import (
	"my-back/authForm"
	"my-back/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	request := authForm.New()

	api := r.Group("/api")
	{
		api.POST("/validate", handler.CreateRequest(request))
		api.GET("/validate", handler.Validate(request))
	}
	r.Run()
}
