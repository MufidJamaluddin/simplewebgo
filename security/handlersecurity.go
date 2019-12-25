package security

/**
 *	HandlerSecurity Handler
 *
 *	@author Mufid Jamaluddin
 **/

import (
	"fmt"
	"net/http"
	"strings"

	model "github.com/MufidJamaluddin/simplewebgo/model"
	jwt "github.com/dgrijalva/jwt-go"
)

// SecureHandleFunc ...
type SecureHandleFunc func(http.ResponseWriter, *http.Request, *model.UserInfo)

// HandlerSecurity mengamankan page yang wajib login
func HandlerSecurity(phandler SecureHandleFunc) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}

		strtoken := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := ParseJWTToken(strtoken)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userInfo := &model.UserInfo{
			Username: fmt.Sprintf("%v", claims["Username"]),
			Group:    fmt.Sprintf("%v", claims["Group"]),
		}

		phandler(w, r, userInfo)
	}

	return http.HandlerFunc(fn)
}
