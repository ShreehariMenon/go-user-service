package routes

import (
    "go-user-service/internal/handler"
    "github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, h *handler.UserHandler) {
    app.Post("/users", h.Create)
    app.Get("/users", h.GetAll)
}