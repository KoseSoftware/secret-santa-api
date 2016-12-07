package main

import (
	"log"
	"net/http"

	"github.com/KoseSoftware/secret-santa-api/config"
	"github.com/KoseSoftware/secret-santa-api/controllers"
	"github.com/gorilla/mux"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

func main() {
	v := render.New()

	hc := controllers.NewHomepageController(v)
	lc := controllers.NewListsController(v)

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", hc.Index).Methods("GET").Name("homepage")

	l := r.PathPrefix("/lists").Subrouter().StrictSlash(true)
	l.HandleFunc("/", lc.GetLists).Methods("GET").Name("get_lists")
	l.HandleFunc("/", lc.PostLists).Methods("POST").Name("post_lists")

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.Use(negroni.NewStatic(http.Dir("public")))
	n.UseHandler(r)

	log.Fatal(config.GetServer(n).ListenAndServe())
}
