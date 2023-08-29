package routes

import (
	"database/sql"

	"github.com/Mkkysh/AvitoTest/handlers/controllers"
	"github.com/Mkkysh/AvitoTest/handlers/services"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
}

func New() *Router {
	return &Router{}
}

func (r *Router) Run(fiberApp *fiber.App, db *sql.DB) {

	logService := services.NewLogService(db)
	logController := controllers.NewLogController(logService)
	logs := fiberApp.Group("/api/logs")

	logs.Get("/", logController.Get)

	segmentService := services.NewSegmentService(db)
	segmentController := controllers.NewSegmentController(segmentService)
	segments := fiberApp.Group("/api/segments")

	segments.Post("/", segmentController.Add)
	segments.Delete("/", segmentController.Delete)

	userService := services.NewUserService(db, logService)
	userController := controllers.NewUserController(userService)
	users := fiberApp.Group("/api/users")

	users.Patch("/:id", userController.UpdateSegment)
	users.Get("/:id", userController.GetSegemnts)

}
