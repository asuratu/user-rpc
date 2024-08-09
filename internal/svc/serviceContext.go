package svc

import (
	"errors"

	"github.com/asuratu/user-rpc/internal/config"
	"github.com/asuratu/user-rpc/model/userDao"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/syncx"
)

type ServiceContext struct {
	Config    config.Config
	Cache     cache.Cache
	Redis     *redis.Redis
	UserModel userDao.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化缓存
	ca := cache.New(c.Cache, syncx.NewSingleFlight(), cache.NewStat("dc"), errors.New("cache not found"))

	// 初始化redis
	rs := redis.MustNewRedis(redis.RedisConf{
		Host: c.Redis.Host,
		Type: c.Redis.Type,
	})

	// 初始化数据库
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: userDao.NewUserModel(sqlConn, c.Cache),
		Cache:     ca,
		Redis:     rs,
	}
}
