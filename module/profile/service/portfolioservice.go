package service

import (
	"github.com/MufidJamaluddin/simplewebgo/baseservice"
	basemodel "github.com/MufidJamaluddin/simplewebgo/model"
	"github.com/MufidJamaluddin/simplewebgo/module/profile/model"
	"net/http"
)

type PortfolioService struct {
	baseservice.BaseService
}

func (service *PortfolioService) Get() (interface{}, error) {

	listabout := []model.AboutItemModel{}

	about := model.AboutItemModel{
		ID:      1,
		Title:   "GET Portfolio!",
		Content: "Bebas!",
	}

	data := append(listabout, about)

	response := &model.AboutResponseModel{
		Data:  data,
		Total: 1,
	}

	return response, nil
}

func (service *PortfolioService) Post() (interface{}, error) {

	response := &basemodel.BaseModel{
		Message: "POST Portfolio!",
	}

	return response, nil
}

func (service *PortfolioService) Put() (interface{}, error) {
	response := &basemodel.BaseModel{
		Message: "PUT Portfolio!",
	}

	return response, nil
}

func (service *PortfolioService) Delete() (interface{}, error) {
	response := &basemodel.BaseModel{
		Message: "DELETE Portfolio!",
	}

	return response, nil
}

func PortfolioServiceRoute(w http.ResponseWriter, r *http.Request) {
	service := PortfolioService{}
	service.Use(w, r)

	baseservice.RouteService(&service)
}
