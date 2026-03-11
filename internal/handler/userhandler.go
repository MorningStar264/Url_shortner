package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MorningStar264/Url_shortner/internal/helper"
	"github.com/MorningStar264/Url_shortner/internal/model"
	"github.com/MorningStar264/Url_shortner/internal/repository"
	"github.com/MorningStar264/Url_shortner/internal/server"
)

type UserHandler struct {
	server       *server.Server
	repositories *repository.Repositories
}

func NewUserHandler(srv *server.Server, repo *repository.Repositories) *UserHandler {
	return &UserHandler{
		server:       srv,
		repositories: repo,
	}
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var u model.User

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	// ensure the unknown field are not there
	// decoder.DisallowUnknownFields()
	// decoding the data
	err := decoder.Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// hash the password and store
	hashPassword, _ := helper.HashPassword(u.PasswordHash)
	u.PasswordHash = hashPassword

	fmt.Fprintf(w, "Received: %+v", u)
	h.repositories.UserMethods.CreateUser(u)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var u model.User

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	// ensure the unknown field are not there
	decoder.DisallowUnknownFields()
	// decoding the data
	err := decoder.Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := h.repositories.UserMethods.GetUser(u)
	if !helper.CheckPassword(u.PasswordHash, user.PasswordHash) {
		fmt.Fprintf(w, "Authentication failed")
	} else {
		jwt_token,_:=helper.CreateToken(u.Username)
		fmt.Fprintf(w, "Logged in Succesfully\n")
		fmt.Fprintf(w,"%s", jwt_token)
	}
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var u model.User

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	// ensure the unknown field are not there
	decoder.DisallowUnknownFields()
	// decoding the data
	err := decoder.Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if u.PasswordHash != "" {
		hashPassword, _ := helper.HashPassword(u.PasswordHash)
		u.PasswordHash = hashPassword
	}
	h.repositories.UserMethods.UpdateUser(u)
	fmt.Println("update user")
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var u model.User

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	// ensure the unknown field are not there
	decoder.DisallowUnknownFields()
	// decoding the data
	err := decoder.Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.repositories.UserMethods.DeleteUser(u)
	fmt.Println("delete user")
}
