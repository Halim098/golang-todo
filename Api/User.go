package Api

import (
	"net/http"
	"time"
	"todo/Controller"
	"todo/Model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserApi interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type userApi struct {
	UserController Controller.UserController
}

func NewUserApi(UserController Controller.UserController) *userApi {
	return &userApi{UserController}
}

func (u *userApi) RegisterUser(c *gin.Context) {
	start := time.Now()
	data := Model.User{}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error":    err.Error(),
			"duration": time.Since(start).String(),
		}).Warn("Failed to bind JSON during user registration")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = u.UserController.RegisterUser(&data)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"email":    data.Email,
			"error":    err.Error(),
			"duration": time.Since(start).String(),
		}).Error("Failed to register user")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"email":    data.Email,
		"duration": time.Since(start).String(),
	}).Info("User registered successfully")

	c.JSON(http.StatusCreated, gin.H{"message": "user registered", "data": data})
}

func (u *userApi) LoginUser(c *gin.Context) {
	start := time.Now()
	data := Model.LoginInput{}

	err := c.ShouldBindJSON(&data)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error":    err.Error(),
			"duration": time.Since(start).String(),
		}).Warn("Failed to bind JSON during user login")

		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if data.Email == "" || data.Password == "" {
		logrus.WithFields(logrus.Fields{
			"duration": time.Since(start).String(),
		}).Warn("Email and password are required")

		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	token, err := u.UserController.LoginUser(&data)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"email":    data.Email,
			"error":    err.Error(),
			"duration": time.Since(start).String(),
		}).Error("Failed to login user")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:  "todo_cookie",
		Value: token,
	})

	logrus.WithFields(logrus.Fields{
		"email":    data.Email,
		"duration": time.Since(start).String(),
	}).Info("User logged in successfully")

	c.JSON(http.StatusOK, gin.H{"message": "login success"})
}
