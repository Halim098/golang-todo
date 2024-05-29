package Repository

import (
	"errors"
	"time"
	"todo/Model"

	"gorm.io/gorm"
)

type TodoRepository interface {
	SaveTodo(data *Model.InputTodo, id int) error
	GetTodoById(id int) (Model.Todo, error)
	GetAllTodosUser(id int) ([]Model.Todo, error)
	UpdateTodo(data *Model.InputTodo, id int) error
	DeleteTodo(id int) error
	CheckAuthorization(todo Model.Todo, id int) error
}

type todoRepository struct {
	Db *gorm.DB
}

func NewTodoRepository(Db *gorm.DB) *todoRepository {
	return &todoRepository{Db}
}

func (t *todoRepository) SaveTodo(data *Model.InputTodo, id int) error {
	date, err := time.Parse("2006-01-02", data.DueDate)
	if err != nil {
		return err
	}

	result := t.Db.Exec("INSERT INTO todos (title, description, status, id_user, duedate) VALUES (?, ?, ?, ?, ?)", data.Title, data.Description, data.Status, id, date)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t *todoRepository) GetTodoById(id int) (Model.Todo, error) {
	todo := Model.Todo{}
	result := t.Db.Raw("SELECT * FROM todos WHERE id = ?", id).Scan(&todo)
	if result.Error != nil {
		return todo, result.Error
	}
	if result.RowsAffected == 0 {
		return todo, errors.New("todo not found")
	}
	return todo, nil
}

func (t *todoRepository) GetAllTodosUser(id int) ([]Model.Todo, error) {
	var todos []Model.Todo
	result := t.Db.Raw("SELECT * FROM todos where id_user = ?", id).Scan(&todos)
	if result.Error != nil {
		return todos, result.Error
	}
	if result.RowsAffected == 0 {
		return todos, errors.New("empty todo")
	}
	return todos, nil
}

func (t *todoRepository) UpdateTodo(data *Model.InputTodo, id int) error {
	date, err := time.Parse("2006-01-02", data.DueDate)
	if err != nil {
		return err
	}

	result := t.Db.Exec("UPDATE todos SET title = ?, description = ?, status = ?,duedate = ? WHERE id = ?", data.Title, data.Description, data.Status, date, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t *todoRepository) DeleteTodo(id int) error {
	result := t.Db.Exec("DELETE FROM todos WHERE id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t *todoRepository) CheckAuthorization(todo Model.Todo, id int) error {
	if todo.Id_user != id {
		return errors.New("unauthorized")
	}
	return nil
}
