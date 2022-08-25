package handlers

import (
	"context"
	"net/http"

	pb "github.com/Services/imanuz_service/api_service/genproto/data_service"
	l "github.com/Services/imanuz_service/api_service/pkg/logger"

	_ "github.com/Services/imanuz_service/api_service/api/models"

	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreatePost godoc
// @Summary Create new post
// @Description This API for creating a new post
// @Tags Post
// @Accept json
// @Param body body models.CreatedPost true "body"
// @Produce json
// @Success 200
// @Router /posts [post]
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		body        pb.Link
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.DataService().Create(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create post", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}
