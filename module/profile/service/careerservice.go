package service

import (
	"github.com/MufidJamaluddin/simplewebgo/baseservice"
	basemodel "github.com/MufidJamaluddin/simplewebgo/model"
	"github.com/MufidJamaluddin/simplewebgo/module/profile/model"
	"net/http"
)

type CareerService struct {
	baseservice.BaseService
}

func (service *CareerService) Get() (interface{}, error) {

	listabout := []model.AboutItemModel{}

	about := model.AboutItemModel{
		ID:      1,
		Title:   "GET Career!",
		Content: "Bebas!",
	}

	data := append(listabout, about)

	response := &model.AboutResponseModel{
		Data:  data,
		Total: 1,
	}

	return response, nil
}

func (service *CareerService) Post() (interface{}, error) {

	response := &basemodel.BaseModel{
		Message: "POST Career!",
	}

	return response, nil
}

func (service *CareerService) Put() (interface{}, error) {
	response := &basemodel.BaseModel{
		Message: "PUT Career!",
	}

	return response, nil
}

func (service *CareerService) Delete() (interface{}, error) {
	response := &basemodel.BaseModel{
		Message: "DELETE Career!",
	}

	return response, nil
}

func CareerServiceRoute(w http.ResponseWriter, r *http.Request) {
	service := CareerService{}
	service.Use(w, r)

	baseservice.RouteService(&service)
}
