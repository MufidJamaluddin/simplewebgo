package handler

import (
	"net/http"
	"time"
	"fmt"
	util "github.com/MufidJamaluddin/simplewebgo/util"
)

/**
 *	LOGGING AKTIVITAS WB
 *	JIKA TERDAPAT ERROR
 *
 *	@author Mufid Jamaluddin
 *
 **/

// ErrorHandler : Lo Setiap Unexpected Error
//					   dan Recover Web
func ErrorHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				clientIP := &r.RemoteAddr
				time := time.Now()
				url := r.URL.String()

				filehandler := &util.FileHandler{Filename: "APP_LOG.txt"}

				filehandler.WriteToFile(fmt.Sprintf("%v \t %s \t %s \t %+v\n", time, *clientIP, url, err))
				
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
