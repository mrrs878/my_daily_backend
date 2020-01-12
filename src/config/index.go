package config

var (
	ApiAuthConfig = map[string] map[string]string {
		"DEMO" : {
			"md5" : "IgkibX71IEf382PT",
			"aes" : "IgkibX71IEf382PT",
			"rsa" : "rsa/public.pem",
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

	AppAccessLogName = "log/" + AppName + "-access.log"
	AppErrorLogName  = "log/" + AppName + "-error.log"
	AppGrpcLogName   = "log/" + AppName + "-grpc.log"

	SystemEmailUser = "xinliangnote@163.com"
	SystemEmailPass = "" //密码或授权码
	SystemEmailHost = "smtp.163.com"
	SystemEmailPort = 465

	ErrorNotifyUser = "xinliangnote@163.com"

	ErrorNotifyOpen = -1

	JaegerHostPort = "127.0.0.1:6831"

	JaegerOpen = 1
)
