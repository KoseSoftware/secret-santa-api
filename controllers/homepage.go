package controllers

import (
	"net/http"

	"github.com/unrolled/render"
)

type HomepageController struct {
	view *render.Render
}

func NewHomepageController(r *render.Render) *HomepageController {
	return &HomepageController{
		view: r,
	}
}

func (hc *HomepageController) Index(w http.ResponseWriter, r *http.Request) {
	hc.view.Text(w, http.StatusOK, "Welcome to the homepage")
}
