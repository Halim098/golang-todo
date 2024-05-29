package Api

import (
	"net/http"
	"strconv"
	"todo/Controller"
	"todo/Model"

	"github.com/gin-gonic/gin"
)

type TodoApi interface {
	SaveTodo(c *gin.Context)
	GetTodoById(c *gin.Context)
	GetAllTodosUser(c *gin.Context)
	UpdateTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
}

type todoApi struct {
	TodoController Controller.TodoController
}

func NewTodoApi(TodoController Controller.TodoController) *todoApi {
	return &todoApi{TodoController}
}

func (t *todoApi) SaveTodo(c *gin.Context) {
	id := c.MustGet("id").(int)
	data := Model.InputTodo{}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err = t.TodoController.SaveTodo(&data, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "todo created", "data": data})
}

func (t *todoApi) GetTodoById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	todo, err := t.TodoController.GetTodoById(idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}

	err = t.TodoController.CheckAuthorization(todo, c.MustGet("id").(int))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func (t *todoApi) GetAllTodosUser(c *gin.Context) {
	id := c.MustGet("id").(int)

	todos, err := t.TodoController.GetAllTodosUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func (t *todoApi) UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	todo, err := t.TodoController.GetTodoById(idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}

	err = t.TodoController.CheckAuthorization(todo, c.MustGet("id").(int))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	data := Model.InputTodo{}

	err = c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err = t.TodoController.UpdateTodo(&data, idInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo updated", "data": data})
}

func (t *todoApi) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	todo, err := t.TodoController.GetTodoById(idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}

	err = t.TodoController.CheckAuthorization(todo, c.MustGet("id").(int))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	err = t.TodoController.DeleteTodo(idInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo deleted"})
}
