package common

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/yin-zt/itsm-workflow/pkg/config"
	"github.com/yin-zt/itsm-workflow/pkg/models/order"
	"github.com/yin-zt/itsm-workflow/pkg/utils/loger"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	defer log.Flush()
	for _, dir := range config.ConstSysDirs {
		os.Mkdir(dir, 0777)
	}
	log.ReplaceLogger(loger.GetLoggerOperate())
	log.Info("database logger init successfully")
}

// 全局mysql数据库变量
var DB *gorm.DB

// 初始化mysql数据库
func InitMysql() {
	defer log.Flush()
	var (
		resp any
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		config.MysqlUsername,
		config.MysqlPassword,
		config.MysqlHost,
		config.MysqlPort,
		config.MysqlDatabase,
		config.MysqlCharset,
		config.MysqlCollation,
		config.MysqlQuery,
	)
	// 隐藏密码
	showDsn := fmt.Sprintf(
		"%s:******@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		config.MysqlUsername,
		config.MysqlHost,
		config.MysqlPort,
		config.MysqlDatabase,
		config.MysqlCharset,
		config.MysqlCollation,
		config.MysqlQuery,
	)
	// Log.Info("数据库连接DSN: ", showDsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用外键(指定外键时不会在mysql创建真实的外键约束)
		DisableForeignKeyConstraintWhenMigrating: true,
		//// 指定表前缀
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix: config.Conf.Mysql.TablePrefix + "_",
		//},
	})
	if err != nil {
		log.Errorf("初始化mysql数据库异常: %v", err)
		resp = fmt.Errorf("初始化mysql数据库异常: %v", err)
		panic(resp)
	}
	// 开启mysql日志
	if config.MysqlLogMode {
		db.Debug()
	}
	// 全局DB赋值
	DB = db
	// 自动迁移表结构
	dbAutoMigrate()
	log.Infof("初始化mysql数据库完成! dsn: %s", showDsn)
}

// 自动迁移表结构
func dbAutoMigrate() {
	_ = DB.AutoMigrate(
		&order.T_Order{},
		//&order.Big_T_Order{},
	)
}
