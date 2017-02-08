package repositories

import (
	"github.com/KoseSoftware/secret-santa-api/models"
	"github.com/markbates/pop"
)

type PopListRepository struct {
	db *pop.Connection
}

func NewPopListRepository(db *pop.Connection) *PopListRepository {
	return &PopListRepository{
		db: db,
	}
}

func (lr *PopListRepository) Create(l models.List) (id int64, err error) {
	return
}

func (lr *PopListRepository) FindAll() (items []models.List, err error) {
	err = lr.db.Store.Select(&items, "SELECT * FROM list")

	return
}

func (lr *PopListRepository) FindByID(id int64) (item models.List, err error) {
	err = lr.db.Store.Get(&item, "SELECT * FROM list WHERE id = ?", id)

	return
}
