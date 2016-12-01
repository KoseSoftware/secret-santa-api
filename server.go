package main

import (
	"log"
	"net/http"
	"time"

	"github.com/KoseSoftware/secret-santa-api/config"
	"github.com/KoseSoftware/secret-santa-api/controllers"
	"github.com/urfave/negroni"
)

func main() {
	bc := controllers.NewBaseController("my database connection")
	r := config.NewRouter(bc)

	n := negroni.Classic()
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
