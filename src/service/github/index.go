package github

import (
	"demo_1/src/constant"
	"demo_1/src/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func fetch(url string, method string, body io.Reader) (res *http.Response, err error) {
	client := &http.Client{}
	var _res *http.Response = nil
	var _err error = nil
	if method == http.MethodPost {
		_res, _err = client.Post(url, "application/json", body)
	} else if method == http.MethodGet {
		_res, _err = client.Get(url)
	}
	return _res, _err
}

func Login(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	code := c.Query("code")
	clientId := "7b961b417e4b3fc83488"
	clientSecret := "43720c59ded9b35c79b7131c4af178e4fe9651a2"
	url := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		clientId,
		clientSecret,
		code,
	)
	_res, _err := fetch(url, http.MethodGet, nil)
	log.Println(_res, _err)
	utilGin.Response(constant.SUCCESS, "登陆成功", nil)
}
