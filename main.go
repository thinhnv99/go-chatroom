package main

import (
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := InitDB()
	r := SetupRouter(db)
	r.Run(":8088")
}
