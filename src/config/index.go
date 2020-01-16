package config

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

	ErrorNotifyOpen = -1

	DatabaseURl      = "localhost"
	DatabasePort     = "3306"
	DatabaseName     = "e_book"
	DatabaseUsername = "root"
	DatabasePassword = "admin888"

	SignKey = "hello world"
)
