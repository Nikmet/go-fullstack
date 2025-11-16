package vacancy

import (
	"go-fullstack/pkg/tadapter"
	"go-fullstack/pkg/validator"
	"go-fullstack/views/components"
	"net/http"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHandler struct {
	router fiber.Router
	logger *zerolog.Logger
	repo   *VacancyRepository
}

func NewHanndler(r fiber.Router, repo *VacancyRepository, cl *zerolog.Logger) {
	h := &VacancyHandler{
		router: r,
		logger: cl,
		repo:   repo,
	}
	vacancyGruop := h.router.Group("/vacancy")
	vacancyGruop.Post("/", h.createVacancy)
	vacancyGruop.Get("/", h.getAll)
}

func (h *VacancyHandler) getAll(c *fiber.Ctx) error {
	var body GetAllRequestBody

	// Парсим тело запроса
	if err := c.BodyParser(&body); err != nil {
		h.logger.Error().Msgf("Failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Устанавливаем значения по умолчанию, если не переданы
	if body.Limit <= 0 {
		body.Limit = 10 // или другое значение по умолчанию
	}
	if body.Offset < 0 {
		body.Offset = 0
	}

	vacancies, err := h.repo.GetAll(body.Limit, body.Offset)
	if err != nil {
		h.logger.Error().Msg(err.Error())
	}
	return c.JSON(vacancies)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	form := VacancyCreateForm{
		Email:    c.FormValue("email"),
		Role:     c.FormValue("role"),
		Location: c.FormValue("location"),
		Type:     c.FormValue("type"),
		Salary:   c.FormValue("salary"),
		Company:  c.FormValue("company"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "неккоректный email"},
		&validators.StringIsPresent{Name: "Location", Field: form.Location, Message: "расположение не задано"},
		&validators.StringIsPresent{Name: "Type", Field: form.Type, Message: "сфера компании не задана"},
		&validators.StringIsPresent{Name: "Company", Field: form.Company, Message: "название компании не задано"},
		&validators.StringIsPresent{Name: "Role", Field: form.Role, Message: "должность не задана"},
		&validators.StringIsPresent{Name: "Salary", Field: form.Salary, Message: "зарплата не задана"},
	)
	if len(errors.Errors) > 0 {
		comp := components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, comp, http.StatusBadRequest)
	}
	err := h.repo.addVacancy(form)

	if err != nil {
		h.logger.Err(err)
		comp := components.Notification("Произзошла ошибка на сервере, попробуйте позже(", components.NotificationFail)
		return tadapter.Render(c, comp, http.StatusBadRequest)
	}

	return tadapter.Render(c, components.Notification("Вакансия успешно создана!", components.NotificationSucces), http.StatusOK)
}
