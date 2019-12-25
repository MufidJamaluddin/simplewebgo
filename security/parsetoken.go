package security

import (
	"fmt"
	constant "github.com/MufidJamaluddin/simplewebgo/constant"
	"github.com/dgrijalva/jwt-go"
)

func validateJWTSigningMethod(token *jwt.Token) (interface{}, error) {
	if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Signing method invalid")
	} else if method != constant.JWTSigningMethod {
		return nil, fmt.Errorf("Signing method invalid")
	}

	return constant.JWTSignatureKey, nil
}

// ParseJWTToken ..
func ParseJWTToken(strtoken string) (*jwt.Token, error) {
	return jwt.Parse(strtoken, validateJWTSigningMethod)
}
