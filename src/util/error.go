package util

import (
	"demo_1/src/config"
	"demo_1/src/constant"
	"github.com/xinliangnote/go-util/json"
	"log"
	"os"
	"time"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func ErrorNew(text string) error {
	alarm(config.ErrorLevel, text)
	return &errorString{text}
}

func alarm(level constant.ErrorLevel, text string) {
	if level == constant.INFO {
		if f, err := os.OpenFile(config.AppErrorLogName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
			log.Println(err)
		} else {
			errorLogMap := make(map[string]interface{})
			errorLogMap["time"] = time.Now().Format("2020/02/04 - 13:45:05")
			errorLogMap["info"] = text

			errorLogJson, _ := json.Encode(errorLogMap)
			_, _ = f.WriteString(errorLogJson + "\n")
		}
	}
}
