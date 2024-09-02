package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rflcnunes/course_validations_tests_and_html_pages/config"
	"github.com/rflcnunes/course_validations_tests_and_html_pages/internal/models"
)

func DisplayAllStudents(c *gin.Context) {
	var students []models.Student
	config.DB.Find(&students)
	c.JSON(200, students)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API says:": "Hey " + name + ", how are you doing?",
	})
}

func CreateNewStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	config.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}
