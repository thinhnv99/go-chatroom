package main

import (
	"chatroom/handlers"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // your frontend origin here
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
	r.POST("/signin", userHandler.SignIn)

	r.Run(":8088")
}
