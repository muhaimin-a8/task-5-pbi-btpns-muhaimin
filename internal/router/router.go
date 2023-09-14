package router

import (
	"github.com/gin-gonic/gin"
	"pbi-btpns-api/internal/controller"
	"pbi-btpns-api/internal/middleware"
)

func InitRouter(engine *gin.Engine, controllers controller.Controllers, middlewares middleware.Middlewares) {
	router := engine.Group("/api/v1")
	router.Use(middlewares.NewApiKeyAuth().Init)

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

	// PHOTOS
	photos := users.Group("/:userId/photos")
	photos.POST("/", controllers.NewPhotoController().AddPhoto)
	photos.GET("/:photoId", controllers.NewPhotoController().GetPhoto)
	photos.PUT("/:photoId", controllers.NewPhotoController().UpdatePhoto)
	photos.DELETE("/:photoId", controllers.NewPhotoController().DeletePhoto)

	// UPLOAD PHOTO
	upload := engine.Group("/uploads")
	upload.POST("/photos", controllers.NewUploadController().UploadPhoto)

	// GET STATIC CONTENT
	engine.Static("/static/photos/", "./public/uploads/photos/")

	// REGISTER NEW API KEY
	apiKey := engine.Group("/api-keys")
	apiKey.POST("/register", controllers.NewApiKeyController().RegisterNewApiKey)
}
