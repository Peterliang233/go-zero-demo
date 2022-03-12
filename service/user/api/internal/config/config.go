package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	// 配置jwt进行鉴权
	Auth struct { // jwt鉴权配置
		AccessSecret string // jwt密钥
		AccessExpire int64  // 有效期，单位：秒
	}
	Mysql struct {
		DataSource string
	}

	//UserRpc zrpc.RpcClientConf

	CacheRedis cache.CacheConf
}
