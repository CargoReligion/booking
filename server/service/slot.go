package service

import (
	"fmt"
	"time"

	"github.com/cargoreligion/booking/server/model"
	"github.com/cargoreligion/booking/server/repository"
	"github.com/google/uuid"
)

type SlotService struct {
	slotRepo *repository.SlotRepository
	userRepo *repository.UserRepository
}

func NewSlotService(slotRepo *repository.SlotRepository, userRepo *repository.UserRepository) *SlotService {
	return &SlotService{
		slotRepo: slotRepo,
		userRepo: userRepo,
	}
}

func (s *SlotService) CreateSlot(coachID uuid.UUID, startTime time.Time) (uuid.UUID, error) {
	localLocation := time.Local
	localStartTime := startTime.In(localLocation)
	// Check if the slot is in the past
	now := time.Now()
	if localStartTime.Before(now) {
		return uuid.Nil, fmt.Errorf("cannot create a slot in the past")
	}

	// Check if the start time is at a 15-minute increment
	if localStartTime.Minute()%15 != 0 || localStartTime.Second() != 0 || localStartTime.Nanosecond() != 0 {
		return uuid.Nil, fmt.Errorf("slot must start at 15-minute increments (e.g., 9:00, 9:15, 9:30, 9:45)")
	}

	// Check if the slot is between 9 AM and 5 PM
	startHour := localStartTime.Hour()
	if startHour < 9 || startHour >= 17 {
		return uuid.Nil, fmt.Errorf("slots must be between 9 AM and 5 PM")
	}

	// Calculate end time (2 hours after start time)
	endTime := localStartTime.Add(2 * time.Hour)

	// Check if the end time is after 5 PM
	if endTime.Hour() >= 17 && endTime.Minute() > 0 {
		return uuid.Nil, fmt.Errorf("slots must end by 5 PM")
	}

	// Fetch the user
	user, err := s.userRepo.GetUserByID(coachID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error fetching user: %w", err)
	}

	// Check if the user is a coach
	if user.Role != model.RoleCoach {
		return uuid.Nil, fmt.Errorf("only coaches can create slots")
	}

	// Check for overlapping slots
	hasOverlap, err := s.slotRepo.HasOverlappingSlot(coachID, localStartTime, endTime)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error checking for overlapping slots: %w", err)
	}
	if hasOverlap {
		return uuid.Nil, fmt.Errorf("slot overlaps with an existing slot")
	}

	// Create the slot
	slot := model.Slot{
		ID:        uuid.New(),
		CoachID:   coachID,
		StartTime: localStartTime.UTC(),
		EndTime:   endTime.UTC(),
		Booked:    false,
	}

	// Save the slot
	id, err := s.slotRepo.CreateSlot(slot)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error creating slot: %w", err)
	}

	return id, nil
}

func (s *SlotService) GetUpcomingSlots(userID uuid.UUID, page, pageSize int) ([]model.Slot, int, error) {
	// First, check if the user is a coach
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, 0, fmt.Errorf("error fetching user: %w", err)
	}

	if user.Role != model.RoleCoach {
		return nil, 0, &ErrNotCoach{UserID: userID.String()}
	}

	offset := (page - 1) * pageSize
	// If the user is a coach, proceed to fetch upcoming slots
	paginatedSlots, totalSlots, err := s.slotRepo.GetUpcomingSlots(userID, offset, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("error fetching upcoming slots: %w", err)
	}
	if paginatedSlots == nil {
		paginatedSlots = []model.Slot{} // Return an empty slice instead of nil
	}
	return paginatedSlots, totalSlots, nil
}

func (s *SlotService) GetAvailableSlots(coachId uuid.UUID, page, pageSize int) ([]model.Slot, int, error) {
	user, err := s.userRepo.GetUserByID(coachId)
	if err != nil {
		return nil, 0, fmt.Errorf("error fetching user: %w", err)
	}

	if user.Role != model.RoleCoach {
		return nil, 0, &ErrNotCoach{UserID: coachId.String()}
	}

	offset := (page - 1) * pageSize
	paginatedSlots, totalSlots, err := s.slotRepo.GetAvailableSlots(coachId, offset, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("error fetching available slots: %w", err)
	}
	if paginatedSlots == nil {
		paginatedSlots = []model.Slot{} // Return an empty slice instead of nil
	}
	return paginatedSlots, totalSlots, nil
}

func (s *SlotService) BookSlot(slotID, studentID uuid.UUID) error {
	// Fetch the user
	user, err := s.userRepo.GetUserByID(studentID)
	if err != nil {
		return fmt.Errorf("error fetching user: %w", err)
	}

	// Check if the user is a student
	if user.Role != model.RoleStudent {
		return fmt.Errorf("only students can book slots")
	}

	// Fetch the slot
	slot, err := s.slotRepo.GetSlotByID(slotID)
	if err != nil {
		return fmt.Errorf("error fetching slot: %w", err)
	}

	// Check if the slot is already booked
	if slot.Booked {
		return fmt.Errorf("slot is already booked")
	}

	// Check if the slot is in the past
	if slot.StartTime.Before(time.Now()) {
		return fmt.Errorf("cannot book a slot that has already begun")
	}

	if slot.EndTime.Before(time.Now()) {
		return fmt.Errorf("cannot book a slot that has already expired")
	}

	// Check for overlapping bookings
	hasOverlap, err := s.slotRepo.HasOverlappingBooking(studentID, slot.StartTime, slot.EndTime)
	if err != nil {
		return fmt.Errorf("error checking for overlapping bookings: %w", err)
	}
	if hasOverlap {
		return fmt.Errorf("booking overlaps with an existing booking")
	}

	// Book the slot
	slot.StudentID = &studentID
	slot.Booked = true

	err = s.slotRepo.UpdateSlot(*slot)
	if err != nil {
		return fmt.Errorf("error updating slot: %w", err)
	}

	return nil
}

func (s *SlotService) GetUpcomingBookingsForStudent(studentID uuid.UUID, page, pageSize int) ([]model.Slot, int, error) {
	// First, check if the user is a student
	user, err := s.userRepo.GetUserByID(studentID)
	if err != nil {
		return nil, 0, fmt.Errorf("error fetching user: %w", err)
	}

	if user.Role != model.RoleStudent {
		return nil, 0, &ErrNotStudent{UserID: studentID.String()}
	}
	offset := (page - 1) * pageSize
	// If the user is a student, proceed to fetch upcoming bookings
	paginatedSlots, totalCount, err := s.slotRepo.GetUpcomingBookingsForStudent(studentID, offset, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("error fetching upcoming bookings: %w", err)
	}
	if paginatedSlots == nil {
		paginatedSlots = []model.Slot{} // Return an empty slice instead of nil
	}

	return paginatedSlots, totalCount, nil
}

func (s *SlotService) GetSlotDetails(userID, slotID uuid.UUID) (*model.SlotDetails, error) {
	// Fetch the slot details
	slotDetails, err := s.slotRepo.GetSlotDetails(slotID)
	if err != nil {
		return nil, fmt.Errorf("error fetching slot details: %w", err)
	}

	// Check if the user is either the coach or the student for this slot
	if slotDetails.CoachID != userID && (slotDetails.StudentID == nil || *slotDetails.StudentID != userID) {
		return nil, &ErrNotAuthorized{UserID: userID.String(), Action: "view slot details"}
	}

	// If the slot is not booked, remove student information
	if !slotDetails.Booked {
		slotDetails.StudentID = nil
		slotDetails.StudentName = ""
		slotDetails.StudentPhoneNumber = ""
	}

	return slotDetails, nil
}
