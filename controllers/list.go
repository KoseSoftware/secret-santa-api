package controllers

import (
	"net/http"
	"time"

	"github.com/KoseSoftware/secret-santa-api/models"
	"github.com/KoseSoftware/secret-santa-api/responses"
	"github.com/unrolled/render"
)

type ListController struct {
	view *render.Render
}

func NewListsController(r *render.Render) *ListController {
	return &ListController{
		view: r,
	}
}

// https://github.com/golang/go/wiki/CodeReviewComments#receiver-type
func (lc *ListController) GetLists(w http.ResponseWriter, r *http.Request) {
	meta := make(map[string]interface{}, 0)
	data := make([]models.List, 0)

	list1 := models.List{
		Organiser: "Stephen McAuley",
		Email:     "steviebiddles@gmail.com",
		Amount:    50.00,
		Date:      time.Date(2016, time.December, 25, 15, 0, 0, 0, time.UTC),
		Location:  "Mums house",
		Notes:     "Try and not spoil it this year by telling anyone who you are buying for!",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	list2 := models.List{
		Organiser: "Sheena Hall",
		Email:     "sheena1hall@gmail.com",
		Amount:    99.99,
		Date:      time.Date(2016, time.December, 25, 10, 0, 0, 0, time.UTC),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// prepare meta
	meta["current-page"] = 1
	meta["total-pages"] = 10
	meta["extra"] = "some extra meta information"

	// prepare links
	links := responses.Links{
		Self: "/lists/?page[pg]=1&page[limit]=2",
		Next: "/lists/?page[pg]=2&page[limit]=2",
		Last: "/lists/?page[pg]=5&page[limit]=2",
	}

	// prepare data
	data = append(data, list1, list2)

	lc.view.JSON(w, http.StatusOK, responses.Success{
		Status: http.StatusOK,
		Title:  http.StatusText(http.StatusOK),
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
