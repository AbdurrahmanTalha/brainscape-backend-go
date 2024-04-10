package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/AbdurrahmanTalha/brainscape-backend-go/api/dto"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/api/helper"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/common"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/config"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/data/db"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/data/models"
	"gorm.io/gorm"
)

type UserService struct {
	cfg      *config.Config
	database *gorm.DB
	base     *BaseService[models.User, dto.RegisterUserRequest, dto.UpdateUserRequest, dto.UserResponse]
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDB()

	return &UserService{
		cfg:      cfg,
		database: database,
		base: &BaseService[models.User, dto.RegisterUserRequest, dto.UpdateUserRequest, dto.UserResponse]{
			database: database,
			preloads: []preload{},
		},
	}
}
func (s *UserService) Register(req *dto.RegisterUserRequest) (*models.User, error) {
	u := models.User{FullName: req.FullName, Email: req.Email, Role: "student"}

	exists, _ := s.isExistByEmail(req.Email)

	if exists {
		return nil, errors.New("user already exists")
	}

	u.Password = []byte(common.HashPassword(string(req.Password)))
	transaction := s.database.Begin()
	result := transaction.Create(&u)

	if result.Error != nil {
		transaction.Rollback()
		fmt.Println(result.Error)
		return nil, result.Error
	}
	transaction.Commit()

	return &u, nil
}

func (s *UserService) Login(req *dto.LoginRequest) (*dto.TokenDetail, error) {
	var user models.User
	err := s.database.Model(&models.User{}).Where("email = ?", req.Email).Find(&user).Error

	if err != nil {
		return nil, err
	}

	err = common.ComparePassword(string(user.Password), req.Password)

	if err != nil {
		return nil, errors.New("password does'nt match")
	}

	tokenData := map[string]interface{}{
		"email":    user.Email,
		"role":     user.Role,
		"fullName": user.FullName,
	}

	accessToken, err := helper.GenerateJSONToken(tokenData, s.cfg.JWT.AccessTokenSecret, s.cfg.JWT.AccessTokenExpiresIn)

	if err != nil {
		return nil, err
	}

	refreshToken, err := helper.GenerateJSONToken(tokenData, s.cfg.JWT.RefreshTokenSecret, s.cfg.JWT.RefreshTokenExpiresIn)

	if err != nil {
		return nil, err
	}

	return &dto.TokenDetail{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := s.database.Find(&users).Error; err != nil {
		return nil, errors.New("something went wrong")
	}

	return users, nil
}

func (s *UserService) GetSpecificUser(id uint) (*models.User, error) {
	var result models.User

	if err := s.database.First(&result, "id = ?", id).Error; err != nil {
		fmt.Println("Error retrieving user:", err)

		return nil, errors.New("failed to retrieve user")
	}

	return &result, nil
}

func (s *UserService) isExistByEmail(email string) (bool, error) {
	var count int64
	err := s.database.Model(&models.User{}).Where("email = ?", email).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *UserService) Update(ctx context.Context, id int, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	return s.base.Update(ctx, id, req)
}
