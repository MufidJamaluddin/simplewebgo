package profile

import (
	"github.com/MufidJamaluddin/simplewebgo/module/profile/service"
	"github.com/MufidJamaluddin/simplewebgo/util"
)

// RegisterRoute ...
func RegisterRoute(module *util.Module) {
	module.Register("page", service.PageServiceRoute)
	module.Register("home", service.HomeServiceRoute)
	module.Register("career", service.CareerServiceRoute)
	module.Register("portfolio", service.PortfolioServiceRoute)
}