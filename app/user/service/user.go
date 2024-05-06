package service

import (
	"context"
	"errors"
	"github.com/reoden/todolist/app/user/repository/db/dao"
	"github.com/reoden/todolist/app/user/repository/db/model"
	"github.com/reoden/todolist/idl/pb"
	"github.com/reoden/todolist/pkg/ecode"
	"sync"
)

type UserSrv struct {
}

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

// GetUserSrv lazy-loading solve the problem of parallel
func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (u *UserSrv) UserLogin(ctx context.Context, req *pb.UserReq, resp *pb.UserResp) (err error) {
	resp.Code = uint32(ecode.Success)

	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		return
	}

	if user.ID == 0 {
		resp.Code = uint32(ecode.Error)
		return errors.New("user is not exits")
	}

	if !user.CheckPassword(req.Password) {
		resp.Code = uint32(ecode.Error)
		return errors.New("password error")
	}

	resp.UserDetail = BuildUser(user)
	return nil
}

func (u *UserSrv) UserRegister(ctx context.Context, req *pb.UserReq, resp *pb.UserResp) (err error) {
	resp.Code = uint32(ecode.Success)

	if req.Password != req.PasswordConfirm {
		resp.Code = uint32(ecode.Error)
		return errors.New("password do not equal to password confirm")
	}

	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		return
	}

	if user.ID > 0 {
		resp.Code = uint32(ecode.Error)
		return errors.New("user name is already exits")
	}

	user = &model.User{
		UserName: req.UserName,
	}

	if err = user.SetPassword(req.Password); err != nil {
		resp.Code = uint32(ecode.Error)
		return
	}

	if err = dao.NewUserDao(ctx).CreateUser(user); err != nil {
		resp.Code = uint32(ecode.Error)
		return
	}

	return nil

	return nil
}

func BuildUser(user *model.User) *pb.UserModel {
	return &pb.UserModel{
		Id:        uint32(user.ID),
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}
}
