package controllers

import "net/http"

type HomepageController struct {
}

func NewHomepageController() *HomepageController {
	return &HomepageController{}
}

func (hc HomepageController) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the home page."))
}
