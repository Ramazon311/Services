package handlers

import (
	"errors"
	"net/http"

	"github.com/Services/imanuz_service/api_service/config"
	"github.com/Services/imanuz_service/api_service/pkg/logger"
	"github.com/Services/imanuz_service/api_service/services"

	"github.com/Services/imanuz_service/api_service/api/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
}

type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
}

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
	}
}

func GetClaims(h *handlerV1, c *gin.Context) jwt.MapClaims {
	var (
		ErrUnauthorized = errors.New("unauthorized")
		claims          jwt.MapClaims
		err             error
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Unauthorized request",
			},
		})
		h.log.Error("Unauthorized request: ", logger.Error(ErrUnauthorized))
		return nil
	}
	return claims
}
