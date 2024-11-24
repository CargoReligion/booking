package model

import (
	"time"

	"github.com/google/uuid"
)

type Slot struct {
	ID        uuid.UUID  `db:"id"`
	CoachID   uuid.UUID  `db:"coach_id"`
	StudentID *uuid.UUID `db:"student_id"`
	StartTime time.Time  `db:"start_time"`
	EndTime   time.Time  `db:"end_time"`
	Booked    bool       `db:"booked"`
}
