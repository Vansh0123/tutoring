package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Connector struct {
	accessor *Queries
}

func (c *Connector) Connect() {
	log.Print(os.Getenv("DB_URI"))
	conn, err := sql.Open("postgres", os.Getenv("DB_URI"))
	if err != nil {
		log.Fatal("Error connect to database")
	}
	c.accessor = New(conn)
}

func (con *Connector) RegisterStudent(c *gin.Context) {
	type reqBody struct {
		Name      string `json:"name" binding:"required"`
		Subject   string `json:"subject" binding:"required"`
		Class     string `json:"class" binding:"required"`
		Fees      int32  `json:"fees" binding:"required"`
		FeeStatus string `json:"fee_status" binding:"required"`
	}
	var rBody reqBody
	if err := c.Bind(&rBody); err != nil {
		c.JSON(400, gin.H{
			"error": "Bad Request : One of the field might be empty or wrong",
		})
		return
	}
	log.Print("Successfuly parsed")

	student, err := con.accessor.RegisterStudent(c.Request.Context(), RegisterStudentParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      rBody.Name,
		Subject:   rBody.Subject,
		Class:     rBody.Class,
		Fees:      rBody.Fees,
		FeeStatus: rBody.FeeStatus,
	})

	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "User succesfully created",
		"student": student,
	})

}

func (con *Connector) GetStudent(c *gin.Context) {
	name := c.Param("name")
	user, err := con.accessor.GetStudentByName(c.Request.Context(), name)
	if err != nil {
		c.Status(400)
	}
	c.JSON(200, gin.H{
		"user": user,
	})
}
