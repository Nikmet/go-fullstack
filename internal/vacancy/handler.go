package vacancy

import (
	"go-fullstack/pkg/tadapter"
	"go-fullstack/pkg/validator"
	"go-fullstack/views/components"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHandler struct {
	router fiber.Router
	cl     *zerolog.Logger
}

func NewHanndler(r fiber.Router, cl *zerolog.Logger) {
	h := &VacancyHandler{
		router: r,
		cl:     cl,
	}
	vacancyGruop := h.router.Group("/vacancy")
	vacancyGruop.Post("/", h.createVacancy)
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
	h.cl.Info().Msg(form.Email)
	if len(errors.Errors) > 0 {
		comp := components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, comp)
	}
	return tadapter.Render(c, components.Notification("Вакансия успешно создана!", components.NotificationSucces))
}
