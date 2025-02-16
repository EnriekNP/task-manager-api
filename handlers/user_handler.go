package handlers

import (
	"log"
	"net/http"
	"task-manager-api/models"
	"task-manager-api/repositories"
	"task-manager-api/utils"

	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Register
func Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestFormat,
		})
	}

	if errors := utils.ValidateStruct(req); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  utils.ErrValidationFailed,
			"fields": errors,
		})
	}

	existingUser, err := repositories.GetUserByEmail(req.Email)
	if err != nil {
		log.Println("❌ Error checking user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrInternalServer,
		})
	}

	if existingUser != nil {
		log.Println("⚠️ Email already exists:", req.Email)
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": utils.ErrEmailExists,
		})
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrInternalServer,
		})
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err = repositories.CreateUser(&user); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrInternalServer,
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestFormat,
		})
	}

	if errors := utils.ValidateStruct(req); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  utils.ErrValidationFailed,
			"fields": errors,
		})
	}

	user, err := repositories.GetUserByEmail(req.Email)
	if err != nil {
		log.Println("❌ Database error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrInternalServer,
		})
	}

	// Jika user tidak ditemukan
	if user == nil {
		log.Println("⚠️ User not found:", req.Email)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": utils.ErrInvalidCredentials,
		})
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		log.Println("⚠️ Invalid password for user:", req.Email)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": utils.ErrInvalidCredentials,
		})
	}

	token, err := utils.GenerateToken(uint(user.ID), user.Email)
	if err != nil {
		log.Println("❌ Token generation error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrTokenGeneration,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}
