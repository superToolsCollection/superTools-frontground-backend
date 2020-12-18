package global

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"superTools-frontground-backend/pkg/mq"
)

/**
* @Author: super
* @Date: 2020-09-18 08:51
* @Description: 全局配置DB
**/

var (
	DBEngine    *gorm.DB
	RedisEngine *redis.Pool
	RabbitMQEngine *mq.RabbitMQ
)
