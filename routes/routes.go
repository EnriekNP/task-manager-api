package routes

import (
	"task-manager-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1") // Hardcoded versioning

	api.Post("/auth/register", handlers.Register)
	api.Post("/auth/login", handlers.Login)

	// taskRoutes := api.Group("/tasks")
	// taskRoutes.Get("/", handlers.GetAllTasks)
	// taskRoutes.Post("/", handlers.CreateTask)
	// taskRoutes.Get("/:id", handlers.GetTaskByID)
	// taskRoutes.Put("/:id", handlers.UpdateTask)
	// taskRoutes.Delete("/:id", handlers.DeleteTask)

	// userRoutes := api.Group("/users")
	// userRoutes.Get("/me", handlers.GetProfile)
	// userRoutes.Put("/me", handlers.UpdateProfile)
}
