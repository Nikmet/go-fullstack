package vacancy

import "time"

type VacancyCreateForm struct {
	Email    string
	Role     string
	Location string
	Type     string
	Salary   string
	Company  string
}

type Vacancy struct {
	Id        int       `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	Email     string    `db:"email"`
	Role      string    `db:"role"`
	Location  string    `db:"location"`
	Type      string    `db:"type"`
	Salary    string    `db:"salary"`
	Company   string    `db:"company"`
}

type GetAllRequestBody struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
