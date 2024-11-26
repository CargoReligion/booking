ALTER TABLE session_feedback
ADD COLUMN coach_id UUID NOT NULL,
ADD COLUMN student_id UUID NOT NULL;

ALTER TABLE session_feedback
ADD CONSTRAINT fk_session_feedback_coach
FOREIGN KEY (coach_id) REFERENCES stepful_user(id);

ALTER TABLE session_feedback
ADD CONSTRAINT fk_session_feedback_student
FOREIGN KEY (student_id) REFERENCES stepful_user(id);