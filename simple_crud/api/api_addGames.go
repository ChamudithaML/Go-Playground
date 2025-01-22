package api

import (
	"fmt"
	"main/dao"
	"main/dto"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AddGames(c *fiber.Ctx) error {
	fmt.Println("in game adding")
	newGame := dto.Game{}

	if err := c.BodyParser(&newGame); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	newGame.CreatedAt = time.Now().Format("2006-01-02")

	if err := dao.CreateGame(newGame); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create game",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Game created successfully",
		"game":    newGame,
	})
}
