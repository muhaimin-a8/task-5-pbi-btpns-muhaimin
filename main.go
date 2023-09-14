package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"log"
	"pbi-btpns-api/app"
	"pbi-btpns-api/controller"
	"pbi-btpns-api/database"
	"pbi-btpns-api/exception"
	"pbi-btpns-api/middleware"
	"pbi-btpns-api/model"
	"pbi-btpns-api/repository"
	"pbi-btpns-api/router"
	"pbi-btpns-api/service"
)

func main() {
	// load env
	app.LoadConfig()

	//db
	db, _ := database.NewDB()

	//dao
	dao := repository.NewDAO(db)

	//jwt config
	key := viper.Get("jwt.key").(string)
	exp := viper.Get("jwt.exp").(int)
	refreshKey := viper.Get("jwt.refreshKey").(string)
	refreshExp := viper.Get("jwt.refreshExp").(int)

	//service
	idGenerator := service.NewIdGenerator()
	hasher := service.NewPasswordHasher()
	tokenManager := service.NewJwtTokenManager(key, exp, refreshKey, refreshExp)
	userService := service.NewUserService(dao, idGenerator, hasher)
	authService := service.NewAuthService(dao, hasher, tokenManager)
	photoService := service.NewPhotoService(dao, idGenerator)

	// controller
	validate := validator.New()
	controllers := controller.NewController(validate, userService, authService, photoService)

	// middlewares
	middlewares := middleware.NewMiddlewares(tokenManager)

	//routes

	engine := gin.New()
	engine.Use(gin.CustomRecovery(errorHandler))

	router.InitRouter(engine, controllers, middlewares)
	err := engine.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}

func errorHandler(c *gin.Context, err any) {
	if v, ok := err.(exception.ValidationError); ok {
		c.JSON(400, model.WebResponse{
			Status:  model.Fail,
			Code:    400,
			Message: v.Msg,
			Data:    nil,
		})
		return
	}

	if v, ok := err.(exception.JsonParseError); ok {
		c.JSON(400, model.WebResponse{
			Status:  model.Fail,
			Code:    400,
			Message: v.Msg,
			Data:    nil,
		})
		return
	}

	if v, ok := err.(exception.InvariantError); ok {
		c.JSON(400, model.WebResponse{
			Status:  model.Fail,
			Code:    400,
			Message: v.Msg,
			Data:    nil,
		})
		return
	}

	if v, ok := err.(exception.AuthenticationError); ok {
		c.JSON(401, model.WebResponse{
			Status:  model.Fail,
			Code:    401,
			Message: v.Msg,
			Data:    nil,
		})
		return
	}

	if v, ok := err.(exception.AuthorizationError); ok {
		c.JSON(401, model.WebResponse{
			Status:  model.Fail,
			Code:    401,
			Message: v.Msg,
			Data:    nil,
		})
		return
	}
	if v, ok := err.(exception.NotFoundError); ok {
		c.JSON(404, model.WebResponse{
			Status:  model.Fail,
			Code:    404,
			Message: v.Msg,
			Data:    nil,
		})
		return
	}

	// INTERNAL SERVER ERROR
	c.JSON(500, model.WebResponse{
		Status:  model.Error,
		Code:    500,
		Message: "Internal Server Error",
		Data:    nil,
	})
}
