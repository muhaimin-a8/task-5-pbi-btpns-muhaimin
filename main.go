package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
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
	db, err := app.NewDB(app.Production)
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
	apiKeyService := service.NewApiKeyService(dao, idGenerator)

	// controller
	validate := validator.New()
	controllers := controller.NewController(validate, userService, authService, photoService, apiKeyService)

	// middlewares
	middlewares := middleware.NewMiddlewares(tokenManager, apiKeyService)

	// server
	port := viper.Get("server.port").(int)
	stage := os.Getenv("STAGE")

	if stage == "production" {
		gin.SetMode(gin.ReleaseMode)

		// write log files
		err = os.MkdirAll("./logs", os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
		file, _ := os.Create("./logs/app.log")
		errLog, _ := os.Create("./logs/error.log")
		gin.DefaultWriter = io.MultiWriter(file)
		gin.DefaultErrorWriter = io.MultiWriter(errLog)
	}

	engine := gin.New()
	engine.Use(gin.CustomRecovery(app.ErrorHandler))

	router.InitRouter(engine, controllers, middlewares)

	err = engine.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln(err)
	}
}
