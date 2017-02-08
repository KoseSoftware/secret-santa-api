package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/KoseSoftware/secret-santa-api/models"
	"github.com/KoseSoftware/secret-santa-api/repositories"
	"github.com/KoseSoftware/secret-santa-api/responses"
	"github.com/gorilla/mux"
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
		log.Print(err.Error())

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

	url, _ := mux.CurrentRoute(r).URL("id", strconv.Itoa(id))
	links := responses.Links{
		Self: url.String(),
	}

	lc.view.JSON(w, http.StatusOK, responses.Success{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Links:  links,
		Data:   item,
	})
}

// https://github.com/golang/go/wiki/CodeReviewComments#receiver-type
func (lc *ListController) GetLists(w http.ResponseWriter, r *http.Request) {
	data := make([]models.List, 0)

	list1 := models.List{
		Organiser: "Stephen McAuley",
		Email:     "steviebiddles@gmail.com",
		Amount:    50.00,
		Date:      time.Date(2016, time.December, 25, 15, 0, 0, 0, time.UTC),
		Location:  "Mums house",
		Notes:     "Try and not spoil it this year by telling anyone who you are buying for!",
		Created:   time.Now(),
		Updated:   time.Now(),
	}
	list2 := models.List{
		Organiser: "Sheena Hall",
		Email:     "sheena1hall@gmail.com",
		Amount:    99.99,
		Date:      time.Date(2016, time.December, 25, 10, 0, 0, 0, time.UTC),
		Created:   time.Now(),
		Updated:   time.Now(),
	}

	// prepare meta
	meta := map[string]interface{}{
		"pages": responses.Pages{
			Current:  1,
			Previous: 1,
			Next:     2,
			First:    1,
			Last:     2,
			Limit:    2,
			Total:    5,
		},
		"extra": "some extra stuff",
		"key":   "value",
	}

	// prepare links
	links := responses.Links{
		Self: "/lists/?page[pg]=1&page[limit]=2",
		Next: "/lists/?page[pg]=2&page[limit]=2",
		Last: "/lists/?page[pg]=5&page[limit]=2",
	}

	// prepare data
	data = append(data, list1, list2)

	lc.view.JSON(w, http.StatusOK, responses.Success{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Meta:   meta,
		Links:  links,
		Data:   data,
	})
}

func (lc *ListController) PostLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post list to here"))
}
