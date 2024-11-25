package model

import "github.com/google/uuid"

type SessionFeedback struct {
	ID           uuid.UUID `json:"id" db:"id"`
	SlotID       uuid.UUID `json:"slotId" db:"slot_id"`
	Satisfaction int       `json:"satisfaction" db:"satisfaction"`
	Notes        string    `json:"notes" db:"notes"`
}
