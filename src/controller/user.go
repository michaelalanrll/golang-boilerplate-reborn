package controller

import (
	"strconv"
	"sync"
	"github.com/gin-gonic/gin"
	"net/http"
	services "example_app/service"
	httpEntity "example_app/entity/http"
)

type UserController struct {
	UserService services.UserServiceInterface
}

func (handler *UserController) TestFunction(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func (service *UserController) GetUserByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	result := service.UserService.GetUserByID(id, &sync.WaitGroup{})
	if result == nil {
		context.JSON(http.StatusOK, gin.H{})
		return
	}
	context.JSON(http.StatusOK, result)
}

type Limitofset struct{
	Limit int `form:"limit"`
	Offset int `form:"offset"`
}
func (service *UserController) GetUsers(context *gin.Context) {
	queryparam := Limitofset{}
	err := context.ShouldBindQuery(&queryparam)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	if queryparam.Limit == 0 {
		queryparam.Limit = 10
	}
	result := service.UserService.GetAllUser(queryparam.Limit, queryparam.Offset)
	context.JSON(http.StatusOK, result)
}

func (service *UserController) StoreUser(context *gin.Context) {
	payload := httpEntity.UserRequest{}
	if err := context.ShouldBind(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	success := service.UserService.StoreUser(payload)
	if !success {
		context.JSON(http.StatusNoContent, gin.H{})
		return
	}

	context.JSON(http.StatusCreated, gin.H{})
	
}

func (service *UserController) UpdateUsersByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	payload := httpEntity.UserRequest{}
	if err := context.ShouldBind(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	success := service.UserService.UpdateUserByID(id,payload)

	if !success {
		context.JSON(http.StatusNoContent, gin.H{})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{})
}

func (service *UserController) DeleteUser(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	
	result := service.UserService.DeleteUser(id)

	if result.ID == 0 {
		context.JSON(http.StatusNoContent, gin.H{})
		return
	}

	context.JSON(http.StatusOK, result)
}
