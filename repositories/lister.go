package repositories

import "github.com/KoseSoftware/secret-santa-api/models"

type ListerRepository interface {
	Create(l models.List) (id int64, err error)
	FindAll() (items []models.List, err error)
	FindByID(id int64) (item models.List, err error)
}
