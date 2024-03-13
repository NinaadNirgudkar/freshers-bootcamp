package controller

import (
	"bootcamp/daythree/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(200, user)
	}
}

func CreateUser(c *gin.Context) {
	var user Models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}
	fmt.Println(&user, err)
	err = Models.CreateUser(&user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusCreated, user)
	}
}
func GetUserByID(c *gin.Context) {
	id, ok := c.Params.Get("id")
	var user Models.User
	err := Models.GetUserByID(&user, id)
	if err != nil || !ok {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
	}
	err = Models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
