package repository

import (
	"github.com/cargoreligion/booking/server/infrastructure/db"
	"github.com/cargoreligion/booking/server/model"
	"github.com/google/uuid"
)

type SessionFeedbackRepository struct {
	dbc db.DbClient
}

func NewSessionFeedbackRepository(dbc db.DbClient) *SessionFeedbackRepository {
	return &SessionFeedbackRepository{dbc: dbc}
}

func (r *SessionFeedbackRepository) CreateSessionFeedback(feedback model.SessionFeedback) error {
	query := `INSERT INTO session_feedback (id, slot_id, satisfaction, notes, created_at) 
			  VALUES (:id, :slot_id, :satisfaction, :notes, NOW())`
	_, err := r.dbc.NamedExec(query, feedback)
	return err
}

func (r *SessionFeedbackRepository) GetPastSessionFeedback(coachID uuid.UUID) ([]model.SessionFeedback, error) {
	var feedbacks []model.SessionFeedback
	query := `SELECT sf.* FROM session_feedback sf
			  JOIN slot s ON sf.slot_id = s.id
			  WHERE s.coach_id = $1 AND s.end_time < NOW()
			  ORDER BY s.start_time DESC`
	err := r.dbc.Select(&feedbacks, query, coachID)
	return feedbacks, err
}
