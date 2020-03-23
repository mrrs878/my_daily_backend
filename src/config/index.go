package config

import (
	"demo_1/src/constant"
	"time"
)

const (
	AppMode = "release" //debug or release
	AppPort = ":9090"
	AppName = "go-gin-api"

	AppSignExpiry = "120"

	AppRsaPrivateFile = "rsa/private.pem"

	AppReadTimeout  = 120
	AppWriteTimeout = 120

	AppAccessLogName = "src/log/" + AppName + "-access.log"
	AppErrorLogName  = "src/log/" + AppName + "-error.log"

	ErrorLevel = constant.INFO

	DatabaseURl      = "118.190.140.52"
	DatabasePort     = "3306"
	DatabaseName     = "my_daily"
	DatabaseUsername = "root"
	DatabasePassword = "admin888_remote"

	TokenSignKey    = "hello world"
	TokenExpireTime = 2400 * time.Hour

	QueryPageSize = 20
)
