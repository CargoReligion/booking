package repository

import (
	"time"

	"github.com/cargoreligion/booking/server/infrastructure/db"
	"github.com/cargoreligion/booking/server/model"
	"github.com/google/uuid"
)

type SlotRepository struct {
	dbc db.DbClient
}

func NewSlotRepository(dbc db.DbClient) *SlotRepository {
	return &SlotRepository{dbc: dbc}
}

func (r *SlotRepository) CreateSlot(slot model.Slot) (uuid.UUID, error) {
	query := `INSERT INTO slot (id, coach_id, start_time, end_time, booked) 
			  VALUES (:id, :coach_id, :start_time, :end_time, :booked)
			  RETURNING id`
	var id uuid.UUID
	err := r.dbc.NamedGetSingleEntity(&id, query, slot)
	return id, err
}

func (r *SlotRepository) GetUpcomingSlots(coachID uuid.UUID) ([]model.Slot, error) {
	var slots []model.Slot
	query := `SELECT * FROM slot WHERE coach_id = $1 AND start_time > NOW() ORDER BY start_time ASC`
	err := r.dbc.Select(&slots, query, coachID)
	return slots, err
}

func (r *SlotRepository) GetAvailableSlots() ([]model.Slot, error) {
	var slots []model.Slot
	query := `SELECT * FROM slot WHERE booked = false AND start_time > NOW() ORDER BY start_time ASC`
	err := r.dbc.Select(&slots, query)
	return slots, err
}

func (r *SlotRepository) GetSlotByID(id uuid.UUID) (*model.Slot, error) {
	var slot model.Slot
	query := `SELECT * FROM slot WHERE id = $1`
	err := r.dbc.GetSingleEntity(&slot, query, id)
	if err != nil {
		return nil, err
	}
	return &slot, nil
}

func (r *SlotRepository) UpdateSlot(slot model.Slot) error {
	query := `UPDATE slot SET student_id = :student_id, booked = :booked WHERE id = :id`
	_, err := r.dbc.NamedExec(query, slot)
	return err
}

func (r *SlotRepository) BookSlot(slotID, studentID uuid.UUID) error {
	query := `UPDATE slot SET student_id = $1, booked = true WHERE id = $2 AND booked = false`
	_, err := r.dbc.NamedExec(query, map[string]interface{}{
		"slot_id":    slotID,
		"student_id": studentID,
	})
	return err
}

func (r *SlotRepository) HasOverlappingSlot(coachID uuid.UUID, startTime, endTime time.Time) (bool, error) {
	var count int
	query := `
		SELECT COUNT(*) 
		FROM slot
		WHERE coach_id = $1 
		AND (
			(start_time <= $2 AND end_time > $2) OR
			(start_time < $3 AND end_time >= $3) OR
			(start_time >= $2 AND end_time <= $3)
		)`
	err := r.dbc.GetSingleEntity(&count, query, coachID, startTime, endTime)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *SlotRepository) HasOverlappingBooking(studentID uuid.UUID, startTime, endTime time.Time) (bool, error) {
	var count int
	query := `
		SELECT COUNT(*) 
		FROM slot
		WHERE student_id = $1 
		AND booked = true
		AND (
			(start_time <= $2 AND end_time > $2) OR
			(start_time < $3 AND end_time >= $3) OR
			(start_time >= $2 AND end_time <= $3)
		)`
	err := r.dbc.GetSingleEntity(&count, query, studentID, startTime, endTime)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
