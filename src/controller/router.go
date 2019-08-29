package controller

import (
	"github.com/gin-gonic/gin"
	srv "example_app/service"
)

func LoadRouter(routers *gin.Engine) {
	user := &UserRouterLoader{}
	post := &PostRouterLoader{}
	user.UserRouter(routers)
	post.PostRouter(routers)
}

type UserRouterLoader struct{
}

type PostRouterLoader struct{
}

func (rLoader *UserRouterLoader) UserRouter(router *gin.Engine) {
	handler := &UserController{
		UserService: srv.UserServiceHandler(),
	}
	rLoader.routerDefinition(router, handler)
}

func (rLoader *PostRouterLoader) PostRouter(router *gin.Engine) {
	handler := &PostController{
		PostService: srv.PostServiceHandler(),
	}
	rLoader.routerDefinition(router,handler)
}

func (rLoader *UserRouterLoader) routerDefinition(router *gin.Engine,handler *UserController) {
	group := router.Group("v1/users")
	group.GET("", handler.GetUsers)
	group.GET(":id", handler.GetUserByID)
	group.PUT(":id", handler.UpdateUsersByID)
	group.POST("", handler.StoreUser)
	group.DELETE(":id", handler.DeleteUser)
}

func (rLoader *PostRouterLoader) routerDefinition(router *gin.Engine,handler *PostController) {
	group := router.Group("v1/posts")
	group.GET("", handler.GetPosts)
	group.GET(":id", handler.GetPostById)
	group.PUT(":id", handler.UpdatePostById)
	group.POST("", handler.StorePost)
	group.DELETE(":id", handler.DeletePost)
}