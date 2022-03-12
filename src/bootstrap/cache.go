package bootstrap

import (
	"evlic/go-printer/src/conf"
	"time"

	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	goCache "github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

func InitCache() {
	log.Info("初始化缓存……")
	c := conf.Conf.Cache
	if c.Expiration == 0 {
		c.Expiration, c.CleanupInterval = 60, 120
	}
	goCahceClient := goCache.New(time.Minute*time.Duration(c.Expiration), time.Minute*time.Duration(c.CleanupInterval))
	goCacheStore := store.NewGoCache(goCahceClient, nil)
	conf.Cache = cache.New(goCacheStore)
}
