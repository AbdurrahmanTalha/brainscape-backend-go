package services

import (
	"github.com/AbdurrahmanTalha/brainscape-backend-go/api/dto"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/common"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/config"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/data/db"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/data/models"
	"gorm.io/gorm"
)

type UserService struct {
	cfg      *config.Config
	database *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDB()

	return &UserService{
		cfg:      cfg,
		database: database,
	}
}

func (s *UserService) Register(req *dto.RegisterUserRequest) error {
	u := models.User{FullName: req.FullName, Email: req.Email, Password: req.Password, Role: "student"}

	exists, err := s.isExistByEmail(req.Email)

	if err != nil {
		return err
	}

	if exists {
		return err
	}

	u.Password = common.HashPassword(u.Password)

	transaction := s.database.Begin()
	err = transaction.Create(&u).Error
	if err != nil {
		transaction.Rollback()
		return err
	}
	transaction.Commit()
	return nil
}

func (s *UserService) isExistByEmail(email string) (bool, error) {
	var count int64
	if err := s.database.Model(&models.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *UserService) Login(req *dto.LoginRequest) (*dto) {
	var user models.User;
	err := s.database.Model(&models.User{}).Where("email = ?", req.Email).Find(&user).Error;
	if err != nil {
		return nil, err;
	}
	return ""
}
