package middlewares

import (
	"bytes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func BufferMiddleware(c *fiber.Ctx) error {
	const maxBufferSize = 1 << 20

	body := c.Body()

	if len(body) > maxBufferSize {
		return c.Status(fiber.StatusRequestEntityTooLarge).SendString("Request body too large")
	}

	buffer := bytes.NewBuffer(body)
	log.Println("Buffered Data:", buffer.String())

	c.Request().SetBody(buffer.Bytes())

	return c.Next()
}
