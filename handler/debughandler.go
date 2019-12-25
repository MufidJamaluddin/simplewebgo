package handler

import (
	"log"
	"net/http"
	"time"
)

/**
 *	LOGGING AKTIVITAS WEB
 *	DALAM PHASE DEVELOPMENT
 *	
 *	@author Mufid Jamaluddin
 *
 **/

// DebugHandler : Handler untuk Logging (Print) Setiap Aktivitas Web
func DebugHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}