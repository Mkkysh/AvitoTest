package main

import (
	"log"

	"github.com/Mkkysh/AvitoTest/app"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	app := app.NewApp()
	app.Run()
}
