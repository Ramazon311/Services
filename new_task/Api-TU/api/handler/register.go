package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github/Services/newpro/Api-TU/api/models"
	_ "github/Services/newpro/Api-TU/api/models"
	"github/Services/newpro/Api-TU/api/token"
	"github/Services/newpro/Api-TU/etc"
	pbml "github/Services/newpro/Api-TU/genproto/email_service"
	pb "github/Services/newpro/Api-TU/genproto/user_service"
	l "github/Services/newpro/Api-TU/pkg/logger"
	"github/Services/newpro/Api-TU/storage/redis"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	Id            string `json:"id"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Username      string `json:"username"`
	Profile_photo string `json:"profile_photo"`
	Bio           string `json:"bio"`
	Email         string `json:"Email"`
	Gender        string `json:"gender"`
	Address       string `json:"address"`
	Phone_number  string `json:"phone_number"`
	Created_at    string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
	Deleted_at    string `json:"deleted_at"`
}

// Register ...
// @Summary Register
// @Description Register - API for registering users
// @Tags register
// @Accept  json
// @Produce  json
// @Param register body models.User true "register"
// @Success 200 {object} models.RegisterResponseModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /register [post]
func (h *handlerV1) Register(c *gin.Context) {
	var (
		body models.User
		// code string
	)

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	body.Email = strings.TrimSpace(body.Email)
	body.Email = strings.ToLower(body.Email)

	err = body.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Message: err.Error(),
			},
		})
		h.log.Error("Error while validating", l.Error(err))
		return
	}

	// Checking uniqueness of username
	checkUsername, err := h.serviceManager.UserService().CheckField(
		context.Background(), &pb.Check{
			Field: "username",
			Value: body.Username,
		},
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Message: err.Error(),
			},
		})
		h.log.Error("Error while checking uniquess", l.Error(err))
		return
	}

	if checkUsername.Status {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Username already exists",
			},
		})
		return
	}

	checkEmail, err := h.serviceManager.UserService().CheckField(
		context.Background(), &pb.Check{
			Field: "email",
			Value: body.Email,
		},
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Message: err.Error(),
			},
		})
		h.log.Error("Error while checking email uniquess", l.Error(err))
		return
	}

	if checkEmail.Status {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Email already exists",
			},
		})
		return
	}

	code := etc.GenerateCode(7)

	_, err = h.serviceManager.EmailService().Send(
		context.Background(), &pbml.EmailText{
			Subject:   "Code for verification",
			Body:      code,
			Recipints: []string{body.Email},
		},
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Message: err.Error(),
			},
		})
		h.log.Error("Error while sending verification code to user", l.Error(err))
		return
	}
	fmt.Println("\n\n\n\n\n\n",err)

	data := models.User{
		First_name:    body.First_name,
		Last_name:     body.Last_name,
		Code:          code,
		Username:      body.Username,
		Password:      body.Password,
		Email:         body.Email,
		Profile_photo: body.Profile_photo,
		Phone_number:  body.Phone_number,
		Bio:           body.Bio,
		Address:       body.Address,
		Gender:        body.Gender,
	}

	bodyJSON, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Error while marshal user data for setting with ttl to redis",
			},
		})
		return
	}

	err = h.inMemoryStorage.SetWithTTL(code, string(bodyJSON), 86400)
	if err != nil {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Error while setting with ttl to redis",
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.RegisterResponseModel{
		Message: "Verification code has been sent to your email, please check and verify",
	})
}

// Verify ...
// @Summary Verify
// @Description returns access token
// @Tags register
// @Accept  json
// @Produce  json
// @Param code path string true "code"
// @Success 200 {object} models.User
// @Failure 404 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /verify/{code} [post]
func (h *handlerV1) Verify(c *gin.Context) {
	code := c.Param("code")

	// Getting code from redis

	user := models.User{}

	userJSON, err := redis.String(h.inMemoryStorage.Get(code))
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	// Checking whether received code is valid
	if code != user.Code {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Message: "ErrorCodeInvalidCode",
			},
		})
		h.log.Error("verification failed", l.Error(err))
		return
	}

	id, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Error while generating new uuid for user",
			},
		})
		return
	}

	h.jwtHandler = token.JWTHandler{
		SigninKey: h.cfg.SigninKey,
		Sub:       id.String(),
		Iss:       "user",
		Role:      "authorized",
		Aud:       []string{"nt"},
		Log:       h.log,
	}

	// Creating access and refresh tokens
	accessTokenString, refreshTokenString, err := h.jwtHandler.GenerateAuthJWT()
	if err != nil {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Error while generating tokens",
			},
		})
		return
	}

	// Creating hash of a password
	hashedPassword, err := etc.GeneratePasswordHash(user.Password)
	if err != nil {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Error while generating hash for password",
			},
		})
		return
	}

	checkEmail, err := h.serviceManager.UserService().CheckField(
		context.Background(), &pb.Check{
			Field: "email",
			Value: user.Email,
		},
	)
	if err != nil {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Error while checking field",
			},
		})
		return
	}

	if checkEmail.Status {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Mail already exists",
			},
		})
		return
	}

	// Creating new user
	resUser, err := h.serviceManager.UserService().Create(context.Background(), &pb.User{
		Id:           id.String(),
		Email:        user.Email,
		Password:     string(hashedPassword),
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		FirstName:    user.First_name,
		LastName:     user.Last_name,
		Username:     user.Username,
		Bio:          user.Bio,
		Address:      user.Address,
		Gender:       user.Gender,
		ProfilePhoto: user.Profile_photo,
		PhoneNumber:  user.Phone_number,
	})
	if err != nil {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Error while create new user",
			},
		})
		return
	}

	c.JSON(http.StatusOK, &models.VerifyResponseModel{
		ID:           resUser.Id,
		AccessToken:  resUser.AccessToken,
		RefreshToken: resUser.RefreshToken,
	})
}
