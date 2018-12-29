package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/morio-pg/stubtool/src/interfaces/api/server/handler"
	"github.com/morio-pg/stubtool/src/interfaces/api/server/middleware"
)

// Router interface
type Router interface {
	Setup() *gin.Engine
}

type router struct {
	authenticationMiddleware middleware.AuthenticationMiddleware
	accountHandler           handler.AccountHandler
	adminHandler             handler.AdminHandler
	stubsHandler             handler.StubsHandler
}

// NewRouter return Router
func NewRouter(
	authenticationMiddleware middleware.AuthenticationMiddleware,
	accountHandler handler.AccountHandler,
	adminHandler handler.AdminHandler,
	stubsHandler handler.StubsHandler,
) Router {
	return &router{authenticationMiddleware, accountHandler, adminHandler, stubsHandler}
}

func (r *router) Setup() *gin.Engine {
	router := gin.Default()

	api := router.Group("api")
	{
		account := api.Group("account", r.authenticationMiddleware.Check)
		{
			account.DELETE("", r.accountHandler.Delete)
		}

		admin := api.Group("admin", r.authenticationMiddleware.Check)
		{
			admin.GET("", r.adminHandler.GetAll)
			admin.POST("", r.adminHandler.Post)
			admin.GET("/:stubId", r.adminHandler.Get)
			admin.PUT("/:stubId", r.adminHandler.Put)
			admin.DELETE(":stubId", r.adminHandler.Delete)
		}

		stubs := api.Group("stubs", cors.Default())
		{
			stubs.GET("/:stubId/:filename", r.stubsHandler.Get)
		}
	}

	return router
}
