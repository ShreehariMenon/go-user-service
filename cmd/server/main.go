package main

import (
    "database/sql"
    "go-user-service/internal/handler"
    "go-user-service/internal/logger"
    "go-user-service/internal/middleware"
    "go-user-service/internal/repository"
    "go-user-service/internal/routes"
    "go-user-service/internal/service"
    "log"

    _ "github.com/go-sql-driver/mysql"
    "github.com/gofiber/fiber/v2"
)

func main() {
    logger.Init()

    // CHANGE PASSWORD HERE
    dsn := "root:YOUR_PASSWORD@tcp(127.0.0.1:3306)/userdb?parseTime=true"
    dbConn, err := sql.Open("mysql", dsn)
    if err != nil { log.Fatal(err) }

    repo := repository.NewRepo(dbConn)
    svc := service.NewUserService(repo)
    h := handler.NewUserHandler(svc)

    app := fiber.New()
    middleware.Setup(app)
    routes.Setup(app, h)

    log.Fatal(app.Listen(":3000"))
}