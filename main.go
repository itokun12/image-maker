package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/Sirupsen/logrus"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Error("Warning: .env file not loading. " + err.Error())
	}
	r := gin.Default()
	r.Run(":" + listenPort())
}

func listenPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8000"
	}
	return port
}
