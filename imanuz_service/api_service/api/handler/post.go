package handlers

import (
	"context"
	"net/http"

	pb "github.com/Services/imanuz_service/api_service/genproto/post_service"
	l "github.com/Services/imanuz_service/api_service/pkg/logger"

	_ "github.com/Services/imanuz_service/api_service/api/models"

	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// GetPost godoc
// @Summary Get post
// @Description  Get post
// @Tags Post
// @Accept json
// @Param id path string true "ID"
// @Produce json
// @Success 200 {object} models.Post
// @Router /post/{id} [get]
func (h *handlerV1) GetPost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().Get(ctx, &pb.IdP{Id: guid})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get post", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)

}

// @Summary Update post
// @Schemes
// @Description  Update post
// @Tags Post
// @Accept json
// @Param body body models.UpdatePost true "body"
// @Produce json
// @Success 200 {object} models.Post
// @Router /post/{id} [put]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var (
		body        pb.UpdatePost
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

	response, err := h.serviceManager.PostService().Update(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update post", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeletePost godoc
// @Summary Delete post
// @Schemes
// @Description  Delete post
// @Tags Post
// @Accept json
// @Param id path string true "ID"
// @Produce json
// @Success 200
// @Router /post/{id} [delete]
func (h *handlerV1) DeletePost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().Delete(ctx, &pb.IdP{Id: guid})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete post", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetList godoc
// @Summary Get posts
// @Description  Get posts
// @Tags Post
// @Accept json
// @Param body body models.List true "body"
// @Produce json
// @Success 200 {object} models.ListRes
// @Router /posts/ [put]
func (h *handlerV1) GetList(c *gin.Context) {
	var (
		body        pb.ListReq
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

	response, err := h.serviceManager.PostService().List(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update post", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
