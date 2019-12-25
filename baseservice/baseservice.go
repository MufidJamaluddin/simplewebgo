package baseservice

import (
	"net/http"
)

/*
	BASE SERVICE
	@author Mufid Jamaluddin <mufidjamaluddin@outlook.com>
*/

// BaseService ..
type BaseService struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

// Use ..
func (bs *BaseService) Use(w http.ResponseWriter, r *http.Request) {
	bs.ResponseWriter = w
	bs.Request = r
}

// GetWriter ..
func (bs *BaseService) GetWriter() http.ResponseWriter {
	return bs.ResponseWriter
}

// GetRequest ..
func (bs *BaseService) GetRequest() *http.Request {
	return bs.Request
}
