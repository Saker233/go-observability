package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	godotenv.Load("app.env")

	router := routes()

	err := router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

}
