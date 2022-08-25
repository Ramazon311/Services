package handlers

import (
	"context"
	"fmt"
	"net/http"
	"github/Services/newpro/Api-TU/api/models"
	_ "github/Services/newpro/Api-TU/api/models"
	pb "github/Services/newpro/Api-TU/genproto/user_service"
	l "github/Services/newpro/Api-TU/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateUser godoc
// @Summary Create new user
// @Description This API for creating a new user
// @Tags User
// @Accept json
// @Param body body models.CreateUser true "body"
// @Produce json
// @Success 201 {object} models.User
// @Router /users [post]
func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		body        pb.User
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

	response, err := h.serviceManager.UserService().Create(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetUser godoc
// @Summary GetUser
// @Schemes
// @Description  Get User
// @Security 	BearerAuth
// @Tags User
// @Accept json
// @Param id path string true "ID"
// @Produce json
// @Success 200 {object} models.User
// @Router /user/{id} [get]
func (h *handlerV1) Get(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().Get(
		ctx, &pb.IdReq{
			Id: guid,
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get User", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)

}

// Login godoc
// @Summary Login
// @Schemes
// @Description  Get My Profile
// @Security 	BearerAuth
// @Tags User
// @Accept json
// @Param body body models.Login true "body"
// @Produce json
// @Success 200 {object} models.User
// @Router /user [post]
func (h *handlerV1) Login(c *gin.Context) {
	var (
		body        pb.Login1
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

	response, err := h.serviceManager.UserService().Login(ctx, &pb.Login1{
		Password: string(body.Password),
		Email:    body.Email,
	})

	// Creating hash of a password
	compareErr := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(body.Password))
	

	if compareErr != nil {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Password Error",
			},
		})
		return
	}

	

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get User Profile", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)

}

// @Summary Update User
// @Schemes
// @Description  Update User
// @Security 	BearerAuth
// @Tags User
// @Accept json
// @Param body body models.UpdateUser true "body"
// @Produce json
// @Success 200 {object} models.User
// @Router /user/{id} [put]
func (h *handlerV1) UpdateUser(c *gin.Context) {
	var (
		body        pb.User
		jspbMarshal protojson.MarshalOptions
	)

	bul, err := h.serviceManager.UserService().CheckField(c, &pb.Check{Field: "username", Value: body.Username})
	fmt.Println(bul.Status)

	if bul.Status != false {
		c.JSON(int(http.StatusAccepted), "This username or email already is registered")
		return
	}

	bul, err = h.serviceManager.UserService().CheckField(c, &pb.Check{Field: "email", Value: body.Email})
	if bul.Status != false {
		c.JSON(int(http.StatusAccepted), "This email already or email is registered")

		return
	}

	jspbMarshal.UseProtoNames = true

	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().Update(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary DeleteUser
// @Schemes
// @Description  Delete User
// @Security 	BearerAuth
// @Tags User
// @Param id path string true "ID"
// @Accept json
// @Produce json
// @Success 200
// @Router /user/{id} [delete]
func (h *handlerV1) DeleteUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().Delete(
		ctx, &pb.IdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete User", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetMyProfile godoc
// @Summary GetMyProfile
// @Schemes
// @Description  Get My Profile
// @Security 	BearerAuth
// @Tags User
// @Accept json
// @Param id path string true "ID"
// @Produce json
// @Success 200 {object} models.User
// @Router /profile [get]
func (h *handlerV1) GetMyProfile(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	claims := GetClaims(h,c)

	guid := claims["sub"].(string)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().Get(
		ctx, &pb.IdReq{
			Id: guid,
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get User Profile", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)

}
