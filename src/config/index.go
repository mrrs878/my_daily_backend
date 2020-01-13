package config

var (
	ApiAuthConfig = map[string]map[string]string{
		"DEMO": {
			"md5": "IgkibX71IEf382PT",
			"aes": "IgkibX71IEf382PT",
			"rsa": "rsa/public.pem",
		},
	}
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
	AppGrpcLogName   = "src/log/" + AppName + "-grpc.log"

	ErrorNotifyOpen = -1

	DatabaseURl      = ""
	DatabasePort     = ""
	DatabaseUsername = ""
	DatabasePassword = ""
)
