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
	query := `INSERT INTO session_feedback (id, slot_id, coach_id, student_id, satisfaction, notes, created_at) 
			  VALUES (:id, :slot_id, :coach_id, :student_id, :satisfaction, :notes, :created_at)`
	_, err := r.dbc.NamedExec(query, feedback)
	return err
}

func (r *SessionFeedbackRepository) GetPastSessionFeedback(coachID uuid.UUID) ([]model.SessionFeedback, error) {
	var feedbacks []model.SessionFeedback
	query := `SELECT sf.* FROM session_feedback sf
			  JOIN slot s ON sf.slot_id = s.id
			  WHERE s.coach_id = $1 AND ((s.end_time < NOW() AND s.status = 'active') OR s.status = 'ended')
			  ORDER BY s.start_time DESC`
	err := r.dbc.Select(&feedbacks, query, coachID)
	return feedbacks, err
}

func (r *SessionFeedbackRepository) GetStudentsWithSessionsByCoach(coachID uuid.UUID) ([]model.User, error) {
	var students []model.User
	query := `
			SELECT DISTINCT u.* 
			FROM stepful_user u
			JOIN session_feedback sf ON u.id = sf.student_id
			WHERE sf.coach_id = $1
			ORDER BY u.name ASC
			`
	err := r.dbc.Select(&students, query, coachID)
	return students, err
}

func (r *SessionFeedbackRepository) GetSessionsForStudent(studentId, coachId uuid.UUID) ([]model.SessionFeedback, error) {
	var feedbacks []model.SessionFeedback
	query := `
			SELECT sf.*
			FROM session_feedback sf
			WHERE sf.student_id = $1 AND sf.coach_id = $2
			ORDER BY sf.created_at DESC
			`
	err := r.dbc.Select(&feedbacks, query, studentId, coachId)
	return feedbacks, err
}
