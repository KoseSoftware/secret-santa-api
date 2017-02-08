package repositories

import (
	"fmt"

	"github.com/KoseSoftware/secret-santa-api/models"
	"github.com/markbates/pop"
)

const table = "lists"

type PopListRepository struct {
	db *pop.Connection
}

func NewPopListRepository(db *pop.Connection) *PopListRepository {
	return &PopListRepository{
		db: db,
	}
}

func (lr *PopListRepository) Create(l models.List) (id int64, err error) {
	result, err := lr.db.Store.NamedExec(fmt.Sprintf("INSERT INTO %s (organiser, email, amount, date, location, notes) VALUES (:organiser, :email, :amount, :date, :location, :notes)", table), &l)
	if err == nil {
		id, err = result.LastInsertId()
	}

	return
}

func (lr *PopListRepository) FindAll() (items []models.List, err error) {
	err = lr.db.Store.Select(&items, fmt.Sprintf("SELECT * FROM %s", table))

	return
}

func (lr *PopListRepository) FindByID(id int64) (item models.List, err error) {
	err = lr.db.Store.Get(&item, fmt.Sprintf("SELECT * FROM %s WHERE id = ? LIMIT 1", table), id)

	return
}
