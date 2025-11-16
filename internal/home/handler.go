package home

import (
	"fmt"
	"go-fullstack/internal/vacancy"
	"go-fullstack/pkg/tadapter"
	"go-fullstack/views"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router fiber.Router
	logger *zerolog.Logger
	repo   *vacancy.VacancyRepository
}

func NewHandler(r fiber.Router, cl *zerolog.Logger, repo *vacancy.VacancyRepository) {
	h := &HomeHandler{
		router: r,
		logger: cl,
		repo:   repo,
	}
	h.router.Get("/", h.home)
}

func (h HomeHandler) home(c *fiber.Ctx) error {
	PAGE_ITEMS := 2
	page := c.QueryInt("page", 1)
	count := h.repo.CountAll()

	// Calculate total pages correctly
	totalPages := (count + PAGE_ITEMS - 1) / PAGE_ITEMS

	h.logger.Info().Msg(fmt.Sprintf("page - %d, totalPages - %d, count - %d", page, totalPages, count))

	vacancies, err := h.repo.GetAll(PAGE_ITEMS, PAGE_ITEMS*(page-1))
	if err != nil {
		h.logger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}

	component := views.Main(vacancies, totalPages, page)
	return tadapter.Render(c, component, http.StatusOK)
}
