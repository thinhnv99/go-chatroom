package main

import (
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := InitDB()
	SetupRouter(db)
}
