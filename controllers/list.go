package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func NewListsController(lr repositories.ListerRepository, r *render.Render) *ListController {
	return &ListController{
		listRepository: lr,
		view:           r,
	}
}

func (lc *ListController) GetList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	item, err := lc.listRepository.FindByID(int64(id))
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
	url, _ := mux.CurrentRoute(r).URL("id", strconv.Itoa(id))
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
	items, err := lc.listRepository.FindAll()
	if err != nil {
		log.Print(err.Error())

		return
	}

	url, _ := mux.CurrentRoute(r).URL()
	links := responses.Links{
		Self: url.String(),
	}

	for i, item := range items {
		items[i].Links.Self = fmt.Sprintf("%s/%s", url.String(), strconv.Itoa(item.ID))
	}

	lc.view.JSON(w, http.StatusOK, responses.Success{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Links:  links,
		Data:   items,
	})
}

func (lc *ListController) PostLists(w http.ResponseWriter, r *http.Request) {
	errors := make([]responses.Error, 0)
	list := new(models.List)

	errs := binding.Bind(r, list)
	if errs != nil {
		for _, msg := range errs {
			errors = append(errors, responses.Error{
				Message: msg.Error(),
			})
		}

		errorResponse(lc, w, http.StatusBadRequest, "invalid data", errors)
		return
	}

	id, err := lc.listRepository.Create(*list)
	if err != nil {
		errors = append(errors, responses.Error{
			Message: err.Error(),
		})

		errorResponse(lc, w, http.StatusUnprocessableEntity, "failed to create list", errors)
		return
	}

	url, _ := mux.CurrentRoute(r).URL()
	location := fmt.Sprintf("%s/%s", url.String(), strconv.Itoa(int(id)))

	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
}

func errorResponse(lc *ListController, w http.ResponseWriter, statusCode int, message string, errors []responses.Error) {
	lc.view.JSON(w, statusCode, responses.Errors{
		Code:    statusCode,
		Status:  http.StatusText(statusCode),
		Message: message,
		Errors:  errors,
	})
}
