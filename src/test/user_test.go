package test

import (
	"demo_1/src/constant"
	"testing"
)

func TestGetInfo(t *testing.T) {
	CreateTest(t, constant.GET, "/user/", nil)
}

func TestWriteOff(t *testing.T) {
	CreateTest(t, constant.DELETE, "/user/3", nil)
}

func TestWriteOffSelf(t *testing.T) {
	CreateTest(t, constant.DELETE, "/user/", nil)
}
