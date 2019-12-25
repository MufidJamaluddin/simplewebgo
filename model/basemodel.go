package model

/*
	BASE MODEL
	@author Mufid Jamaluddin <mufidjamaluddin@outlook.com>
*/

// BaseModel ...
type BaseModel struct {
	Message string
}

// BaseResponseModel ...
type BaseResponseModel struct {
	Data  []interface{}
	Total int
	Start int
	Size  int
}
