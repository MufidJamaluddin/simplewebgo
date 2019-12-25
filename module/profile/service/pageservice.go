package service

import (
	"github.com/MufidJamaluddin/simplewebgo/baseservice"
	basemodel "github.com/MufidJamaluddin/simplewebgo/model"
	"github.com/MufidJamaluddin/simplewebgo/module/profile/model"
	"net/http"
)

type PageService struct {
	baseservice.BaseService
}

func (service *PageService) Get() (interface{}, error) {

	listabout := []model.AboutItemModel{}

	about := model.AboutItemModel{
		ID:      1,
		Title:   "GET Halaman Page!",
		Content: "Bebas!",
	}

	data := append(listabout, about)

	response := &model.AboutResponseModel{
		Data:  data,
		Total: 1,
	}

	return response, nil
}

func (service *PageService) Post() (interface{}, error) {

	response := &basemodel.BaseModel{
		Message: "POST Halaman Page!",
	}

	return response, nil
}

func (service *PageService) Put() (interface{}, error) {
	response := &basemodel.BaseModel{
		Message: "PUT halaman page!",
	}

	return response, nil
}

func (service *PageService) Delete() (interface{}, error) {
	response := &basemodel.BaseModel{
		Message: "DELETE halaman page!",
	}

	return response, nil
}

func PageServiceRoute(w http.ResponseWriter, r *http.Request) {
	service := PageService{}
	service.Use(w, r)

	baseservice.RouteService(&service)
}
