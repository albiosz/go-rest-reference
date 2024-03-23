package main

import (
	"net/http"

	"github.com/albiosz/honeycombs/internal/controller"
	"github.com/albiosz/honeycombs/internal/database"
	"github.com/albiosz/honeycombs/internal/repository"
	"github.com/albiosz/honeycombs/internal/service"
)

var userController *controller.User

func injectDependencies(db *database.DB) {

	// repositories
	userRepo := repository.NewUser(db)

	// services
	userService := service.NewUser(userRepo)

	// controllers
	userController = controller.NewUser(userService)
}

func getRouter(db *database.DB) *http.ServeMux {
	injectDependencies(db)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/{id}", userController.FindByID)
	mux.HandleFunc("POST /users", userController.Create)
	mux.HandleFunc("PATCH /users/{id}", userController.Update)
	mux.HandleFunc("DELETE /users/{id}", userController.Delete)

	return mux
}
