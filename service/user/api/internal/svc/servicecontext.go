package svc

import (
	"book/service/user/api/internal/config"
	"book/service/user/model"
	"book/service/user/rpc/userclient"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	UserRpc   userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	coon := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(coon, c.CacheRedis),
		//UserRpc:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
