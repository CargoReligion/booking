package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cargoreligion/booking/server/api/middleware"
	"github.com/cargoreligion/booking/server/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type SessionFeedbackHandler struct {
	service *service.SessionFeedbackService
}

func NewSessionFeedbackHandler(service *service.SessionFeedbackService) *SessionFeedbackHandler {
	return &SessionFeedbackHandler{service: service}
}

func (h *SessionFeedbackHandler) CreateSessionFeedback(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserID(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	var req struct {
		SlotID       uuid.UUID `json:"slotId"`
		Satisfaction int       `json:"satisfaction"`
		Notes        string    `json:"notes"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Satisfaction < 1 || req.Satisfaction > 5 {
		http.Error(w, "Satisfaction must be between 1 and 5", http.StatusBadRequest)
		return
	}
	if err := h.service.CreateSessionFeedback(userID, req.SlotID, req.Satisfaction, req.Notes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *SessionFeedbackHandler) GetPastSessionFeedbacks(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserID(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	feedbacks, err := h.service.GetPastSessionFeedbacks(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(feedbacks)
}

func (h *SessionFeedbackHandler) GetStudentsWithSessionsByCoach(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID, err := middleware.GetUserID(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	students, err := h.service.GetStudentsWithSessionsByCoach(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(students)
}

func (h *SessionFeedbackHandler) GetSessionsForStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID, err := middleware.GetUserID(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	studentId, err := uuid.Parse(vars["studentId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sessions, err := h.service.GetSessionsForStudent(studentId, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(sessions)
}
