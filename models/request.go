package models

import "time"

type TaskRequest struct {
	Task_detail string    `json:"task_detail" binding:"required"`
	Assigne     string    `json:"assigne" binding:"required"`
	Deadline    time.Time `json:"deadline" binding:"required"`
}
