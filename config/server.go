package config

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/negroni"
)

var httpAddr = 3000

func init() {
	if os.Getenv("HTTP_ADDR") != "" {
		httpAddr, _ = strconv.Atoi(os.Getenv("HTTP_ADDR"))
	}
}

func GetServer(n *negroni.Negroni) *http.Server {
	return &http.Server{
		Addr:           fmt.Sprintf(":%d", httpAddr),
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
