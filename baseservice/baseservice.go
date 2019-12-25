package baseservice

import (
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
)

/*
	BASE SERVICE
	@author Mufid Jamaluddin <mufidjamaluddin@outlook.com>
*/

// BaseService ..
type BaseService struct {
	ResponseWriter   http.ResponseWriter
	Request          *http.Request
	RequestBodyModel interface{}
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

// GetRequestBody ...
func (bs *BaseService) GetRequestBody() interface{} {
	if bs.RequestBodyModel == nil {
		return nil
	}

	modelType := reflect.TypeOf((*bs).RequestBodyModel)

	modelValue := reflect.New(modelType).Interface()
	err := json.NewDecoder(bs.GetRequest().Body).Decode(modelValue)

	if err != nil {
		panic(err)
	}

	return modelValue
}

// GetURIParam ...
func (bs *BaseService) GetURIParam() url.Values {
	return bs.GetRequest().URL.Query()
}
