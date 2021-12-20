package main

import (
	"my-gram/database"
	"my-gram/router"
	"os"
)

func main() {

	database.StartDB()
	r := router.StartApp()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
