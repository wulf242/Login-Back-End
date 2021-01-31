package handler

import (
	"fmt"
	"my-back/authForm"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Validate(request *authForm.AuthRequest) gin.HandlerFunc {
	return func(c *gin.Context) {
		form := request.Unpack()
		hours, minutes, _ := time.Now().Clock()
		code := fmt.Sprintf("%d%02d", hours, minutes)
		if form["email"] == "c137@onecause.com" && form["password"] == "#th@nH@rm#y#r!$100%D0p#" && form["token"] == code {
			fmt.Print(code)
			c.String(http.StatusOK, code)
		} else {
			c.JSON(http.StatusOK, code)
		}
	}
}

func CreateRequest(request *authForm.AuthRequest) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Bind(&request)
		c.Status(http.StatusOK)
	}
}
