package models

import "time"

type List struct {
	ID        int       `json:"id,omitempty" db:"id,omitempty"`
	Organiser string    `json:"organiser" db:"organiser"`
	Email     string    `json:"email" db:"email"`
	Amount    float64   `json:"amount" db:"amount"`
	Date      time.Time `json:"date" db:"date"`
	Location  string    `json:"location,omitempty" db:"location"`
	Notes     string    `json:"notes,omitempty" db:"notes"`
	Created   time.Time `json:"created" db:"created"`
	Updated   time.Time `json:"updated" db:"updated"`
}
