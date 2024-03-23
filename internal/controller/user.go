package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/albiosz/honeycombs"
	"github.com/albiosz/honeycombs/internal/config/errs"
)

type User struct {
	userService honeycombs.UserService
}

func NewUser(userService honeycombs.UserService) *User {
	return &User{
		userService: userService,
	}
}

func (u *User) FindByID(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := u.userService.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, errs.ErrResourceNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user.Password = ""

	switch r.Header.Get("Accept") {
	case "application/json":
		w.Header().Set("Content-type", "application/json")
		if err := json.NewEncoder(w).Encode(user); err != nil {
			log.Println(err)
			return
		}
	}
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	var userToCreate honeycombs.User
	err := json.NewDecoder(r.Body).Decode(&userToCreate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	createdUser, err := u.userService.Create(&userToCreate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch r.Header.Get("Accept") {
	case "application/json":
		w.Header().Set("Content-type", "application/json")
		if err := json.NewEncoder(w).Encode(createdUser); err != nil {
			log.Println(err)
			return
		}
	}
}

func (u *User) Update(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var updates honeycombs.UserUpdate
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedUser, err := u.userService.Update(uint(id), updates)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch r.Header.Get("Accept") {
	case "application/json":
		w.Header().Set("Content-type", "application/json")
		if err := json.NewEncoder(w).Encode(updatedUser); err != nil {
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (u *User) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := u.userService.Delete(uint(id)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
