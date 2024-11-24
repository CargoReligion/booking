CREATE TABLE stepful_user (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    user_role TEXT NOT NULL
);

INSERT INTO stepful_user (id, name, phone_number, user_role) VALUES 
  ('f47ac10b-58cc-4372-a567-0e02b2c3d479', 'John Smith', '555-0101', 'coach'),
  ('b9eb36bd-5388-4c89-91d4-c2710c45d42a', 'Sarah Johnson', '555-0102', 'coach'),
  ('6ba7b810-9dad-11d1-80b4-00c04fd430c8', 'Michael Chen', '555-0103', 'coach'),
  ('550e8400-e29b-41d4-a716-446655440000', 'Emma Davis', '555-0104', 'coach'),
  ('67e55044-10b1-426f-9247-bb680e5fe0c8', 'Robert Wilson', '555-0105', 'coach'),
  ('8c725a10-0abd-4712-a88e-c944c6273806', 'Alice Brown', '555-0201', 'student'),
  ('91a85a9e-1d1d-4d5a-8f6b-4ecc1934aa3d', 'David Lee', '555-0202', 'student'),
  ('d5f38b87-7b1a-4e87-b6e9-8e8f3d9d5c5b', 'Maria Garcia', '555-0203', 'student'),
  ('c2e15c48-97d9-4d5c-b7d7-e2b1b4f5c6d7', 'James Taylor', '555-0204', 'student'),
  ('4ce28ee2-4d08-44f3-96dd-5d796fdafb4a', 'Sophie Martin', '555-0205', 'student');