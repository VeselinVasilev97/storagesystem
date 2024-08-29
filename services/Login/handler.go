package login

import (
	"net/http"
	"os"
	"storage/configuration"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Secret key used to sign the JWT
var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// User struct to represent the expected request body
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims struct to include within the JWT token
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// LoginHandler handles the login requests
func LoginHandler(conf *configuration.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User

		// Bind JSON to the user struct and check for errors
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		// Dummy authentication check
		if user.Username != "admin" || user.Password != "password" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Create and sign a new token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"exp":      time.Now().Add(time.Hour * 1).Unix(),
		})

		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
			return
		}

		// Send the token to the client
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"token":   tokenString,
		})
	}
}
