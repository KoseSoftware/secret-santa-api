package repositories

import (
	"database/sql"
	"fmt"

	"github.com/KoseSoftware/secret-santa-api/models"
	_ "github.com/go-sql-driver/mysql"
)

type ListRepository struct {
	db *sql.DB
}

func NewListRepository(db *sql.DB) *ListRepository {
	return &ListRepository{
		db: db,
	}
}

func (lr *ListRepository) Create(l models.List) error {
	fmt.Println("create")

	return nil
}
