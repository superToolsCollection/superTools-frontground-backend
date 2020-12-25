package service

import "errors"

/**
* @Author: super
* @Date: 2020-09-23 20:09
* @Description: 用于Auth入参验证与service代码
**/

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(
		param.AppKey,
		param.AppSecret,
	)
	if err != nil {
		return err
	}

	if auth.ID == "" {
		return nil
	}

	return errors.New("auth info does not exist")
}
