package Model

import (
	"time" // Import the "time" package
)

type Todo struct {
	ID          int        `json:"id"`
	Id_user     int        `json:"id_user" sql:"id_user"`
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description" binding:"required"`
	Status      string     `json:"status" binding:"required"`
	Duedate     *time.Time `json:"duedate" binding:"required"`
}

type InputTodo struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status" binding:"required"`
	DueDate     string `json:"duedate" binding:"required"`
}
