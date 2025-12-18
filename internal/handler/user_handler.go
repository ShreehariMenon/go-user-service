package handler

import (
    "go-user-service/internal/models"
    "go-user-service/internal/service"
    "github.com/go-playground/validator/v10"
    "github.com/gofiber/fiber/v2"
)

type UserHandler struct {
    svc *service.UserService
    val *validator.Validate
}

func NewUserHandler(s *service.UserService) *UserHandler {
    return &UserHandler{svc: s, val: validator.New()}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
    var req models.CreateUserRequest
    if err := c.BodyParser(&req); err != nil { return c.SendStatus(400) }
    if err := h.val.Struct(req); err != nil { return c.SendStatus(400) }

    id, err := h.svc.Create(c.Context(), req.Name, req.Dob)
    if err != nil { return c.SendStatus(500) }

    return c.JSON(models.UserResponse{ID: id, Name: req.Name, Dob: req.Dob})
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
    users, _ := h.svc.Repo().Q.ListUsers(c.Context())
    var resp []models.UserResponse
    for _, u := range users {
        resp = append(resp, models.UserResponse{
            ID: u.ID, Name: u.Name, Dob: u.Dob.Format("2006-01-02"), Age: h.svc.CalculateAge(u.Dob),
        })
    }
    return c.JSON(resp)
}