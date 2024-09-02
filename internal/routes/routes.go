package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rflcnunes/course_validations_tests_and_html_pages/internal/controllers"
)

func Setup() {
	r := gin.Default()
	r.GET("/students", controllers.DisplayAllStudents)
	r.GET("/greeting/:name", controllers.Greeting)
	r.POST("/students", controllers.CreateNewStudent)
	r.GET("/students/:id", controllers.GetStudentByID)
	r.Run()
}
