package httprest

import (
	"fmt"
	"net/http"
)

type (
	Middleware  func(http.Handler) http.Handler
	handlerFunc func(http.ResponseWriter, *http.Request)
)

type CorsConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	MaxAge           int
	AllowCredentials bool
}

type Route struct {
	method         string
	path           string
	handler        handlerFunc
	authentication bool
	corsConfig     *CorsConfig
}

type RouteBuilder struct {
	prefix         string
	method         string
	rootPath       string
	handler        handlerFunc
	authentication bool
	corsConfig     *CorsConfig
}

func (rb *RouteBuilder) Build() *Route {
	path := fmt.Sprintf("%s %s%s", rb.method, rb.prefix, rb.rootPath)

	route := &Route{
		method:         rb.method,
		path:           path,
		handler:        rb.handler,
		authentication: rb.authentication,
		corsConfig:     rb.corsConfig,
	}

	return route
}

func (rb *RouteBuilder) To(function handlerFunc) *RouteBuilder {
	rb.handler = function
	return rb
}

func (rb *RouteBuilder) Cors(config *CorsConfig) *RouteBuilder {
	rb.corsConfig = config
	return rb
}

func GET(rootPath string) *RouteBuilder {
	return &RouteBuilder{
		method:   http.MethodGet,
		rootPath: rootPath,
	}
}

func POST(rootPath string) *RouteBuilder {
	return &RouteBuilder{
		method:   http.MethodPost,
		rootPath: rootPath,
	}
}

func PUT(rootPath string) *RouteBuilder {
	return &RouteBuilder{
		method:   http.MethodPut,
		rootPath: rootPath,
	}
}

func DELETE(rootPath string) *RouteBuilder {
	return &RouteBuilder{
		method:   http.MethodDelete,
		rootPath: rootPath,
	}
}

func PATCH(rootPath string) *RouteBuilder {
	return &RouteBuilder{
		method:   http.MethodPatch,
		rootPath: rootPath,
	}
}

func PrivateRoutes(prefix string, builders ...*RouteBuilder) []*Route {
	return addRoute(true, prefix, builders...)
}

func PublicRoutes(prefix string, builders ...*RouteBuilder) []*Route {
	return addRoute(false, prefix, builders...)
}

func addRoute(isPrivateRoute bool, prefix string, builders ...*RouteBuilder) []*Route {
	var routes []*Route
	for _, builder := range builders {
		builder.authentication = isPrivateRoute
		builder.prefix = prefix
		routes = append(routes, builder.Build())
	}

	return routes
}
