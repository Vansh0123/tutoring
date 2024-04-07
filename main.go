package main

import (
	"log"
	"tutoring/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var cnx middleware.Connector

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	cnx.Connect()
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	router := gin.Default()
	router.POST("/students", cnx.RegisterStudent)
	router.GET("/students/:name", cnx.GetStudent)
	router.GET("/students", cnx.GetAllStudents)
	router.Run()
}
