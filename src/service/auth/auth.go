package auth

import (
	"demo_1/src/config"
	"demo_1/src/constant"
	"demo_1/src/functions"
	"demo_1/src/middleware"
	"demo_1/src/model"
	"demo_1/src/repositories/user"
	"demo_1/src/types"
	"demo_1/src/util"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func Login(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	loginForm := types.LoginForm{}
	if err := c.BindJSON(&loginForm); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}

	hash := functions.MD5(loginForm.Password)
	_user := model.User{Name: loginForm.Name, PasswordHash: hash}
	if err := user.Index(&_user); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}

	token, err := middleware.NewJWT().CreateToken(middleware.CustomClaims{
		ID:             _user.ID,
		Name:           _user.Name,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(config.TokenExpireTime).Unix()},
	})
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}

	utilGin.Response(constant.SUCCESS, "登录成功", types.UserInfoForm{
		Model: _user.Model,
		Name:  _user.Name,
		Role:  _user.Role,
		Token: token,
	})
}

func Register(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	registerForm := types.RegisterForm{}
	if err := c.BindJSON(&registerForm); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}

	_user := model.User{}
	_user.PasswordHash = functions.MD5(registerForm.Password)
	_user.Name = registerForm.Name
	_, err := user.Add(&_user)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "注册成功", nil)
}

func fetch(url string, method string, body url.Values) (res []byte, err error) {
	var _res *http.Response = nil
	var _err error = nil
	if method == http.MethodPost {
		_res, _err = http.PostForm(url, body)
	} else if method == http.MethodGet {
		_res, _err = http.Get(url)
	}
	if _err != nil {
		return nil, _err
	}
	defer _res.Body.Close()
	_body, _err := ioutil.ReadAll(_res.Body)
	if _err != nil {
		return nil, _err
	}
	return _body, _err
}

func LoginByGitHub(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	code := c.Param("code")
	clientId := "7b961b417e4b3fc83488"
	clientSecret := "43720c59ded9b35c79b7131c4af178e4fe9651a2"
	body := url.Values{
		"client_id":     {clientId},
		"client_secret": {clientSecret},
		"code":          {code},
	}
	_res, _err := fetch("https://github.com/login/oauth/access_token", http.MethodPost, body)
	if _err != nil {
		utilGin.Response(constant.FAILED, _err.Error(), nil)
		return
	}
	var str = string(_res)
	data, _ := url.ParseQuery(str)
	token := data.Get("access_token")
	tokenType := data.Get("token_type")
	_url := fmt.Sprintf("https://api.github.com/user?access_token=%s&token_type=%s",
		token,
		tokenType,
	)

	_res, _err = fetch(_url, http.MethodGet, nil)
	if _res == nil {
		utilGin.Response(constant.FAILED, "登陆失败", nil)
		return
	}
	var githubUserInfo struct {
		Login  string `json:"login"`
		NodeId string `json:"node_id"`
	}
	_err = json.Unmarshal(_res, &githubUserInfo)
	_user := model.User{}
	_user.Name = githubUserInfo.Login
	_ = user.Index(&_user)
	if _user.ID == 0 {
		_user.PasswordHash = githubUserInfo.NodeId
		_, err := user.Add(&_user)
		if err != nil {
			utilGin.Response(constant.FAILED, err.Error(), nil)
			return
		}
	}

	token, err := middleware.NewJWT().CreateToken(middleware.CustomClaims{
		ID:             _user.ID,
		Name:           _user.Name,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(config.TokenExpireTime).Unix()},
	})
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}

	utilGin.Response(constant.SUCCESS, "登录成功", types.UserInfoForm{
		Model: _user.Model,
		Name:  _user.Name,
		Role:  _user.Role,
		Token: token,
	})
}
