package handlers

import (
	"context"
	"net/http"
	_ "github/Services/newpro/Api-TU/api/models"
	pb "github/Services/newpro/Api-TU/genproto/task_service"
	l "github/Services/newpro/Api-TU/pkg/logger"

	//user "newpro/Api-TU/genproto/user_service"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateTask godoc
// @Summary Create new task
// @Description This API for creating a new task
// @Tags Task
// @Accept json
// @Param body body models.CreateTask true "body"
// @Produce json
// @Success 201 {object} models.Task
// @Router /tasks [post]
func (h *handlerV1) CreateTask(c *gin.Context) {
	var (
		body        pb.Task
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

	response, err := h.serviceManager.TaskService().Create(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create task", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// GetTask godoc
// @Summary Get task
// @Description  Get task
// @Tags Task
// @Accept json
// @Param id path string true "ID"
// @Produce json
// @Success 200 {object} models.Task
// @Router /task/{id} [get]
func (h *handlerV1) GetTask(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaskService().Get(
		ctx, &pb.IdReq{
			Id: guid,
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get task", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)

}

// @Summary Update task
// @Schemes
// @Description  Update task
// @Tags Task
// @Accept json
// @Param body body models.UpdateTasks true "body"
// @Produce json
// @Success 200 {object} models.Task
// @Router /task/{id} [put]
func (h *handlerV1) UpdateTask(c *gin.Context) {
	var (
		body        pb.Task
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
	//body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaskService().Update(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteTask godoc
// @Summary Delete task
// @Schemes
// @Description  Delete task
// @Tags Task
// @Accept json
// @Param id path string true "ID"
// @Produce json
// @Success 200
// @Router /task/{id} [delete]
func (h *handlerV1) DeleteTask(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaskService().Delete(
		ctx, &pb.IdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete task", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
