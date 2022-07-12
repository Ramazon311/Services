package api

import (
	"github.com/Services/imanuz_service/api_service/config"
	"github.com/Services/imanuz_service/api_service/pkg/logger"
	"github.com/Services/imanuz_service/api_service/services"

	v1 "github.com/Services/imanuz_service/api_service/api/handler"

	"github.com/Services/imanuz_service/api_service/api/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 			Post api
// @version			1.0
// @description		This is Data and Post service Api
// @termsOfService	http://swagger.io/terms/

// @securityDefinitions.apikey BearerAuth
// @in  header
// @name Authorization

// @contact.name	Api Support
// @contact.url		http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url		http://www.apache.org/license/LICENSE-2.0.html

// @host	localhost:8080
// @BasePath /v1

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	docs.SwaggerInfo.BasePath = "/v1"

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	api.POST("/posts", handlerV1.CreatePost)
	api.GET("/post/:id", handlerV1.GetPost)
	api.PUT("/post/:id", handlerV1.UpdatePost)
	api.DELETE("/post/:id", handlerV1.DeletePost)
	api.GET("/posts/:id", handlerV1.GetList)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
