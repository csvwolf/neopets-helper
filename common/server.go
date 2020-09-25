package common

import (
	"fmt"
	"log"
	"net/http"
)

/**
Middleware is the hooks for controller
*/
type Middleware func(http.HandlerFunc) http.HandlerFunc

type WebServer struct {
	Host string
	Port int
}

/**
Define the method and default handler
*/
type Route struct {
	Method  string
	Handler http.HandlerFunc
}

var routes = make(map[string][]Route)         // pattern-route list
var globalMiddlewares = make([]Middleware, 0) // middleware list

/**
handle with http method and return status method not allowed if status is not in routes
*/
func httpMethodMiddleware(routes []Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, route := range routes {
			if r.Method == route.Method {
				route.Handler.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

/**
Set Routes to the http server
*/
func setRoutes() {
	for pattern, routeList := range routes {
		http.HandleFunc(pattern, ChainMiddleware(httpMethodMiddleware(routeList), globalMiddlewares...))
	}
}

/**
Add router to pattern-route list
*/
func addRoute(pattern string, method string, handler http.HandlerFunc) {
	if _, ok := routes[pattern]; ok {
		routes[pattern] = append(routes[pattern], Route{Method: method, Handler: handler})
	} else {
		routes[pattern] = []Route{{Method: method, Handler: handler}}
	}
}

/**
Flatten the http.HandlerFunc Functions
*/
func ChainMiddleware(basicHandler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	var handler = basicHandler

	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

/**
Use Current Routes and Start the server by WebServer Object Config
*/
func (server *WebServer) Start(handler http.Handler) {
	setRoutes()
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", server.Host, server.Port), handler)
	log.Fatal(err)
}

/**
Set Global Middleware
*/
func (server *WebServer) Use(middleware Middleware) {
	globalMiddlewares = append(globalMiddlewares, middleware)
}

/**
Set Get Method Handler to Specific Pattern
*/
func (server *WebServer) Get(pattern string, handler http.HandlerFunc) {
	addRoute(pattern, http.MethodGet, handler)
}

/**
Set Post Method Handler to Specific Pattern
*/
func (server *WebServer) Post(pattern string, handler http.HandlerFunc) {
	addRoute(pattern, http.MethodPost, handler)
}

/**
Set Put Method Handler to Specific Pattern
*/
func (server *WebServer) Put(pattern string, handler http.HandlerFunc) {
	addRoute(pattern, http.MethodPut, handler)
}

/**
Set Patch Method Handler to Specific Pattern
*/
func (server *WebServer) Patch(pattern string, handler http.HandlerFunc) {
	addRoute(pattern, http.MethodPatch, handler)
}

/**
Set Delete Method Handler to Specific Pattern
*/
func (server *WebServer) Delete(pattern string, handler http.HandlerFunc) {
	addRoute(pattern, http.MethodDelete, handler)
}
