package util

/**
 *	Route Berdasarkan Module
 **/

import (
	"fmt"
	security "github.com/MufidJamaluddin/simplewebgo/security"
	"github.com/justinas/alice"
	"net/http"
)

// Module : Module Program
type Module struct {
	Path    string
	Handler *alice.Chain
}

// Register Register Module Route Handler
func (module *Module) Register(subpath string, routeHandler http.HandlerFunc) {
	path := fmt.Sprintf("%s/%s", (*module).Path, subpath)
	http.Handle(path, (*module).Handler.ThenFunc(routeHandler))
}

// RegisterSecure Register Module Route Handler yang dilindungi oleh JWT
func (module *Module) RegisterSecure(subpath string, routeHandler security.SecureHandleFunc) {
	path := fmt.Sprintf("%s/%s", (*module).Path, subpath)
	http.Handle(path, (*module).Handler.Then(security.HandlerSecurity(routeHandler)))
}
