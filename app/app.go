package app

import (
	"database/sql"
	"log"
	"os"

	"github.com/Mkkysh/AvitoTest/handlers/routes"
	"github.com/Mkkysh/AvitoTest/utils/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	db *sql.DB
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {

	DB, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	a.db = DB

	// defer db.Close(a.db)

	fiberApp := fiber.New()
	fiberApp.Use(recover.New())

	router := routes.New()
	router.Run(fiberApp, a.db)

	PORT, _ := os.LookupEnv("PORT")

	log.Fatal(fiberApp.Listen(":" + PORT))

}
