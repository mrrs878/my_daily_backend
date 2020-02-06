package test

import (
	"demo_1/src/constant"
	"demo_1/src/types"
	"testing"
)

func TestLogin(t *testing.T) {
	CreateTest(t, constant.POST, "/auth/login", types.LoginForm{Name: "admin", Password: "admin888"})
}

func TestRegister(t *testing.T) {
	CreateTest(t, constant.POST, "/auth/register", types.RegisterForm{Name: "test1", Password: "test111"})
}
