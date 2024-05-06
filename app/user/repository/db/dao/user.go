package dao

import (
	"context"
	"github.com/reoden/todolist/app/user/repository/db/model"
	"gorm.io/gorm"
	"log"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}

	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) CreateUser(user *model.User) (err error) {
	return dao.Model(&model.User{}).Create(&user).Error
}

func (dao *UserDao) FindUserByUserName(userName string) (r *model.User, err error) {
	err = dao.Model(&model.User{}).
		Where("userName = ? ", userName).Find(&r).Error
	if err != nil {
		log.Println("FindUserNameByUserName err:", err)
		return nil, err
	}

	return
}
