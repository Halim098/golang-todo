package Router

import (
	"todo/Api"
	"todo/Controller"
	"todo/Database"
	"todo/Middleware"
	"todo/Repository"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserApi Api.UserApi
}

type TodoHandler struct {
	TodoApi Api.TodoApi
}

func Router(router *gin.Engine) *gin.Engine {
	Db := Database.Connect()
	UserRepository := Repository.NewUserRepository(Db)
	UserController := Controller.NewUserController(UserRepository)
	UserApi := Api.NewUserApi(UserController)

	apiHandler := UserHandler{
		UserApi: UserApi,
	}

	router.POST("/register", apiHandler.UserApi.RegisterUser)
	router.POST("/login", apiHandler.UserApi.LoginUser)

	return router
}

func RouterTodo(router *gin.Engine) *gin.Engine {
	Db := Database.Connect()
	TodoRepository := Repository.NewTodoRepository(Db)
	TodoController := Controller.NewTodoController(TodoRepository)
	TodoApi := Api.NewTodoApi(TodoController)

	apiHandler := TodoHandler{
		TodoApi: TodoApi,
	}

	router.Use(Middleware.Auth())
	router.POST("/todo/add", apiHandler.TodoApi.SaveTodo)
	router.GET("/todo/:id", apiHandler.TodoApi.GetTodoById)
	router.GET("/todo", apiHandler.TodoApi.GetAllTodosUser)
	router.PUT("/todo/:id", apiHandler.TodoApi.UpdateTodo)
	router.DELETE("/todo/:id", apiHandler.TodoApi.DeleteTodo)

	return router
}
