package main

/**
 *	@author Mufid Jamaluddin
 **/

import (
	"fmt"
	"github.com/MufidJamaluddin/simplewebgo/config"
	"github.com/MufidJamaluddin/simplewebgo/handler"
	"github.com/MufidJamaluddin/simplewebgo/module/profile"
	"github.com/MufidJamaluddin/simplewebgo/util"
	"github.com/MufidJamaluddin/simplewebgo/ui"
	"github.com/justinas/alice"
	"log"
	"net/http"
)

func main() {

	var _handler alice.Chain

	configo := config.GetConfig()

	if configo.Server.IsEnvDev() {
		_handler = alice.New(handler.DebugHandler, handler.ErrorHandler)
	} else {
		_handler = alice.New(handler.ErrorHandler)
	}

	profileModule := &util.Module{Path: "/api/profile", Handler: &_handler}
	profile.RegisterRoute(profileModule)

	http.Handle("/", _handler.ThenFunc(ui.IndexHandler))

	serverPath := fmt.Sprintf("%s:%d", configo.Server.Host, configo.Server.Port)

	fmt.Println(serverPath)

	log.Fatal(http.ListenAndServe(serverPath, nil))
}
