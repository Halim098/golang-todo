package Controller

import (
	"todo/Model"
	"todo/Repository"
)

type TodoController interface {
	SaveTodo(data *Model.InputTodo, id int) error
	GetTodoById(id int) (Model.Todo, error)
	GetAllTodosUser(id int) ([]Model.Todo, error)
	UpdateTodo(data *Model.InputTodo, id int) error
	DeleteTodo(id int) error
	CheckAuthorization(todo Model.Todo, id int) error
}

type todoController struct {
	TodoRepository Repository.TodoRepository
}

func NewTodoController(TodoRepository Repository.TodoRepository) *todoController {
	return &todoController{TodoRepository}
}

func (t *todoController) SaveTodo(data *Model.InputTodo, id int) error {
	err := t.TodoRepository.SaveTodo(data, id)
	if err != nil {
		return err
	}
	return nil
}

func (t *todoController) GetTodoById(id int) (Model.Todo, error) {
	todo, err := t.TodoRepository.GetTodoById(id)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (t *todoController) GetAllTodosUser(id int) ([]Model.Todo, error) {
	todos, err := t.TodoRepository.GetAllTodosUser(id)
	if err != nil {
		return todos, err
	}
	return todos, nil
}

func (t *todoController) UpdateTodo(data *Model.InputTodo, id int) error {
	err := t.TodoRepository.UpdateTodo(data, id)
	if err != nil {
		return err
	}
	return nil
}

func (t *todoController) DeleteTodo(id int) error {
	err := t.TodoRepository.DeleteTodo(id)
	if err != nil {
		return err
	}
	return nil
}

func (t *todoController) CheckAuthorization(todo Model.Todo, id int) error {
	err := t.TodoRepository.CheckAuthorization(todo, id)
	if err != nil {
		return err
	}
	return nil
}
