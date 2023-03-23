package main

import (
	"demoG/config"
	"demoG/controller"
	"demoG/helper"
	"demoG/model"
	"demoG/repository"
	"demoG/router"
	"demoG/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})

	// Repository
	usersRepository := repository.NewUsersREpositoryImpl(db)

	// Service
	usersService := service.NewUsersServiceImpl(usersRepository, validate)

	// Controller
	usersController := controller.NewUsersController(usersService)

	// Router
	routes := router.NewRouter(usersController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
