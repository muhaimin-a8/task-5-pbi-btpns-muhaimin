package router

import (
	"github.com/gin-gonic/gin"
	"pbi-btpns-api/controller"
	"pbi-btpns-api/middleware"
)

func InitRouter(engine *gin.Engine, controllers controller.Controllers, middlewares middleware.Middlewares) {
	router := engine.Group("/api/v1")
	//router.Use(middleware.ApiKeyAuth)

	// USERS
	users := router.Group("/users")
	users.POST("/", controllers.NewUserController().RegisterUser)

	users.Use(middlewares.NewJwtAuth().Init)
	users.PUT("/:userId", controllers.NewUserController().UpdateUser)
	users.DELETE("/:userId", controllers.NewUserController().DeleteUser)

	// AUTHENTICATIONS
	auth := router.Group("/users/auth")
	auth.POST("/login", controllers.NewAuthController().Login)
	auth.DELETE("/logout", controllers.NewAuthController().Logout)
	auth.PUT("/token", controllers.NewAuthController().UpdateAccessToken)

	//PHOTOS
	//photo := router.Group("/users/:userId/photos")

	// UPLOAD PHOTO
	//photo.POST("/upload", controllers.NewPhotoControlle)
}
