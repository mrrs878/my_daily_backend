package functions

import (
	"crypto/md5"
	"demo_1/src/middleware"
	"demo_1/src/model"
	"demo_1/src/repositories/user"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
)

func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func GetUserIDFormContext(c *gin.Context) (uint, error) {
	value, exists := c.Get("claims")
	if !exists {
		return 0, errors.New("token失效")
	}
	result, ok := (value).(*middleware.CustomClaims)
	if !ok {
		return 0, errors.New("token失效")
	}
	return result.ID, nil
}

func GetUserInfoFromContext(c *gin.Context, _user *model.User) error {
	value, exists := c.Get("claims")
	if !exists {
		return errors.New("token失效")
	}
	result, ok := (value).(*middleware.CustomClaims)
	if !ok {
		return errors.New("token失效")
	}
	_user.ID = result.ID
	if err := user.Index(_user); err != nil {
		return err
	}
	return nil
}
