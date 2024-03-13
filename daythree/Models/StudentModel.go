package Models

import (
	"bootcamp/daythree/Config"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Student struct {
	gorm.Model
	name  string
	phone string
	email string
	dob   time.Time
}

type Subject struct {
	gorm.Model
	title string
}

type TestResults struct {
	ID        string  `json:"id" gorm:"primary_key;auto_increment"`
	SubjectID uint    `json:"-"`
	Subject   Subject `json:"subject_id" gorm:"foreignKey:SubjectID;references:ID"`
	StudentID uint    `json:"-"`
	Student   Student `json:"student_id" gorm:"foreignKey:StudentID;references:ID"`
}

func (t *Student) TableName() string {
	return "student"
}

func (t *Subject) TableName() string {
	return "subject"
}

func (t *TestResults) TableName() string {
	return "test_result"
}

func GetAllStudents(Student *[]Student) (err error) {
	if err = Config.DB.Find(Student).Error; err != nil {
		return err
	}
	return nil
}

// CreateStudent ... Insert New data
func CreateStudent(Student *Student) (err error) {
	fmt.Println("Students doc Students", Student)
	return Config.DB.Create(Student).Error
}

// GetStudentByID ... Fetch only one Student by Id
func GetStudentByID(Student *Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(Student).Error; err != nil {
		return err
	}
	return nil
}
func GetAllSubjects(Subject *[]Subject) (err error) {
	if err = Config.DB.Find(Subject).Error; err != nil {
		return err
	}
	return nil
}

// CreateSubject ... Insert New data
func CreateSubject(Subject *Subject) (err error) {
	fmt.Println("Subjects doc Subjects", Subject)
	return Config.DB.Create(Subject).Error
}

// GetSubjectByID ... Fetch only one Subject by Id
func GetSubjectByID(Subject *Subject, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(Subject).Error; err != nil {
		return err
	}
	return nil
}

func GetAllTestResults(TestResults *[]TestResults) (err error) {
	if err = Config.DB.Find(TestResults).Error; err != nil {
		return err
	}
	return nil
}

// CreateTestResults ... Insert New data
func CreateTestResults(TestResults *TestResults) (err error) {
	fmt.Println("TestResultss doc TestResultss", TestResults)
	return Config.DB.Create(TestResults).Error
}

// GetTestResultsByID ... Fetch only one TestResults by Id
func GetTestResultsByID(TestResults *TestResults, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(TestResults).Error; err != nil {
		return err
	}
	return nil
}
