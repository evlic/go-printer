package conf

import (
	"github.com/eko/gocache/v2/cache"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

var (
	Conf       *Config
	DB         *gorm.DB
	Cache      *cache.Cache
	Cron       *cron.Cron
	Debug      bool
	ConfigFile string
)
