package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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
				false,
				http.StatusBadRequest,
				err,
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
				false,
				http.StatusBadRequest,
				err,
			),
		)
		return
	}

	fmt.Printf("%+v", user)

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(
		&user,
		true,
		http.StatusCreated,
	))
}

func (h *UserController) Login(c *gin.Context) {
	req := new(dto.LoginRequest)
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(
			http.StatusCreated,
			helper.GenerateBaseResponseWithError(
				http.StatusBadRequest,
				false,
				http.StatusBadRequest,
				err,
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
				false,
				http.StatusBadRequest,
				err,
			),
		)
		return
	}
	/* result any, success bool, resultCode ResultCode */
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(token, true, http.StatusCreated))
}

func (h *UserController) GetAllUsers(c *gin.Context) {
	result, err := h.service.GetAllUsers()

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			helper.GenerateBaseResponseWithError(
				http.StatusBadRequest,
				false,
				http.StatusBadRequest,
				err,
			),
		)
		return
	}

	c.JSON(http.StatusOK,
		helper.GenerateBaseResponse(
			result,
			true,
			http.StatusOK,
		),
	)
}

func (h *UserController) GetSpecificUser(c *gin.Context) {
	id := c.Param("id")
	convertedId, err := strconv.ParseUint(id, 0, 64)

	if err != nil {
		helper.GenerateBaseResponseWithError(
			http.StatusBadRequest,
			false,
			http.StatusBadRequest,
			err,
		)
	}

	result, err := h.service.GetSpecificUser(uint(convertedId))

	if err != nil {
		helper.GenerateBaseResponseWithError(
			http.StatusBadRequest,
			false,
			http.StatusBadRequest,
			err,
		)
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(result, true, http.StatusOK))
}

func (h *UserController) UpdateSpecificUser(c *gin.Context) {
	Update(c, h.service.Update)
}
