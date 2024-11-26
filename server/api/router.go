package api

import (
	"net/http"

	"github.com/cargoreligion/booking/server/api/handler"
	"github.com/cargoreligion/booking/server/api/middleware"
	"github.com/cargoreligion/booking/server/infrastructure/db"
	"github.com/cargoreligion/booking/server/repository"
	"github.com/cargoreligion/booking/server/service"
	"github.com/gorilla/handlers"
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
	r.HandleFunc("/api/slots/available/{coachId}", slotHandler.GetAvailableSlots).Methods("GET")
	r.HandleFunc("/api/slots/{id}/book", slotHandler.BookSlot).Methods("POST")
	r.HandleFunc("/api/students/bookings", slotHandler.GetUpcomingBookingsForStudent).Methods("GET")
	r.HandleFunc("/api/slots/{id}/details", slotHandler.GetSlotDetails).Methods("GET")

	// Session feedback routes
	r.HandleFunc("/api/session-feedback", sessionFeedbackHandler.CreateSessionFeedback).Methods("POST")
	r.HandleFunc("/api/session-feedback/past", sessionFeedbackHandler.GetPastSessionFeedbacks).Methods("GET")
	r.HandleFunc("/api/session-feedback/studentswithsessions", sessionFeedbackHandler.GetStudentsWithSessionsByCoach).Methods("GET")
	r.HandleFunc("/api/session-feedback/sessionsforstudent/{studentId}", sessionFeedbackHandler.GetSessionsForStudent).Methods("GET")

	// User routes
	r.HandleFunc("/api/users", userHandler.GetAllUsers).Methods("GET")

	// CORS configuration
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Allow all origins
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-User-Id"}),
	)
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Use(corsMiddleware)
	r.Use(middleware.WithUserID)
	return r
}
