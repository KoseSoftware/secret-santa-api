package main

import (
	"log"
	"net/http"
	"time"

	"github.com/KoseSoftware/secret-santa-api/controllers"
	"github.com/gorilla/mux"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"
)

func main() {
	hc := controllers.NewHomepageController()
	lc := controllers.NewListsController()

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

	s := &http.Server{
		Addr:           ":3000",
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
