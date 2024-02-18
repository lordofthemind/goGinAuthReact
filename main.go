package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type GoUser struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

var (
	users []GoUser
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	router.Use(cors.New(config))

	router.GET("/", goGinHello)
	router.POST("/login", goGinLogin)
	router.POST("/register", goGinRegister)

	dummyUsers := []GoUser{
		{Username: "user1", Email: "user1@example.com", Password: "password1"},
		{Username: "user2", Email: "user2@example.com", Password: "password2"},
		// Add more dummy users as needed
	}

	// Append dummy users to the existing users slice
	users = append(users, dummyUsers...)

	err := router.Run("localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
}

func goGinHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func goGinLogin(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, user := range users {
		if user.Username == loginData.Username && user.Password == loginData.Password {
			c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
}

func goGinRegister(c *gin.Context) {
	var registerData GoUser

	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, user := range users {
		if user.Username == registerData.Username {
			c.JSON(http.StatusConflict, gin.H{"message": "Username already taken"})
			return
		}
	}

	users = append(users, registerData)

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
