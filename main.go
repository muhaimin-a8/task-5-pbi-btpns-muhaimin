package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"log"
	"pbi-btpns-api/internal/app"
	"pbi-btpns-api/internal/controller"
	"pbi-btpns-api/internal/middleware"
	"pbi-btpns-api/internal/repository"
	"pbi-btpns-api/internal/router"
	"pbi-btpns-api/internal/service"
)

func main() {
	// load env
	err := app.LoadConfig("config")
	if err != nil {
		log.Fatalln(err)
	}

	//db
	db, err := app.NewDB()
	if err != nil {
		log.Fatalln(err)
	}

	// check db connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Can't connect to database => %v\n", err)
	}

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
	engine.Use(gin.CustomRecovery(app.ErrorHandler))

	router.InitRouter(engine, controllers, middlewares)
	err = engine.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
