package main

import (
	"log"

	"github.com/KoseSoftware/secret-santa-api/config"
	"github.com/KoseSoftware/secret-santa-api/controllers"
	"github.com/KoseSoftware/secret-santa-api/repositories"
	"github.com/gorilla/mux"
	"github.com/markbates/pop"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

func main() {
	v := render.New()

	db, err := pop.Connect("development")
	if err != nil {
		log.Fatal(err)
	}

	lr := repositories.NewPopListRepository(db)

	hc := controllers.NewHomepageController(v)
	lc := controllers.NewListsController(lr, v)

	r := mux.NewRouter()
	r.HandleFunc("/", hc.Index).Methods("GET").Name("homepage")

	r.HandleFunc("/lists", lc.GetLists).Methods("GET").Name("get_lists")
	r.HandleFunc("/lists", lc.PostLists).Methods("POST").Name("post_lists")
	r.HandleFunc("/lists/{id:[0-9]+}", lc.GetList).Methods("GET").Name("get_list")

	n := negroni.Classic()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.UseHandler(r)

	log.Fatal(config.GetServer(n).ListenAndServe())
}
