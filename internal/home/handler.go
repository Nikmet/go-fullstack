package home

import (
	"go-fullstack/pkg/tadapter"
	"go-fullstack/views"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router fiber.Router
	cl     *zerolog.Logger
}

func NewHandler(r fiber.Router, cl *zerolog.Logger) {
	h := &HomeHandler{
		router: r,
		cl:     cl,
	}
	h.router.Get("/", h.home)
}

func (h HomeHandler) home(c *fiber.Ctx) error {
	component := views.Main()
	return tadapter.Render(c, component)
}
