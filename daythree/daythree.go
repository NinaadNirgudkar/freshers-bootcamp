package daythree

import (
	"bootcamp/daythree/Config"
	"bootcamp/daythree/Models"
	"bootcamp/daythree/Routes"
	"bootcamp/daythree/middlewares"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"os"
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

// server program for day3,4,5 programs

func DayThree() {
	var err error
	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	//dsn := "root@tcp(127.0.0.1:3306)/bootcamp?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := Config.DbURL(Config.BuildDBConfig())

	Config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Conn to db failed")
	}
	err = Config.DB.AutoMigrate(&Models.User{}, &Models.Student{}, &Models.Subject{}, &Models.TestResults{})
	if err != nil {
		panic("Migration failed")
	}
	Routes.UserRouter(server)
	Routes.StudentTestRouter(server)
	err = server.Run(":8080")
	if err != nil {
		panic(err.Error())
	}
}
