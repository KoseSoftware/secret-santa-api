package controllers

import (
	"net/http"

	"github.com/KoseSoftware/secret-santa-api/models"
	"github.com/unrolled/render"
)

type Self struct {
	Href string `json:"href,omitempty"`
}
type Next struct {
	Href string `json:"href,omitempty"`
}

// http://stackoverflow.com/questions/33447334/golang-json-marshal-how-to-omit-empty-nested-struct
type Links struct {
	Self *Self `json:"self,omitempty"`
	Next *Next `json:"next,omitempty"`
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

func NewListController(r *render.Render) *ListController {
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
	}
	list2 := models.List{
		Organiser: "Sheena Hall",
		Email:     "sheena1hall@gmail.com",
	}

	lc.view.JSON(w, http.StatusOK, JsonResponse{
		Status: http.StatusText(http.StatusOK),
		Code:   http.StatusOK,
		Links: Links{
			Self: &Self{
				Href: "/lists/",
			},
		},
		Data: append(lists, list1, list2),
	})
}

func (lc *ListController) PostLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post list to here"))
}
