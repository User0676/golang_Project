
CREATE TABLE IF NOT exists Instructors (
    instructor_id SERIAL PRIMARY KEY,
    instructor_name VARCHAR(100),
    specialization VARCHAR(100),
    gender GenderEnum
);