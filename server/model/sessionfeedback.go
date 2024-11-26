package model

import (
	"time"

	"github.com/google/uuid"
)

type SessionFeedback struct {
	ID           uuid.UUID `json:"id" db:"id"`
	SlotID       uuid.UUID `json:"slotId" db:"slot_id"`
	CoachId      uuid.UUID `json:"coachId" db:"coach_id"`
	StudentId    uuid.UUID `json:"studentId" db:"student_id"`
	Satisfaction int       `json:"satisfaction" db:"satisfaction"`
	Notes        string    `json:"notes" db:"notes"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
}
