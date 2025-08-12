package main

import (
	"chatroom/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil || sqlDB.Ping() != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "DB down",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
	
	userHandler := handlers.UserHandler{DB: db}
	r.POST("/signup", userHandler.SignUp)

	r.LoadHTMLGlob("templates/*")

	r.GET("/", handlers.HomePage)
	r.GET("/signup", handlers.SignUpPage)
	r.GET("/signin", handlers.SignInPage)

	return r
}
