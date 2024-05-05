package models

type Instructor struct {
	InstructorId   int64  `json:"instructorId"`
	InstructorName string `json:"instructorName"`
	Specialization string `json:"specialization"`
	Gender         string `json:"gender"`
}

type UserPermission struct {
	UserID       int64 `json:"user_id"`
	PermissionID int64 `json:"permission_id"`
}

type Permission struct {
	ID          int64  `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
