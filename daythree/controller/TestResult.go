package controller

import (
	"bootcamp/daythree/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStudents(c *gin.Context) {
	var Student []Models.Student
	err := Models.GetAllStudents(&Student)
	if err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(200, Student)
	}
}

func CreateStudent(c *gin.Context) {
	var Student Models.Student
	err := c.ShouldBindJSON(&Student)
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}
	fmt.Println(&Student, err)
	err = Models.CreateStudent(&Student)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusCreated, Student)
	}
}
func GetStudentByID(c *gin.Context) {
	id, ok := c.Params.Get("id")
	var Student Models.Student
	err := Models.GetStudentByID(&Student, id)
	if err != nil || !ok {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Student)
	}
}

func GetSubjects(c *gin.Context) {
	var Subject []Models.Subject
	err := Models.GetAllSubjects(&Subject)
	if err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(200, Subject)
	}
}

func CreateSubject(c *gin.Context) {
	var Subject Models.Subject
	err := c.ShouldBindJSON(&Subject)
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}
	fmt.Println(&Subject, err)
	err = Models.CreateSubject(&Subject)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusCreated, Subject)
	}
}
func GetSubjectByID(c *gin.Context) {
	id, ok := c.Params.Get("id")
	var Subject Models.Subject
	err := Models.GetSubjectByID(&Subject, id)
	if err != nil || !ok {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Subject)
	}
}

func GetTestResults(c *gin.Context) {
	var TestResult []Models.TestResults
	err := Models.GetAllTestResults(&TestResult)
	if err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(200, TestResult)
	}
}

func CreateTestResult(c *gin.Context) {
	var TestResult Models.TestResults
	err := c.ShouldBindJSON(&TestResult)
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}
	fmt.Println(&TestResult, err)
	err = Models.CreateTestResults(&TestResult)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusCreated, TestResult)
	}
}
func GetTestResultByID(c *gin.Context) {
	id, ok := c.Params.Get("id")
	var TestResult Models.TestResults
	err := Models.GetTestResultsByID(&TestResult, id)
	if err != nil || !ok {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, TestResult)
	}
}
