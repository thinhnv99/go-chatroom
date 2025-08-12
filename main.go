package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	InitDB()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		sqlDB, err := DB.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"db": "failed to get DB instance",
			})
			return
		}

		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "error",
				"db": "unreachable",
			})
			return
		}
		
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"db": "connected",
		})
	})

	r.Run(":8088")
}

