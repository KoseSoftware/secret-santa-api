package repositories

import (
	"log"

	"time"

	"github.com/KoseSoftware/secret-santa-api/models"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

type UpperListRepository struct {
	sess sqlbuilder.Database
}

func NewUpperListRepository(sess sqlbuilder.Database) *UpperListRepository {
	return &UpperListRepository{
		sess: sess,
	}
}

func (lr *UpperListRepository) Create(l models.List) (id int64, err error) {
	// set timestamp
	l.Created = time.Now()
	l.Updated = time.Now()

	q := lr.sess.InsertInto("lists").Values(l)

	result, err := q.Exec()
	if err == nil {
		id, err = result.LastInsertId()
	}

	return
}

func (lr *UpperListRepository) FindAll() (items []models.List, err error) {
	res := lr.sess.Collection("lists").Find()

	err = res.All(&items)
	if err != nil {
		log.Fatalf("res.All(): %q\n", err)
	}

	return
}

func (lr *UpperListRepository) FindByID(id int64) (item models.List, err error) {
	res := lr.sess.Collection("lists").Find().Where(db.Cond{
		id: id,
	})

	err = res.One(&item)
	if err != nil {
		log.Fatalf("res.One(): %q\n", err)
	}

	return
}
