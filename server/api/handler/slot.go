package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cargoreligion/booking/server/api/middleware"
	"github.com/cargoreligion/booking/server/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type SlotHandler struct {
	service *service.SlotService
}

func NewSlotHandler(service *service.SlotService) *SlotHandler {
	return &SlotHandler{service: service}
}

func (h *SlotHandler) CreateSlot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID, err := middleware.GetUserID(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	var req struct {
		StartTime time.Time `json:"start_time"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateSlot(userID, req.StartTime)
	if err != nil {
		switch err.Error() {
		case "only coaches can create slots":
			http.Error(w, err.Error(), http.StatusForbidden)
		case "slot overlaps with an existing slot":
			http.Error(w, err.Error(), http.StatusConflict)
		case "cannot create a slot in the past":
			http.Error(w, err.Error(), http.StatusBadRequest)
		case "slot must start at 15-minute increments (e.g., 9:00, 9:15, 9:30, 9:45)":
			http.Error(w, err.Error(), http.StatusBadRequest)
		case "slots must be between 9 AM and 5 PM":
			http.Error(w, err.Error(), http.StatusBadRequest)
		case "slots must end by 5 PM":
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	response := struct {
		ID uuid.UUID `json:"id"`
	}{
		ID: id,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *SlotHandler) GetUpcomingSlots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID, err := middleware.GetUserID(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	slots, err := h.service.GetUpcomingSlots(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(slots)
}

func (h *SlotHandler) GetAvailableSlots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	slots, err := h.service.GetAvailableSlots()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(slots)
}

func (h *SlotHandler) BookSlot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID, err := middleware.GetUserID(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	slotID, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid slot ID", http.StatusBadRequest)
		return
	}
	if err := h.service.BookSlot(slotID, userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *SlotHandler) GetUpcomingBookingsForStudent(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserID(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	slots, err := h.service.GetUpcomingBookingsForStudent(userID)
	if err != nil {
		fmt.Println(err.Error())
		var errNotStudent *service.ErrNotStudent
		if errors.As(err, &errNotStudent) {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(slots)
}

func (h *SlotHandler) GetSlotDetails(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserID(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Extract slot ID from the URL path
	vars := mux.Vars(r)
	slotID, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "Invalid slot ID", http.StatusBadRequest)
		return
	}

	slotDetails, err := h.service.GetSlotDetails(userID, slotID)
	if err != nil {
		var errNotAuthorized *service.ErrNotAuthorized
		if errors.As(err, &errNotAuthorized) {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(slotDetails)
}