package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rflcnunes/course_validations_tests_and_html_pages/internal/controllers"
)

var PATH = "/students"

func Setup() {
	r := gin.Default()
	r.GET(PATH, controllers.DisplayAllStudents)
	r.GET("/greeting/:name", controllers.Greeting)
	r.POST(PATH, controllers.CreateNewStudent)
	r.GET(PATH+"/:id", controllers.GetStudentByID)
	r.DELETE(PATH+"/:id", controllers.DeleteStudent)
	r.PUT(PATH+"/:id", controllers.EditStudent)
	r.GET(PATH+"/cpf/:cpf", controllers.GetStudentByCPF)
	r.Run()
}
