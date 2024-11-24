package api

import (
	"github.com/cargoreligion/booking/server/api/handler"
	"github.com/cargoreligion/booking/server/api/middleware"
	"github.com/cargoreligion/booking/server/infrastructure/db"
	"github.com/cargoreligion/booking/server/repository"
	"github.com/cargoreligion/booking/server/service"
	"github.com/gorilla/mux"
)

func NewRouter(dbc db.DbClient) *mux.Router {
	userRepo := repository.NewUserRepository(dbc)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	slotRepo := repository.NewSlotRepository(dbc)
	slotService := service.NewSlotService(slotRepo, userRepo)
	slotHandler := handler.NewSlotHandler(slotService)

	sessionRepo := repository.NewSessionFeedbackRepository(dbc)
	sessionService := service.NewSessionFeedbackService(sessionRepo, slotRepo, userRepo)
	sessionFeedbackHandler := handler.NewSessionFeedbackHandler(sessionService)
	r := mux.NewRouter()

	// Slot routes
	r.HandleFunc("/api/slots", slotHandler.CreateSlot).Methods("POST")
	r.HandleFunc("/api/slots/upcoming", slotHandler.GetUpcomingSlots).Methods("GET")
	r.HandleFunc("/api/slots/available", slotHandler.GetAvailableSlots).Methods("GET")
	r.HandleFunc("/api/slots/{id}/book", slotHandler.BookSlot).Methods("POST")
	r.HandleFunc("/api/students/bookings", slotHandler.GetUpcomingBookingsForStudent).Methods("GET")
	r.HandleFunc("/api/slots/{id}/details", slotHandler.GetSlotDetails).Methods("GET")

	// Session feedback routes
	r.HandleFunc("/api/session-feedback", sessionFeedbackHandler.CreateSessionFeedback).Methods("POST")
	r.HandleFunc("/api/session-feedback/past", sessionFeedbackHandler.GetPastSessionFeedbacks).Methods("GET")

	// User routes
	r.HandleFunc("/api/users", userHandler.GetAllUsers).Methods("GET")
	r.Use(middleware.WithUserID)
	return r
}
