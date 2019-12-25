package util

/**
 *	Route Berdasarkan Module
 **/

import (
	"fmt"
	"net/http"
	"github.com/justinas/alice"
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
