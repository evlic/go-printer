package bootstrap

import (
	"evlic/go-printer/src/conf"
	"flag"

	log "github.com/sirupsen/logrus"
)

// InitLog init log
func InitLog() {
	if conf.Debug {
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	}

	log.SetFormatter(&log.TextFormatter{
		//DisableColors: true,
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           "2006-01-02 15:04:05",
		FullTimestamp:             true,
	})
	log.Infof("初始化「日志」")
}

func init() {
	flag.StringVar(&conf.ConfigFile, "conf", "data/config.json", "配置文件 url")
	flag.BoolVar(&conf.Debug, "debug", false, "是否开启debug，使用 -debug=true 开启")
	flag.Parse()
	InitLog()
}
