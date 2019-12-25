package constant

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// ApplicationName Nama Aplikasi
var ApplicationName = "Simple GO Web App"

// LoginExpirationDuration durasi login expired
var LoginExpirationDuration = time.Duration(2) * time.Hour

// JWTSigningMethod Method buat JWT
var JWTSigningMethod = jwt.SigningMethodHS256

// JWTSignatureKey kode rahasia
var JWTSignatureKey = []byte("kode rahasia yach...")
