package model

import "github.com/google/uuid"

type SessionFeedback struct {
	ID           uuid.UUID `db:"id"`
	SlotID       uuid.UUID `db:"slot_id"`
	Satisfaction int       `db:"satisfaction"`
	Notes        string    `db:"notes"`
}
