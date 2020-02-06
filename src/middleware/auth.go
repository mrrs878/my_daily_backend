package middleware

import (
	"demo_1/src/config"
	"demo_1/src/constant"
	"demo_1/src/util"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"time"
)

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	jwt.StandardClaims
	ID   uint   `json:"user_id"`
	Name string `json:"name"`
}

var CustomValidateMsg = map[uint32]string{
	jwt.ValidationErrorExpired:     "token已过期",
	jwt.ValidationErrorMalformed:   "token格式错误",
	jwt.ValidationErrorNotValidYet: "token未生效",
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		utilGin := util.GinS{Ctx: c}
		tokenTmp := strings.Split(c.Request.Header.Get("Authorization"), " ")
		if len(tokenTmp) != 2 {
			utilGin.Response(constant.FAILED, "请求未携带token，无权限访问", nil)
			c.Abort()
			return
		}

		token := tokenTmp[1]
		log.Print("get token: ", token)

		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			utilGin.Response(constant.FAILED, err.Error(), nil)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

func GetSignKey() string {
	return config.TokenSignKey
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			return nil, errors.New(CustomValidateMsg[ve.Errors])
		}
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, nil
}

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(config.TokenExpireTime).Unix()
		return j.CreateToken(*claims)
	}
	return "", errors.New(CustomValidateMsg[jwt.ValidationErrorMalformed])
}
