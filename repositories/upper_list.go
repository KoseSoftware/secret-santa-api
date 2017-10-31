package repositories

import (
	"log"
	"time"

	"github.com/KoseSoftware/secret-santa-api/models"
	"github.com/dchest/uniuri"
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

func (lr *UpperListRepository) Create(l models.List) (id string, err error) {
	// set id
	l.ID = uniuri.NewLen(7)

	// set timestamp
	l.Created = time.Now()
	l.Updated = time.Now()

	q := lr.sess.InsertInto(listTable).Values(l)

	_, err = q.Exec()
	if err != nil {
		log.Fatalf("sess.InsertInto(): %q\n", err)
	}

	id = l.ID

	return
}

func (lr *UpperListRepository) FindByID(id string) (item models.List, err error) {
	res := lr.sess.Collection(listTable).Find().Where(db.Cond{
		"id": id,
	})

	err = res.One(&item)
	if err != nil {
		log.Fatalf("res.One(): %q\n", err)
	}

	return
}

func (lr *UpperListRepository) DeleteByID(id string) (rowsAffected int64, err error) {
	q := lr.sess.DeleteFrom(listTable).Where(db.Cond{
		"id": id,
	})

	result, err := q.Exec()
	if err != nil {
		log.Fatalf("sess.DeleteFrom(): %q\n", err)
	}

	rowsAffected, err = result.RowsAffected()

	return
}
