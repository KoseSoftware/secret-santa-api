package models

import (
	"net/http"
	"time"

	"github.com/KoseSoftware/secret-santa-api/responses"
	"github.com/mholt/binding"
)

type List struct {
	ID        int             `json:"id,omitempty" db:"id,omitempty"`
	Organiser string          `json:"organiser" db:"organiser"`
	Email     string          `json:"email" db:"email"`
	Amount    float64         `json:"amount" db:"amount"`
	Date      time.Time       `json:"date" db:"date"`
	Location  string          `json:"location,omitempty" db:"location"`
	Notes     string          `json:"notes,omitempty" db:"notes"`
	Created   time.Time       `json:"created" db:"created"`
	Updated   time.Time       `json:"updated" db:"updated"`
	Links     responses.Links `json:"_links,omitempty"`
}

func (l *List) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&l.Organiser: binding.Field{
			Form:     "organiser",
			Required: true,
		},
		&l.Email: binding.Field{
			Form:     "email",
			Required: true,
		},
		&l.Amount: binding.Field{
			Form:     "amount",
			Required: true,
		},
		&l.Date: binding.Field{
			Form:       "date",
			TimeFormat: time.RFC3339,
			Required:   true,
		},
		&l.Location: binding.Field{
			Form:     "location",
			Required: false,
		},
		&l.Notes: binding.Field{
			Form:     "notes",
			Required: false,
		},
	}
}
