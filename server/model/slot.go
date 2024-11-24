package model

import (
	"time"

	"github.com/google/uuid"
)

type Slot struct {
	ID        uuid.UUID  `db:"id"`
	CoachID   uuid.UUID  `db:"coach_id"`
	CoachName string     `db:"coach_name"`
	StudentID *uuid.UUID `db:"student_id"`
	StartTime time.Time  `db:"start_time"`
	EndTime   time.Time  `db:"end_time"`
	Booked    bool       `db:"booked"`
}

type SlotDetails struct {
	Slot
	CoachName          string `db:"coach_name" json:"coach_name"`
	CoachPhoneNumber   string `db:"coach_phone_number" json:"coach_phone_number"`
	StudentName        string `db:"student_name" json:"student_name,omitempty"`
	StudentPhoneNumber string `db:"student_phone_number" json:"student_phone_number,omitempty"`
}
