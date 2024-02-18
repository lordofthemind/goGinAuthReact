package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	validUsername = "goUser"
	validPassword = "goPassword"
)

func main() {
	router := gin.Default()

	// Enable CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	router.Use(cors.New(config))

	router.GET("/", goGinHello)
	router.POST("/login", goGinLogin)

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

	if loginData.Username == validUsername && loginData.Password == validPassword {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
	}
}
