package model

import (
	"fmt"
	"github.com/lackone/gin-blog/global"
	"github.com/lackone/gin-blog/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/tracing"
)

type Model struct {
	Id        uint   `gorm:"primaryKey;autoIncrement;size:32;not null;comment:ID" json:"id"`
	Created   uint   `gorm:"size:32;not null;autoCreateTime;comment:创建时间" json:"created"`
	CreatedBy string `gorm:"size:32;not null;comment:创建人" json:"created_by"`
	Updated   uint   `gorm:"size:32;not null;autoUpdateTime;comment:更新时间" json:"updated"`
	UpdatedBy string `gorm:"size:32;not null;comment:更新人" json:"updated_by"`
	Deleted   uint   `gorm:"size:32;not null;comment:删除时间" json:"deleted"`
	IsDel     uint8  `gorm:"size:1;not null;default:0;comment:是否删除" json:"is_del"`
}

func NewDb(setting *setting.DatabaseSetting) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		setting.Username,
		setting.Password,
		setting.Host,
		setting.DBName,
		setting.Charset,
		setting.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: setting.TablePrefix,
		},
	})
	if err != nil {
		return nil, err
	}
	if err = db.Use(tracing.NewPlugin()); err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.Logger = logger.Default.LogMode(logger.Info)
	}
	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDb.SetMaxIdleConns(setting.MaxIdleConns)
	sqlDb.SetMaxOpenConns(setting.MaxOpenConns)
	return db, nil
}
