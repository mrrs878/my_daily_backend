package test

import (
	"bytes"
	"demo_1/src/constant"
	"demo_1/src/controller"
	"demo_1/src/database"
	"demo_1/src/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xinliangnote/go-util/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var engine *gin.Engine

const accessToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODA5OTE1NTUsInVzZXJfaWQiOjE4LCJuYW1lIjoidGVzdDEifQ.900X9t3qveW0NGn3_FhYxHXlgsLAdl1BkjfKxNlBNyE"

func CreateRequest(method constant.HttpRequestMethod, url string, params interface{}) (status constant.ResultCode, resBody *util.ResS, returnError error) {
	if engine == nil {
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
		database.SetUpDatabase()
		controller.SetupRouter(engine)
	}

	var (
		rr        = httptest.NewRecorder()
		req       = new(http.Request)
		err error = nil
	)
	if params != nil {
		tmp, err := json.Encode(params)
		if err != nil {
			return constant.FAILED, nil, err
		}
		req, err = http.NewRequest(string(method), url, bytes.NewBuffer([]byte(tmp)))
	} else {
		req, err = http.NewRequest(string(method), url, nil)
	}
	if err != nil {
		return constant.FAILED, nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bear %s", accessToken))
	engine.ServeHTTP(rr, req)

	formatRes := util.ResS{}
	err = json.Decode(rr.Body.Bytes(), &formatRes)
	if err != nil {
		return constant.FAILED, nil, err
	}
	return constant.SUCCESS, &formatRes, nil
}

func CreateTest(t *testing.T, method constant.HttpRequestMethod, url string, params interface{}) {
	_, resBody, err := CreateRequest(method, url, params)
	if err != nil {
		t.Errorf("got error: %s", err.Error())
		return
	}
	if resBody != nil {
		if resBody.Code != constant.SUCCESS {
			t.Errorf("GetInfo test failed, got: %v, info: %s", resBody.Data, resBody.Message)
		}
	}
}
