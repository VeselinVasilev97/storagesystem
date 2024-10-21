package login

import (
	"net/http"
	"os"
	"storage/configuration"
	"strings"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

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
func LoginHandler(conf *configuration.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		jwtKey := os.Getenv("JWT_SECRET_KEY")
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		user.Username = strings.TrimSpace(strings.ToLower(user.Username))

		var dbUser User // Use the User model from the users package
		if err := conf.Db.Where("lower(username) = ?", user.Username).First(&dbUser).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"exp":      time.Now().Add(time.Hour * 1).Unix(),
		})

		tokenString, err := token.SignedString([]byte(jwtKey))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"token":   tokenString,
		})
	}
}
