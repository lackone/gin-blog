package global

import (
	"github.com/lackone/gin-blog/pkg/logger"
	"github.com/lackone/gin-blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	JwtSetting      *setting.JWTSetting
	EmailSetting    *setting.EmailSetting
	Logger          *logger.Logger
)
