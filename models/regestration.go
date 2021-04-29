package models

import (
	"github.com/SwapRane12/scamdog/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Registration struct {
	gorm.Model
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	ReferalID uint   `json:"referalID"`
}

func NewRegestration(c *fiber.Ctx) error {
	db := database.DBConn
	registration := new(Registration)
	if err := c.BodyParser(registration); err != nil {
		return c.Status(503).SendString("Error Occured")
	}
	db.Create(&registration)
	return c.JSON(registration)
}
