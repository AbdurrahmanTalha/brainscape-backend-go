package controllers

import (
	"fmt"
	"net/http"

	"github.com/AbdurrahmanTalha/brainscape-backend-go/api/dto"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/api/helper"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/config"
	services "github.com/AbdurrahmanTalha/brainscape-backend-go/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(cfg *config.Config) *UserController {
	service := services.NewUserService(cfg)
	return &UserController{
		service: service,
	}
}

func (h *UserController) Register(c *gin.Context) {
	req := new(dto.RegisterUserRequest)
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(
			http.StatusCreated,
			helper.GenerateBaseResponseWithError(
				http.StatusBadRequest,
				err,
				"Failed to create user",
			),
		)
		return
	}

	user, err := h.service.Register(req)

	if err != nil {
		c.JSON(
			http.StatusCreated,
			helper.GenerateBaseResponseWithError(
				http.StatusBadRequest,
				err,
				"Failed to create user",
			),
		)
		return
	}

	fmt.Printf("%+v", user)

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(true, "Successfully created user", http.StatusCreated, &user))
}

func (h *UserController) Login(c *gin.Context) {
	req := new(dto.LoginRequest)
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(
			http.StatusCreated,
			helper.GenerateBaseResponseWithError(
				http.StatusBadRequest,
				err,
				"Failed to bind json",
			),
		)
		return
	}

	token, err := h.service.Login(req)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			helper.GenerateBaseResponseWithError(
				http.StatusBadRequest,
				err,
				"Failed to login user",
			),
		)
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(true, "Successfully logged in user", http.StatusCreated, token))
}

func (h *UserController) GetAllUsers(c *gin.Context) {
	result, err := h.service.GetAllUsers()

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			helper.GenerateBaseResponseWithError(
				http.StatusBadRequest,
				err,
				"Failed to find all users",
			),
		)
		return
	}

	c.JSON(http.StatusOK,
		helper.GenerateBaseResponse(
			true,
			"Successfully found all users",
			http.StatusOK,
			result,
		),
	)
}
