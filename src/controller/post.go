package controller

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"net/http"
	services "example_app/service"
	httpEntity "example_app/entity/http"
)

type PostController struct {
	PostService services.PostServiceInterface
}

func (service *PostController) GetPostById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	result := service.PostService.GetPostById(id)
	if result == nil {
		context.JSON(http.StatusOK, gin.H{})
		return
	}
	context.JSON(http.StatusOK, result)
}

func (service *PostController) GetPosts(context *gin.Context) {
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
	result := service.PostService.GetAllPosts(queryparam.Limit, queryparam.Offset)
	context.JSON(http.StatusOK, result)
}

func (service *PostController) StorePost(context *gin.Context) {
	payload := httpEntity.PostRequest{}
	if err := context.ShouldBind(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	success := service.PostService.StorePost(payload)
	if !success {
		context.JSON(http.StatusNoContent, gin.H{})
		return
	}

	context.JSON(http.StatusCreated, gin.H{})
	
}

func (service *PostController) UpdatePostById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	payload := httpEntity.PostRequest{}
	if err := context.ShouldBind(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	success := service.PostService.UpdatePostById(id,payload)

	if !success {
		context.JSON(http.StatusNoContent, gin.H{})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{})
}

func (service *PostController) DeletePost(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	
	result := service.PostService.DeletePost(id)

	if result.Id == 0 {
		context.JSON(http.StatusNoContent, gin.H{})
		return
	}

	context.JSON(http.StatusOK, result)
}
