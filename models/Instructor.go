package models

type Instructor struct {
	InstructorId   int64  `json:"instructorId"`
	InstructorName string `json:"instructorName"`
	Specialization string `json:"specialization"`
	Gender         string `json:"gender"`
}
