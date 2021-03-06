package jwt

import (
	"fmt"
	"github.com/gofuncchan/ginger/config"
	"github.com/gofuncchan/ginger/logger"
	"go.uber.org/zap"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func JwtInit() {
	JwtService = new(TokenService)
	fmt.Println("JWT service init ready!")

}

type ITokenService interface {
	Encode(user TokenUserClaim) (string, error)
	Decode(tokenString string) (*CustomerClaim, error)
}

// 编解码私钥，在生产环境中，该私钥请使用生成器生成，并妥善保管，此处使用简单字符串。
var privateKey = []byte(config.TokenPrivateKey)
var JwtService *TokenService

// JWT中携带的用户个人信息
type TokenUserClaim struct {
	Id     int64  `json:id`
	Name   string `json:name`
	Avatar string `json:avatar`
}

// 聚合jwt内部实现的Claims
type CustomerClaim struct {
	TokenUserClaim
	*jwt.StandardClaims
}

// 实现token服务
type TokenService struct{}

// 传入用户信息编码成token
func (tks *TokenService) Encode(user TokenUserClaim) (string, error) {
	// privateKey, _ := base64.URLEncoding.DecodeString(string(privateKey))

	// 设置超时时间
	expTime := time.Now().Add(time.Hour * 24 * 3).Unix()

	// 设置Claim
	customer := CustomerClaim{user, &jwt.StandardClaims{ExpiresAt: expTime}}

	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customer)

	return token.SignedString(privateKey)

}

// token字符串解码成用户信息
func (tks *TokenService) Decode(tokenString string) (*CustomerClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomerClaim{}, func(token *jwt.Token) (interface{}, error) {
		// return base64.URLEncoding.DecodeString(string(privateKey))
		return privateKey, nil
	})

	if err != nil {
		logger.ZapLogger.Error("JWT Decode Wrong",
			zap.String("path", "util/jwt.TokenService.Decode"),
			zap.String("warming", err.Error()),
			zap.String("receive_token", tokenString),
		)
		return nil, err
	}

	if claim, ok := token.Claims.(*CustomerClaim); ok && token.Valid {
		return claim, nil
	} else {
		return nil, err
	}
}
