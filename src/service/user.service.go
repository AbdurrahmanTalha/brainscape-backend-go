package services

import (
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
	/* base     *BaseService[models.User, dto.RegisterUserRequest, dto.UpdatePropertyRequest, dto.PropertyResponse] */
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDB()

	return &UserService{
		cfg:      cfg,
		database: database,
	}
}
func (s *UserService) Register(req *dto.RegisterUserRequest) (any, error) {
	u := models.User{FullName: req.FullName, Email: req.Email, Role: "student"}

	exists, _ := s.isExistByEmail(req.Email)

	if exists {
		return nil, errors.New("user already exists")
	}

	u.Password = []byte(common.HashPassword(string(u.Password)))
	transaction := s.database.Begin()
	result := transaction.Create(&u)

	if result.Error != nil {
		transaction.Rollback()
		fmt.Println(result.Error)
		return nil, result.Error
	}
	transaction.Commit()

	return result, nil
}

func (s *UserService) Login(req *dto.LoginRequest) (*dto.TokenDetail, error) {
	var user models.User
	err := s.database.Model(&models.User{}).Where("email = ?", req.Email).Find(&user).Error

	if err != nil {
		return nil, err
	}

	err = common.ComparePassword(string(user.Password), req.Password)
	
	if err != nil {
		return nil, err
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

func (s *UserService) GetAllUsers() {
	var users []models.User
	if err := s.database.Find(&users).Error; err != nil {
		fmt.Printf("Error retrieving users: %v", err)
	}

	for _, user := range users {
		fmt.Printf("ID: %d, Full Name: %s, Email: %s", user.ID, user.FullName, user.Email)
	}
}

func (s *UserService) isExistByEmail(email string) (bool, error) {
	var count int64
	err := s.database.Model(&models.User{}).Where("email = ?", email).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
