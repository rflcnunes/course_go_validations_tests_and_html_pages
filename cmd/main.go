package main

import (
	"fmt"

	"github.com/rflcnunes/course_validations_tests_and_html_pages/config"
	"github.com/rflcnunes/course_validations_tests_and_html_pages/internal/routes"
)

func main() {
	fmt.Println("Hello, Friend!")

	config.Setup()
	routes.Setup()
}
