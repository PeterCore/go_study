package global

import (
	"my-blog-sevice/pkg/logger"
	"my-blog-sevice/pkg/setting"
)

var (
	ServerSetting   *setting.SeverSettings
	AppSetting      *setting.AppSettings
	DataBaseSetting *setting.DataBaseSettings
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettingS
)
