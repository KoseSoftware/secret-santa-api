package models

import "time"

type List struct {
	ID        int       `db:"id,omitempty" json:"id,omitempty"`
	Organiser string    `db:"organiser" json:"organiser"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}
