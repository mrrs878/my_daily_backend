package auth

import (
	"demo_1/src/middleware"
	"demo_1/src/model"
	"demo_1/src/repositories/user"
	"demo_1/src/security"
	"demo_1/src/types"
	"demo_1/src/util"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	loginForm := types.LoginForm{}
	if err := c.BindJSON(&loginForm); err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}

	hash := security.MD5(loginForm.Password)
	_user := model.User{Name: loginForm.Name, PasswordHash: hash}
	if err := user.Index(&_user); err != nil {
		utilGin.Response(-1, err.Error(), nil)
	}

	token, err := middleware.NewJWT().CreateToken(middleware.CustomClaims{ID: _user.ID, Name: _user.Name})
	if err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}

	utilGin.Response(0, "登录成功", types.UserInfoForm{
		Model:  _user.Model,
		Name:   _user.Name,
		Emails: _user.Emails,
		Role:   _user.Role,
		Token:  token,
	})
}

func Register(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	registerForm := types.RegisterForm{}
	if err := c.BindJSON(&registerForm); err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}

	_user := model.User{}
	_user.PasswordHash = security.MD5(registerForm.Password)
	_user.Name = registerForm.Name
	_, err := user.Add(&_user)
	if err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}
	utilGin.Response(0, "注册成功", nil)
}
