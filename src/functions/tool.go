package functions

import (
	"crypto/md5"
	"demo_1/src/middleware"
	"demo_1/src/model"
	"demo_1/src/repositories/user"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func GetUserIDFormContext(c *gin.Context) (uint, error) {
	tokenTmp := strings.Split(c.Request.Header.Get("Authorization"), " ")
	if len(tokenTmp) != 2 {
		return 0, errors.New("请求未携带token")
	}
	token := tokenTmp[1]
	log.Print("get token: ", token)

	j := middleware.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		return 0, err
	}
	return claims.ID, nil
}

func GetUserInfoFromContext(c *gin.Context) (*model.User, error) {
	id, err := GetUserIDFormContext(c)
	if err != nil {
		return nil, err
	}
	_user := model.User{}
	_user.ID = id
	err = user.Index(&_user)
	return &_user, err
}
