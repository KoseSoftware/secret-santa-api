package controllers

import (
	"net/http"
	"time"

	"github.com/KoseSoftware/secret-santa-api/models"
	"github.com/unrolled/render"
)

type Links struct {
	Self string `json:"self"`
}

type JsonResponse struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Links  Links       `json:"links"`
	Data   interface{} `json:"data"`
}

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
	lists := make([]models.List, 0)

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

	lc.view.JSON(w, http.StatusOK, JsonResponse{
		Status: http.StatusText(http.StatusOK),
		Code:   http.StatusOK,
		Links: Links{
			Self: "/lists/",
		},
		Data: append(lists, list1, list2),
	})
}

func (lc *ListController) PostLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post list to here"))
}
