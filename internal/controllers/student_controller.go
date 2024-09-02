package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rflcnunes/course_validations_tests_and_html_pages/config"
	"github.com/rflcnunes/course_validations_tests_and_html_pages/internal/models"
)

func DisplayAllStudents(c *gin.Context) {
	var students []models.Student
	config.DB.Find(&students)
	c.JSON(200, students)
}
