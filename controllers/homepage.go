package controllers

import (
	"fmt"
	"net/http"
)

type HomepageController struct {
	bc *BaseController
}

func NewHomepageController(bc *BaseController) *HomepageController {
	return &HomepageController{
		bc: bc,
	}
}

func (c *HomepageController) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Gorilla! Welcome to the home page. DB String: %s\n", c.bc.dbConn)))
}

func (c *HomepageController) Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Gorilla! Test that my file save reloads!"))
}
