package Middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

var jwtKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		token, err := c.Request.Cookie("todo_cookie")
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"status":   http.StatusUnauthorized,
				"error":    err.Error(),
				"duration": time.Since(start).String(),
			}).Warn("No token cookie found")

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		claims := &jwt.MapClaims{
			"id":    0,
			"email": "",
		}

		tokenString := token.Value
		parsedtoken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error":    err.Error(),
				"status":   http.StatusUnauthorized,
				"token":    tokenString,
				"duration": time.Since(start).String(),
			}).Error("Failed to parse JWT token")

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		if !parsedtoken.Valid {
			logrus.WithFields(logrus.Fields{
				"status":   http.StatusUnauthorized,
				"token":    tokenString,
				"duration": time.Since(start).String(),
			}).Warn("Invalid JWT token")

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		id := int((*claims)["id"].(float64))

		c.Set("id", id)
		c.Next()

		logrus.WithFields(logrus.Fields{
			"method":   c.Request.Method,
			"endpoint": c.Request.RequestURI,
			"user_id":  id,
			"status":   c.Writer.Status(),
			"duration": time.Since(start).String(),
		}).Info("Request processed")
	}
}
