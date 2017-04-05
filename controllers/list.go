package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KoseSoftware/secret-santa-api/models"
	"github.com/KoseSoftware/secret-santa-api/repositories"
	"github.com/KoseSoftware/secret-santa-api/responses"
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
		errors := make([]responses.Error, 0)
		errors = append(errors, responses.Error{
			Message: err.Error(),
		})

		lc.view.JSON(w, http.StatusNotFound, responses.Errors{
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
			Message: "list item not found",
			Errors:  errors,
		})

		return
	}

	// links
	url, _ := mux.CurrentRoute(r).URL("id", vars["id"])
	links := responses.Links{
		Self: url.String(),
	}

	item.Links.Self = url.String()

	lc.view.JSON(w, http.StatusOK, responses.Success{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Links:  links,
		Data:   item,
	})
}

// https://github.com/golang/go/wiki/CodeReviewComments#receiver-type
func (lc *ListController) GetLists(w http.ResponseWriter, r *http.Request) {
	items, err := lc.listRepository.FindAll(r.URL.Query().Get("email"))
	if err != nil {
		log.Print(err.Error())

		return
	}

	url, _ := mux.CurrentRoute(r).URL()
	links := responses.Links{
		Self: url.String(),
	}

	for i, item := range items {
		items[i].Links.Self = fmt.Sprintf("%s/%s", url.String(), item.ID)
	}

	lc.view.JSON(w, http.StatusOK, responses.Success{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Links:  links,
		Data:   items,
	})
}

func (lc *ListController) PostList(w http.ResponseWriter, r *http.Request) {
	errors := make([]responses.Error, 0)
	list := new(models.List)

	errs := binding.Bind(r, list)
	if errs != nil {
		for _, msg := range errs {
			errors = append(errors, responses.Error{
				Message: msg.Error(),
			})
		}

		lc.view.JSON(w, http.StatusBadRequest, responses.Errors{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "invalid data",
			Errors:  errors,
		})

		return
	}

	id, err := lc.listRepository.Create(*list)
	if err != nil {
		errors = append(errors, responses.Error{
			Message: err.Error(),
		})

		lc.view.JSON(w, http.StatusBadRequest, responses.Errors{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "failed to create list",
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

}

func (lc *ListController) DeleteList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	rowsAffected, err := lc.listRepository.DeleteByID(vars["id"])
	if err != nil {

	}

	if rowsAffected == 1 {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	// https://developers.google.com/youtube/v3/docs/core_errors
	w.WriteHeader(http.StatusNotFound)
}
