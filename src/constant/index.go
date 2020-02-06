package constant

type ResultCode int8
type ErrorLevel int8
type HttpRequestMethod string

const (
	FAILED  ResultCode = -1
	SUCCESS ResultCode = 0
)

const (
	INFO ErrorLevel = 0
	MAIL ErrorLevel = 1
	SMS  ErrorLevel = 2
	WX   ErrorLevel = 3
)

const (
	GET    HttpRequestMethod = "GET"
	POST   HttpRequestMethod = "POST"
	PUT    HttpRequestMethod = "PUT"
	DELETE HttpRequestMethod = "DELETE"
	OPTION HttpRequestMethod = "OPTION"
)
