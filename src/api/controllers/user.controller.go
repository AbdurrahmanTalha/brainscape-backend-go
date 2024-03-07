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
		return
	}

	err = h.service.Register(req)
	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse("", true, helper.Success))
}

func (h *UserController) Login(c *gin.Context) {
	req := new(dto.LoginRequest)
	err := c.ShouldBindJSON(&req)

	if err != nil {
		fmt.Println(err, "1");
		return
	}

	token, err := h.service.Login(req)
	if err != nil {
		fmt.Println(err, "2");
		return;
	}
	
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(token, true, helper.Success))

}
