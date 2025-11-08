package home

import "github.com/gofiber/fiber/v2"

type HomeHandler struct {
	router fiber.Router
}

func NewHandler(r fiber.Router) {
	h := &HomeHandler{
		router: r,
	}
	h.router.Get("/", h.home)
}

func (h HomeHandler) home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
