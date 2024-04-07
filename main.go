package main

import (
	"log"
	"tutoring/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var cnx database.Connector

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
	router.POST("/student", cnx.RegisterStudent)
	router.GET("/student/:name", cnx.GetStudent)
	router.Run()
}
