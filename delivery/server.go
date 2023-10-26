package delivery

import (
	"payment-application/config"
	"payment-application/delivery/controller"
	"payment-application/delivery/middleware"
	"payment-application/manager"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type appServer struct {
	usecaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
	log            *logrus.Logger
}

func (a *appServer) initController() {
	a.engine.Use(middleware.LogRequestMiddleware(a.log))
	controller.NewProfileController(a.engine, a.usecaseManager.ProfileUsecase())
	controller.NewMerchantController(a.engine, a.usecaseManager.MerchantUsecase())
	controller.NewUserController(a.engine, a.usecaseManager.UserUsecase())
	controller.NewAuthController(a.engine, a.usecaseManager.AuthUsecase())

}

func (a *appServer) Run() {
	a.initController()

	err := a.engine.Run(a.host)
	if err != nil {
		panic(err.Error())
	}
}

func Server() *appServer {
	engine := gin.Default()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln("Error Config : ()", err.Error())
	}

	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		log.Fatalln("Error Conection : ", err.Error())
	}

	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	host := fmt.Sprintf("%s:%s", cfg.APIHost, cfg.APIPort)

	return &appServer{
		engine:         engine,
		host:           host,
		usecaseManager: useCaseManager,
		log:            logrus.New(),
	}
}
