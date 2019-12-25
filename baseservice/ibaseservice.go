package baseservice

import (
	"net/http"
)

/*
	BASE SERVICE INTERFACE
	
	Base Service Fommon REST API : GET, POST, PUT, DELETE
	@author Mufid Jamaluddin <mufidjamaluddin@outlook.com> 
*/

// IBaseService ...
type IBaseService interface {
	GetWriter() http.ResponseWriter 
	GetRequest() *http.Request
	Get() (interface{}, error)
	Post() (interface{}, error)
	Put() (interface{}, error)
	Delete() (interface{}, error)
}