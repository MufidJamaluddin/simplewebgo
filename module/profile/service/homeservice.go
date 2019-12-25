package service

import (
	"github.com/MufidJamaluddin/simplewebgo/baseservice"
	basemodel "github.com/MufidJamaluddin/simplewebgo/model"
	"github.com/MufidJamaluddin/simplewebgo/module/profile/model"
	"net/http"
)

type HomeService struct {
	baseservice.BaseService
}

func (service *HomeService) Get() (interface{}, error) {

	listabout := []model.AboutItemModel{}

	about := model.AboutItemModel{
		ID:      1,
		Title:   "GET halaman home!",
		Content: "Bebas!",
	}

	data := append(listabout, about)

	response := &model.AboutResponseModel{
		Data:  data,
		Total: 1,
	}

	return response, nil
}

func (service *HomeService) Post() (interface{}, error) {

	response := &basemodel.BaseModel{
		Message: "POST halaman home!",
	}

	return response, nil
}

func (service *HomeService) Put() (interface{}, error) {
	response := &basemodel.BaseModel{
		Message: "PUT halaman home!",
	}

	return response, nil
}

func (service *HomeService) Delete() (interface{}, error) {
	response := &basemodel.BaseModel{
		Message: "DELETE halaman home!",
	}

	return response, nil
}

func HomeServiceRoute(w http.ResponseWriter, r *http.Request) {
	service := HomeService{}
	service.Use(w, r)

	baseservice.RouteService(&service)
}
