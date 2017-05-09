package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KoseSoftware/secret-santa-api/models"
	"github.com/KoseSoftware/secret-santa-api/repositories"
	"github.com/gorilla/mux"
	"github.com/mholt/binding"
	"github.com/unrolled/render"
)

type ListController struct {
	listRepository repositories.ListerRepository
	view           *render.Render
}

func NewListController(lr repositories.ListerRepository, v *render.Render) *ListController {
	return &ListController{
		listRepository: lr,
		view:           v,
	}
}

func (lc *ListController) GetList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	item, err := lc.listRepository.FindByID(vars["id"])
	if err != nil {
		errors := make([]models.Error, 0)
		errors = append(errors, models.Error{
			Message: err.Error(),
		})

		lc.view.JSON(w, http.StatusNotFound, models.Errors{
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
			Message: "List item not found",
			Errors:  errors,
		})

		return
	}

	// links
	url, _ := mux.CurrentRoute(r).URL("id", vars["id"])
	links := models.Links{
		Self: url.String(),
	}

	item.Links.Self = url.String()

	lc.view.JSON(w, http.StatusOK, models.Success{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Links:  links,
		Data:   item,
	})
}

// https://github.com/golang/go/wiki/CodeReviewComments#receiver-type
func (lc *ListController) GetLists(w http.ResponseWriter, r *http.Request) {
	var items []models.List
	var errors []models.Error

	// find a better solution
	email := r.URL.Query().Get("email")

	if email == "" {
		errors = append(errors, models.Error{
			Message: "Email: Provide an email address",
		})

		lc.view.JSON(w, http.StatusBadRequest, models.Errors{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Query parameter missing",
			Errors:  errors,
		})

		return
	}

	items, err := lc.listRepository.FindAll(email)
	if err != nil {
		log.Print(err.Error())

		return
	}

	url, _ := mux.CurrentRoute(r).URL()
	links := models.Links{
		Self: url.String(),
	}

	for i, item := range items {
		items[i].Links.Self = fmt.Sprintf("%s/%s", url.String(), item.ID)
	}

	if len(items) == 0 {
		items = make([]models.List, 0)
	}

	lc.view.JSON(w, http.StatusOK, models.Success{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Links:  links,
		Data:   items,
	})
}

func (lc *ListController) PostList(w http.ResponseWriter, r *http.Request) {
	errors := make([]models.Error, 0)
	list := new(models.List)

	errs := binding.Bind(r, list)
	if errs != nil {
		for _, msg := range errs {
			errors = append(errors, models.Error{
				Message: msg.Error(),
			})
		}

		lc.view.JSON(w, http.StatusBadRequest, models.Errors{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Invalid data",
			Errors:  errors,
		})

		return
	}

	id, err := lc.listRepository.Create(*list)
	if err != nil {
		errors = append(errors, models.Error{
			Message: err.Error(),
		})

		lc.view.JSON(w, http.StatusBadRequest, models.Errors{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Failed to create list",
			Errors:  errors,
		})

		return
	}

	url, _ := mux.CurrentRoute(r).URL()
	location := fmt.Sprintf("%s/%s", url.String(), id)

	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
}

func (lc *ListController) PutList(w http.ResponseWriter, r *http.Request) {
	// implement
}

func (lc *ListController) DeleteList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	rowsAffected, err := lc.listRepository.DeleteByID(vars["id"])
	if err != nil {
		// handle
	}

	if rowsAffected == 1 {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	// https://developers.google.com/youtube/v3/docs/core_errors
	w.WriteHeader(http.StatusNotFound)
}
