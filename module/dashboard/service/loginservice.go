package service

import (
	"encoding/json"
	"net/http"

	constant "github.com/MufidJamaluddin/simplewebgo/constant"
	model "github.com/MufidJamaluddin/simplewebgo/model"
	security "github.com/MufidJamaluddin/simplewebgo/security"
)

// LoginServiceRoute Login dengan JWT
func LoginServiceRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Unsupported http method", http.StatusBadRequest)
		return
	}

	var credential model.CredentialModel
	err := json.NewDecoder(r.Body).Decode(&credential)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	// PENGECEKAN USERNAME DAN PASSWORD DI DATABASE!
	//
	// ada := security.CheckUser(credential)
	// if ada {
	//	  http.Error(w, "Username atau Password Salah!", http.StatusBadRequest)
	//	  return
	// }
	//
	// END PENGECEKAN!

	token := security.GenerateJWTToken(credential.Username, credential.Group)
	signedToken, err := token.SignedString(constant.JWTSignatureKey)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenString, _ := json.Marshal(model.TokenModel{Token: signedToken})
	w.Write([]byte(tokenString))
}
