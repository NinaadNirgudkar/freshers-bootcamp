package Routes

import (
	"bootcamp/daythree/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(server *gin.Engine) *gin.Engine {
	grp1 := server.Group("user-api")
	{
		grp1.GET("user", controller.GetUsers)
		grp1.POST("user", controller.CreateUser)
		grp1.GET("user/:id", controller.GetUserByID)
		grp1.PUT("user/:id", controller.UpdateUser)
		grp1.DELETE("user/:id", controller.DeleteUser)
	}
	return server
}
func StudentTestRouter(server *gin.Engine) *gin.Engine {
	grp2 := server.Group("student")
	{
		grp2.GET("/", controller.GetStudents)
		grp2.POST("/", controller.CreateStudent)
		grp2.GET("/:id", controller.GetStudentByID)
	}
	grp3 := server.Group("subject")
	{
		grp3.GET("/", controller.GetSubjects)
		grp3.POST("/", controller.CreateSubject)
		grp3.GET("/:id", controller.GetStudentByID)
	}
	grp4 := server.Group("test-results")
	{
		grp4.GET("/", controller.GetTestResults)
		grp4.POST("/", controller.CreateTestResult)
		grp4.GET("/:id", controller.GetTestResults)
	}
	return server
}
