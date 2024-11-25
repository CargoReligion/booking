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

func (r *SlotRepository) GetUpcomingSlots(coachID uuid.UUID, offset, pagesize int) ([]model.Slot, int, error) {
	var totalCount int
	query := `SELECT COUNT(*) FROM slot WHERE coach_id = $1 AND start_time > NOW()`
	err := r.dbc.GetSingleEntity(&totalCount, query, coachID)
	if err != nil {
		return nil, 0, err
	}
	var slots []model.Slot
	query = `
		SELECT 
			* 
		FROM 
			slot 
		WHERE 
			coach_id = $1 AND 
			start_time > NOW() 
		ORDER BY 
			start_time ASC
		LIMIT $2 OFFSET $3`
	err = r.dbc.Select(&slots, query, coachID, pagesize, offset)
	return slots, totalCount, err
}

func (r *SlotRepository) GetAvailableSlots(coachID uuid.UUID, offset, pagesize int) ([]model.Slot, int, error) {
	var totalCount int
	query := `SELECT COUNT(*) FROM slot WHERE coach_id = $1 AND booked = false AND start_time > NOW()`
	err := r.dbc.GetSingleEntity(&totalCount, query, coachID)
	if err != nil {
		return nil, 0, err
	}
	var slots []model.Slot
	query = `
		SELECT 
			* 
		FROM 
			slot 
		WHERE 
			coach_id = $1 AND
			booked = false AND 
			start_time > NOW() 
		ORDER BY 
			start_time ASC
			LIMIT $2 OFFSET $3
		`
	err = r.dbc.Select(&slots, query, coachID, pagesize, offset)
	return slots, totalCount, err
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

func (r *SlotRepository) GetUpcomingBookingsForStudent(studentID uuid.UUID, offset, pagesize int) ([]model.Slot, int, error) {
	var totalCount int
	query := `SELECT COUNT(*) FROM slot WHERE student_id = $1 AND booked = true AND start_time > NOW()`
	err := r.dbc.GetSingleEntity(&totalCount, query, studentID)
	if err != nil {
		return nil, 0, err
	}
	var slots []model.Slot
	query = `
		SELECT s.*, u.name as coach_name
		FROM slot s
		JOIN stepful_user u ON s.coach_id = u.id
		WHERE s.student_id = $1 
		AND s.start_time > $2
		AND s.booked = true
		ORDER BY s.start_time ASC
		LIMIT $3 OFFSET $4
	`
	err = r.dbc.Select(&slots, query, studentID, time.Now(), pagesize, offset)
	return slots, totalCount, err
}

func (r *SlotRepository) GetSlotDetails(slotID uuid.UUID) (*model.SlotDetails, error) {
	var slotDetails model.SlotDetails
	query := `
        SELECT 
            s.*,
            c.name AS coach_name,
            c.phone_number AS coach_phone_number,
            st.name AS student_name,
            st.phone_number AS student_phone_number
        FROM slot s
        JOIN stepful_user c ON s.coach_id = c.id
        LEFT JOIN stepful_user st ON s.student_id = st.id
        WHERE s.id = $1
    `
	err := r.dbc.GetSingleEntity(&slotDetails, query, slotID)
	if err != nil {
		return nil, err
	}
	return &slotDetails, nil
}
