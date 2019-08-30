package controller

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
)

func LoadRouterTestMock() (*gin.Context, *gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	resp := httptest.NewRecorder()
	context, routers := gin.CreateTestContext(resp)
	
	routerLoader := &UserRouterLoader{}
	routerLoader.UserRouterTestMock(routers)
	routerPost := &PostRouterLoader{}
	routerPost.PostRouterTestMock(routers)
	
	return context, routers, resp
}

func (rLoader *UserRouterLoader) UserRouterTestMock(router *gin.Engine){
	handler := &UserController{
		UserService: &UserServiceMock{},
	}
	rLoader.routerDefinition(router,handler)
}

func (rLoader *PostRouterLoader) PostRouterTestMock(router *gin.Engine){
	handler := &PostController{
		PostService: &PostServiceMock{},
	}
	rLoader.routerDefinition(router,handler)
}