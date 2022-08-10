package models

import "time"

type TaskResponse struct {
	ID          int       `json:"id"`
	Task_detail string    `json:"task_detail"`
	Assigne     string    `json:"assigne"`
	Deadline    time.Time `json:"deadline"`
}
