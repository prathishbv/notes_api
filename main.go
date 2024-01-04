package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/prathishbv/notes-api/config"
	"github.com/prathishbv/notes-api/controller"
	"github.com/prathishbv/notes-api/helper"
	"github.com/prathishbv/notes-api/model"
	"github.com/prathishbv/notes-api/repository"
	"github.com/prathishbv/notes-api/router"
	"github.com/prathishbv/notes-api/service"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	// Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})


	userRepository := repository.NewUsersRepositoryImpl(db)
    noteRepository := repository.NewNotesRepositoryImpl(db)

    authenticationService := service.NewAuthenticationServiceImpl(userRepository, validate)

    authenticationController := controller.NewAuthenticationController(authenticationService)
    usersController := controller.NewUsersController(userRepository)
    notesController := controller.NewNotesController(noteRepository)

    routes := router.NewRouter(userRepository, authenticationController, usersController, notesController)

	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
