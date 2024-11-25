package model

import (
	"time"

	"github.com/google/uuid"
)

type Slot struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	CoachID   uuid.UUID  `json:"coachId" db:"coach_id"`
	CoachName string     `json:"coachName" db:"coach_name"`
	StudentID *uuid.UUID `json:"studentId" db:"student_id"`
	StartTime time.Time  `json:"startTime" db:"start_time"`
	EndTime   time.Time  `json:"endTime" db:"end_time"`
	Booked    bool       `json:"booked" db:"booked"`
}

type SlotDetails struct {
	Slot
	CoachName          string `db:"coach_name" json:"coachName"`
	CoachPhoneNumber   string `db:"coach_phone_number" json:"coachPhoneNumber"`
	StudentName        string `db:"student_name" json:"studentName,omitempty"`
	StudentPhoneNumber string `db:"student_phone_number" json:"studentPhoneNumber,omitempty"`
}
