package repositories

import "github.com/KoseSoftware/secret-santa-api/models"

const listTable = "lists"

type ListerRepository interface {
	Create(l models.List) (id string, err error)
	Update(l models.List) (rowsAffected int64, err error)
	FindByID(id string) (item models.List, err error)
	DeleteByID(id string) (rowsAffected int64, err error)
}
