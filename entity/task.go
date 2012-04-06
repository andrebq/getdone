package entity

type Task struct {
	Id          int64
	Title       string
	Description string
	Done        bool
	Project     *Project
}
