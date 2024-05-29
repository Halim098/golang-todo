package main

import (
	"fmt"
	Logging "todo/Log"
	"todo/Router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	loadEnv()
	Logging.Init()

	gin := gin.Default()
	router := Router.Router(gin)
	routertodo := Router.RouterTodo(gin)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error running router")
	}

	err = routertodo.Run(":8080")
	if err != nil {
		log.Fatal("Error running router")
	}

	fmt.Println("Server running on port 8080")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
