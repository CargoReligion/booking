CREATE TABLE slot (
    id UUID PRIMARY KEY,
    coach_id UUID NOT NULL,
    student_id UUID,
    start_time TIMESTAMP WITH TIME ZONE NOT NULL,
    end_time TIMESTAMP WITH TIME ZONE NOT NULL,
    booked BOOLEAN NOT NULL DEFAULT FALSE
);

ALTER TABLE slots 
ADD CONSTRAINT fk_slots_coach 
FOREIGN KEY (coach_id) REFERENCES users(id);

ALTER TABLE slots 
ADD CONSTRAINT fk_slots_student 
FOREIGN KEY (student_id) REFERENCES users(id);

ALTER TABLE slots
ADD CONSTRAINT check_slot_times
CHECK (end_time > start_time);