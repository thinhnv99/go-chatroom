package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.Redirect(http.StatusFound, "/signup")
}

func SignUpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func SignInPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", nil)
}
