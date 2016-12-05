package controllers

import "net/http"

type ListsController struct {
}

func NewListsController() *ListsController {
	return &ListsController{}
}

// https://github.com/golang/go/wiki/CodeReviewComments#receiver-type
func (lc *ListsController) GetLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("List for a single user (email address)"))
}

func (lc *ListsController) PostLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post list to here"))
}
