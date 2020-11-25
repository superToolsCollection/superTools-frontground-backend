package service

import (
	"context"

	"superTools-frontground-backend/global"
	"superTools-frontground-backend/internal/dao"
	"superTools-frontground-backend/pkg/otgorm"
)

/**
* @Author: super
* @Date: 2020-09-22 09:41
* @Description: 统一配置context
**/

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}
