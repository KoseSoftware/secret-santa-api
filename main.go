package main

import (
	"log"
	"net/http"

	"github.com/KoseSoftware/secret-santa-api/config"
	"github.com/KoseSoftware/secret-santa-api/controllers"
	"github.com/KoseSoftware/secret-santa-api/repositories"
	"github.com/gorilla/mux"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"upper.io/db.v3/mysql"
)

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})

func main() {
	v := render.New()

	sess, err := mysql.Open(config.GetDbSettings())
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}
	defer sess.Close()

	lr := repositories.NewUpperListRepository(sess)

	hc := controllers.NewHomepageController(v)
	lc := controllers.NewListController(lr, v)

	r := mux.NewRouter()
	r.HandleFunc("/", hc.Index).Methods("GET").Name("homepage")

	v1 := r.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/lists", lc.GetLists).Methods("GET").Name("get_lists")
	v1.HandleFunc("/lists", lc.PostList).Methods("POST").Name("post_list")
	v1.HandleFunc("/lists/{id:[a-zA-Z0-9]+}", lc.GetList).Methods("GET").Name("get_list")
	v1.HandleFunc("/lists/{id:[a-zA-Z0-9]+}", lc.PutList).Methods("PUT").Name("put_list")
	v1.HandleFunc("/lists/{id:[a-zA-Z0-9]+}", lc.DeleteList).Methods("DELETE").Name("delete_list")
	v1.HandleFunc("/lists/{id:[a-zA-Z0-9]+}/exclusions", NotImplemented).Methods("GET")
	v1.HandleFunc("/lists/{id:[a-zA-Z0-9]+}/exclusions", NotImplemented).Methods("POST")
	v1.HandleFunc("/lists/{id:[a-zA-Z0-9]+}/exclusions", NotImplemented).Methods("PUT")

	v1.HandleFunc("/status", StatusHandler).Methods("GET")

	n := negroni.Classic()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.UseHandler(r)

	log.Fatal(config.GetServer(n).ListenAndServe())
}
