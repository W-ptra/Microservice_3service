package controller

import "github.com/gofiber/fiber/v2"

func NotFound404(c *fiber.Ctx)error{
	return c.Status(404).JSON(fiber.Map{
		"Message":"404 Not Found",
	})
}