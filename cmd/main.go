package main

import (
	"log"

	"github.com/Mkkysh/AvitoTest/app"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func main() {
	app := app.NewApp()
	app.Run()
}
