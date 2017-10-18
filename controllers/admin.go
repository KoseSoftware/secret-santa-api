package controllers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type AdminController struct {
}

func NewAdminController() *AdminController {
	return &AdminController{}
}

func (ac *AdminController) Index(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	claims := user.(*jwt.Token).Claims

	fmt.Println(claims)
	fmt.Fprintf(w, "This is an authenticated request")
}
