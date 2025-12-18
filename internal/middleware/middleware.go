package middleware

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/requestid"
)

func Setup(app *fiber.App) {
    app.Use(requestid.New())
    app.Use(logger.New())
}