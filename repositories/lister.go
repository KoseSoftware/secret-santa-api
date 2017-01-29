package repositories

import "github.com/KoseSoftware/secret-santa-api/models"

type Lister interface {
	Create(l models.List) (id int64, err error)
}
