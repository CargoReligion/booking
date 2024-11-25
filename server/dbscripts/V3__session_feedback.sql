CREATE TABLE session_feedback (
    id UUID PRIMARY KEY,
    slot_id UUID NOT NULL,
    satisfaction INT NOT NULL,
    notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT check_satisfaction_range
        CHECK (satisfaction >= 1 AND satisfaction <= 5)
);

ALTER TABLE session_feedback
ADD CONSTRAINT fk_slot
FOREIGN KEY (slot_id) 
REFERENCES slot(id);

-- Index for faster lookups by slot_id
CREATE INDEX idx_session_feedback_slot_id ON session_feedback(slot_id);