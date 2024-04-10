package services

import (
	"context"
	"fmt"

	"github.com/AbdurrahmanTalha/brainscape-backend-go/common"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/config"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/data/db"
	"gorm.io/gorm"
)

type preload struct {
	string
}

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	database *gorm.DB
	preloads []preload
}

func NewBaseService[T any, Tc any, Tu any, Tr any](cfg *config.Config) *BaseService[T, Tc, Tu, Tr] {
	return &BaseService[T, Tc, Tu, Tr]{
		database: db.GetDB(),
	}
}

func (s *BaseService[T, Tc, Tu, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) {
	data, _ := common.TypeConverter[T](req)
	transaction := s.database.WithContext(ctx).Begin()
	err := transaction.Create(data).Error

	if err != nil {
		transaction.Rollback()
		return nil, err
	}

	transaction.Commit()
	bm, _ := common.TypeConverter[Tr](data)

	return bm, nil
}

func (s *BaseService[T, Tc, Tu, Tr]) GetById(ctx context.Context, id int) (*Tr, error) {
	model := new(T)
	db := Preload(s.database, s.preloads)

	err := db.Where("id = ?", id).First(model).Error
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	fmt.Println(model)
	return common.TypeConverter[Tr](model)
}

func Preload(db *gorm.DB, preloads []preload) *gorm.DB {
	for _, item := range preloads {
		db = db.Preload(item.string)
	}

	return db
}

func (s *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, id int, req *Tu) (*Tr, error) {
	updateMap, _ := common.TypeConverter[map[string]interface{}](req)
	snakeMap := map[string]interface{}{}
	fmt.Println(updateMap)
	fmt.Println(snakeMap)
	for k, v := range *updateMap {
		snakeMap[common.ToSnakeCase(k)] = v
	}
	
	model := new(T)
	tx := s.database.WithContext(ctx).Begin()
	if err := tx.Model(model).Where("id = ?", id).Updates(snakeMap).Error; err != nil {
		tx.Rollback()

		return nil, err
	}
	tx.Commit()

	return s.GetById(ctx, id)
}
