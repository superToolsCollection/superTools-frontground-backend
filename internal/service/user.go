package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"

	"superTools-frontground-backend/internal/dao"
	"superTools-frontground-backend/pkg/idGenerator"
)

/**
* @Author: super
* @Date: 2020-11-24 14:33
* @Description:
**/
type UserSignInRequest struct {
	UserName  string `form:"user_name" binding:"required,min=2,max=4294967295"`
	Password  string `form:"password" binding:"required,min=2,max=4294967295"`
	IPAddress string `form:"ip_address" binding:"required,min=7, max=15"`
}

type UserRegisterRequest struct {
	UserName string `form:"user_name" binding:"required,min=2,max=4294967295"`
	NickName string `form:"nick_name" binding:"required,min=2,max=4294967295"`
	Password string `form:"password" binding:"required,min=2,max=4294967295"`
}

type UserUpdateInfoRequest struct {
	ID       string `form:"id" binding:"required,min=2,max=4294967295"`
	UserName string `form:"user_name" binding:"required,min=2,max=4294967295"`
	NickName string `form:"nick_name" binding:"required,min=2,max=4294967295"`
	Password string `form:"password" binding:"required,min=2,max=4294967295"`
}

type User struct {
	ID           string `json:"id"`
	NickName     string `json:"nick_name"`
	UserName     string `json:"user_name"`
	HashPassword string `json:"-"`
}

type LoginUser struct {
	ID        string `json:"id"`
	UserName  string `json:"user_name"`
	IPAddress string `json:"ip_address"`
}

type IUserService interface {
	SignIn(param *UserSignInRequest) (*User, error)
	Register(param *UserRegisterRequest) (string, error)
	UpdateInfo(param *UserUpdateInfoRequest) error
}

type UserService struct {
	userDao dao.IUser
}

func (s *UserService) SignIn(param *UserSignInRequest) (*User, error) {
	user, err := s.userDao.SelectByUserName(param.UserName)
	if err != nil {
		return nil, errors.New("获取用户失败")
	}
	isOk, err := ValidatePassword(param.Password, user.HashPassword)
	if !isOk {
		return nil, err
	}
	return &User{
		ID:       user.ID,
		NickName: user.NickName,
		UserName: user.UserName,
	}, nil
}

func (s *UserService) Register(param *UserRegisterRequest) (string, error) {
	hashedPassword, err := GeneratePassword(param.Password)
	if err != nil {
		return "", err
	}
	userId, err := s.userDao.Insert(&dao.User{
		ID:           idGenerator.GenerateID(),
		UserName:     param.UserName,
		NickName:     param.NickName,
		HashPassword: string(hashedPassword),
	})
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (s *UserService) UpdateInfo(param *UserUpdateInfoRequest) error {
	hashedPassword, err := GeneratePassword(param.Password)
	if err != nil {
		return err
	}
	err = s.userDao.Update(&dao.User{
		ID:           param.ID,
		UserName:     param.UserName,
		NickName:     param.NickName,
		HashPassword: string(hashedPassword),
	})
	if err != nil {
		return err
	}
	return nil
}

func NewUserService(userDao dao.IUser) IUserService {
	return &UserService{userDao: userDao}
}

//将明文密码加密
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)

}

//验证登录密码是否正确
func ValidatePassword(userPassword string, hashed string) (isOK bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码校验错误")
	}
	return true, nil
}
