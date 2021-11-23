package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Query(context *gin.Context) {
	value := context.Query("name")
	age := context.Query("age")
	atoi, err := strconv.Atoi(age)
	if err != nil {
		context.String(http.StatusBadRequest, "age is not num")
	}
	context.String(http.StatusOK, "name is: %s age is: %d", value, atoi)
}

func QueryParam(context *gin.Context) {
	value := context.Param("name")
	age := context.Param("age")
	atoi, err := strconv.Atoi(age)
	if err != nil {
		context.String(http.StatusBadRequest, "age is not num")
	}
	context.String(http.StatusOK, "name is: %s age is: %d", value, atoi)
}

type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	UserName string `json:"user_name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func Json(context *gin.Context) {
	var login Login
	if err := context.ShouldBindJSON(&login); err != nil {
		context.String(http.StatusBadRequest, "param error")
		return
	}

	context.JSON(http.StatusOK, login)
}
