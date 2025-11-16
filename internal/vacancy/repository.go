package vacancy

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type VacancyRepository struct {
	Dbpool *pgxpool.Pool
	Logger *zerolog.Logger
}

func NewRepo(dbpool *pgxpool.Pool, cl *zerolog.Logger) *VacancyRepository {
	return &VacancyRepository{
		Dbpool: dbpool,
		Logger: cl,
	}
}

func (r *VacancyRepository) GetAll(limit, offset int) ([]Vacancy, error) {
	q := `SELECT * FROM vacancies ORDER BY created_at desc LIMIT @limit OFFSET @offset`
	args := pgx.NamedArgs{
		"limit":  limit,
		"offset": offset,
	}
	rows, err := r.Dbpool.Query(context.Background(), q, args)
	if err != nil {
		return nil, err
	}
	vacancies, err := pgx.CollectRows(rows, pgx.RowToStructByName[Vacancy])
	if err != nil {
		return nil, err
	}
	return vacancies, nil
}

func (r *VacancyRepository) CountAll() int {
	query := `SELECT count(*) FROM vacancies`
	var count int
	r.Dbpool.QueryRow(context.Background(), query).Scan(&count)
	return count
}

func (r *VacancyRepository) addVacancy(form VacancyCreateForm) error {
	query := `INSERT INTO vacancies (email, role, salary, company, location, type, created_at) VALUES(@email, @role, @salary, @company, @location, @type, @created_at)`
	args := pgx.NamedArgs{
		"email":      form.Email,
		"role":       form.Role,
		"salary":     form.Salary,
		"company":    form.Company,
		"location":   form.Location,
		"type":       form.Type,
		"created_at": time.Now(),
	}
	_, err := r.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("не удвлось добавить запись вакансии в бд %w", err)
	}
	return nil
}
