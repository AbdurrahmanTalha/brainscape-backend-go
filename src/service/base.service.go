package services

import (
	"context"

	"github.com/AbdurrahmanTalha/brainscape-backend-go/common"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/config"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/data/db"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/data/models"
	"gorm.io/gorm"
)

type preload struct {
	string
}

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	Database *gorm.DB
	Preloads []preload
}

func NewBaseService[T any, Tc any, Tu any, Tr any](cfg *config.Config) *BaseService[T, Tc, Tu, Tr] {
	return &BaseService[T, Tc, Tu, Tr]{
		Database: db.GetDB(),
	}
}

func (s *BaseService[T, Tc, Tu, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) {
	data, _ := common.TypeConverter[T](req)
	transaction := s.Database.WithContext(ctx).Begin()
	err := transaction.Create(data).Error

	if err != nil {
		transaction.Rollback()
		return nil, err
	}

	transaction.Commit()
	bm, _ := common.TypeConverter[models.BaseModel](data)

	return s.GetById(ctx, bm.Id)
}

func (s *BaseService[T, Tc, Tu, Tr]) GetById(ctx context.Context, id int) (*Tr, error) {
	model := new(T)
	db := Preload(s.Database, s.Preloads)

	err := db.Where("id = ? and deleted_by is null", id).First(model).Error

	if err != nil {
		return nil, err
	}

	return common.TypeConverter[Tr](model)
}


func Preload(db *gorm.DB, preloads []preload) *gorm.DB {
	for _, item := range preloads {
		db = db.Preload(item.string);
	}
	
	return db;
}