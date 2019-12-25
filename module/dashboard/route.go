package dashboard

import (
	"github.com/MufidJamaluddin/simplewebgo/module/dashboard/service"
	"github.com/MufidJamaluddin/simplewebgo/util"
)

// RegisterRoute ...
func RegisterRoute(module *util.Module) {
	module.Register("login", service.LoginServiceRoute)
	module.RegisterSecure("user", service.UserServiceRoute)
}
