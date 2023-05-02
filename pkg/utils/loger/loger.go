package loger

import (
	log "github.com/cihub/seelog"
	"github.com/yin-zt/itsm-workflow/pkg/config"
	"os"
)

var (
	LoggerAcc log.LoggerInterface
	LoggerOpe log.LoggerInterface
)

func GetLoggerAcc() log.LoggerInterface {
	os.MkdirAll("/var/loger/", 0777)
	logger, err := log.LoggerFromConfigAsBytes([]byte(config.LogAccessConfigStr))
	if err != nil {
		log.Error("init loger fail")
		os.Exit(1)
	}
	LoggerAcc = logger
	return LoggerAcc
}

func GetLoggerOperate() log.LoggerInterface {
	os.MkdirAll("/var/loger/", 0777)
	logger, err := log.LoggerFromConfigAsBytes([]byte(config.LogOperateConfigStr))
	if err != nil {
		log.Error("init loger fail")
		os.Exit(1)
	}
	LoggerOpe = logger
	return LoggerOpe
}

func GetCmdbLogger() log.LoggerInterface {
	os.MkdirAll("/var/loger/", 0777)
	logger, err := log.LoggerFromConfigAsBytes([]byte(config.LogOperateCmdbStr))
	if err != nil {
		log.Error("init cmdb loger fail")
		os.Exit(1)
	}
	return logger
}
