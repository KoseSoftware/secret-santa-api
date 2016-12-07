package models

import "time"

type List struct {
	ID        int       `db:"id,omitempty" json:"id,omitempty"`
	Organiser string    `db:"organiser" json:"organiser"`
	Email     string    `db:"email" json:"email"`
	Amount    float64   `db:"amount" json:"amount"`
	Date      time.Time `db:"date" json:"date"`
	Location  string    `db:"location" json:"location,omitempty"`
	Notes     string    `db:"notes" json:"notes,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}
