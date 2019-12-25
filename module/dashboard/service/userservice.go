package service

import (
	"encoding/json"
	"net/http"

	model "github.com/MufidJamaluddin/simplewebgo/model"
)

// UserServiceRoute Login dengan JWT
func UserServiceRoute(w http.ResponseWriter, r *http.Request, claim *model.UserInfo) {
	user, err := json.Marshal(claim)

	if err != nil {
		http.Error(w, "Claims is not valid!", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(user))
}
