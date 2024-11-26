package service

import (
	"fmt"
	"time"

	"github.com/cargoreligion/booking/server/model"
	"github.com/cargoreligion/booking/server/repository"
	"github.com/google/uuid"
)

type SessionFeedbackService struct {
	sessionFeedbackRepo *repository.SessionFeedbackRepository
	slotRepo            *repository.SlotRepository
	userRepo            *repository.UserRepository
}

func NewSessionFeedbackService(
	sessionFeedbackRepo *repository.SessionFeedbackRepository,
	slotRepo *repository.SlotRepository,
	userRepo *repository.UserRepository,
) *SessionFeedbackService {
	return &SessionFeedbackService{
		sessionFeedbackRepo: sessionFeedbackRepo,
		slotRepo:            slotRepo,
		userRepo:            userRepo,
	}
}

func (s *SessionFeedbackService) CreateSessionFeedback(coachID uuid.UUID, slotID uuid.UUID, satisfaction int, notes string) error {
	// Check if the user is a coach
	user, err := s.userRepo.GetUserByID(coachID)
	if err != nil {
		return fmt.Errorf("error fetching user: %w", err)
	}
	if user.Role != model.RoleCoach {
		return &ErrNotAuthorized{UserID: coachID.String(), Action: "create session feedback"}
	}

	// Check if the slot is assigned to this coach
	slot, err := s.slotRepo.GetSlotByID(slotID)
	if err != nil {
		return fmt.Errorf("error fetching slot: %w", err)
	}
	if slot.CoachID != coachID {
		return &ErrSlotNotAssignedToCoach{SlotID: slotID.String(), CoachID: coachID.String()}
	}

	// Create the session feedback
	feedback := model.SessionFeedback{
		ID:           uuid.New(),
		SlotID:       slotID,
		CoachId:      coachID,
		StudentId:    *slot.StudentID,
		Satisfaction: satisfaction,
		Notes:        notes,
		CreatedAt:    time.Now(),
	}

	err = s.sessionFeedbackRepo.CreateSessionFeedback(feedback)
	if err != nil {
		return fmt.Errorf("error creating session feedback: %w", err)
	}

	return nil
}

func (s *SessionFeedbackService) GetPastSessionFeedbacks(coachID uuid.UUID) ([]model.SessionFeedback, error) {
	// Check if the user is a coach
	user, err := s.userRepo.GetUserByID(coachID)
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %w", err)
	}
	if user.Role != model.RoleCoach {
		return nil, &ErrNotAuthorized{UserID: coachID.String(), Action: "retrieve session feedback"}
	}

	// Fetch the session feedback for this coach
	feedbacks, err := s.sessionFeedbackRepo.GetPastSessionFeedback(coachID)
	if err != nil {
		return nil, fmt.Errorf("error fetching session feedback: %w", err)
	}

	return feedbacks, nil
}

func (s *SessionFeedbackService) GetStudentsWithSessionsByCoach(coachID uuid.UUID) ([]model.User, error) {
	// Check if the user is a coach
	user, err := s.userRepo.GetUserByID(coachID)
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %w", err)
	}
	if user.Role != model.RoleCoach {
		return nil, &ErrNotAuthorized{UserID: coachID.String(), Action: "retrieve students with sessions"}
	}

	// Fetch the students with sessions for this coach
	students, err := s.sessionFeedbackRepo.GetStudentsWithSessionsByCoach(coachID)
	if err != nil {
		return nil, fmt.Errorf("error fetching students with sessions: %w", err)
	}

	return students, nil
}

func (s *SessionFeedbackService) GetSessionsForStudent(studentID, coachID uuid.UUID) ([]model.SessionFeedback, error) {
	// Check if the user is a coach
	user, err := s.userRepo.GetUserByID(coachID)
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %w", err)
	}
	if user.Role != model.RoleCoach {
		return nil, &ErrNotAuthorized{UserID: coachID.String(), Action: "retrieve sessions for student"}
	}
	// Fetch the sessions for this student and coach
	sessions, err := s.sessionFeedbackRepo.GetSessionsForStudent(studentID, coachID)
	if err != nil {
		return nil, fmt.Errorf("error fetching sessions for student: %w", err)
	}
	return sessions, nil
}
