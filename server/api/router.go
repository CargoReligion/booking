package api

import (
	"github.com/cargoreligion/booking/server/api/handler"
	"github.com/cargoreligion/booking/server/infrastructure/db"
	"github.com/cargoreligion/booking/server/repository"
	"github.com/cargoreligion/booking/server/service"
	"github.com/gorilla/mux"
)

func NewRouter(dbc db.DbClient) *mux.Router {
	userRepo := repository.NewUserRepository(dbc)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	r := mux.NewRouter()
	r.HandleFunc("/api/users", userHandler.GetAllUsers).Methods("GET")
	return r
}
