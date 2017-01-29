package repositories

import (
	"database/sql"
	"log"

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

func (lr *ListRepository) Create(l models.List) (id int64, err error) {
	stmt, err := lr.db.Prepare("INSERT INTO list (organiser, email) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(
		l.Organiser,
		l.Email,
	)
	if err != nil {
		log.Fatal(err)
	}

	id, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return
}
