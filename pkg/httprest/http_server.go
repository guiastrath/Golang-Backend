package httprest

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type HttpServer interface {
	Use(middleware ...Middleware) *httpServer
	UseAuth(authMiddleware Middleware) *httpServer
	UseCors(authMiddleware Middleware) *httpServer
	AddHandlers(handlers ...Handler) *httpServer
	ListenAndServe()

	applyCors(handler http.Handler, corsConfig *CorsConfig) http.Handler
	applyAuth(handler http.Handler) http.Handler
	applyMiddlewares(handler http.Handler) http.Handler
	getServer() *http.ServeMux
	getConfig() *WSConfig
	getAuth() Middleware
}

type Handler interface {
	Handlers() []*Route
}

type WSConfig struct {
	ServerPort string
}

type httpServer struct {
	server      *http.ServeMux
	config      *WSConfig
	auth        Middleware
	cors        Middleware
	middlewares []Middleware
	routes      []*Route
}

func NewWebService(config *WSConfig) HttpServer {
	ws := &httpServer{
		server: http.NewServeMux(),
		config: config,
	}

	return ws
}

func (ws *httpServer) Use(middleware ...Middleware) *httpServer {
	ws.middlewares = append(ws.middlewares, middleware...)
	return ws
}

func (ws *httpServer) UseAuth(authMiddleware Middleware) *httpServer {
	ws.auth = authMiddleware
	return ws
}

func (ws *httpServer) UseCors(corsMiddleware Middleware) *httpServer {
	ws.cors = corsMiddleware
	return ws
}

func (ws *httpServer) AddHandlers(handlers ...Handler) *httpServer {
	for _, handler := range handlers {
		routes := handler.Handlers()

		for _, route := range routes {
			routeHandler := http.Handler(http.HandlerFunc(route.handler))

			if route.corsConfig != nil {
				routeHandler = ws.applyCors(routeHandler, route.corsConfig)
			}

			if route.authentication && ws.auth != nil {
				routeHandler = ws.applyAuth(routeHandler)
			}

			routeHandler = ws.applyMiddlewares(routeHandler)

			ws.routes = append(ws.routes, route)
			ws.server.Handle(route.path, routeHandler)
		}
	}

	return ws
}

func (ws *httpServer) applyCors(handler http.Handler, corsConfig *CorsConfig) http.Handler {
	return ws.cors(handler)
}

func (ws *httpServer) applyAuth(handler http.Handler) http.Handler {
	return ws.auth(handler)
}

func (ws *httpServer) applyMiddlewares(handler http.Handler) http.Handler {
	for i := len(ws.middlewares) - 1; i >= 0; i-- {
		handler = ws.middlewares[i](handler)
	}
	return handler
}

func (ws *httpServer) getConfig() *WSConfig {
	return ws.config
}

func (ws *httpServer) getServer() *http.ServeMux {
	return ws.server
}

func (ws *httpServer) getAuth() Middleware {
	return ws.auth
}

func (ws *httpServer) ListenAndServe() {
	log.Printf("HTTP service started at port %s", ws.config.ServerPort)
	go func() {
		if err := http.ListenAndServe(ws.config.ServerPort, GlobalCorsMiddlewares(ws.server)); err != nil {
			log.Fatalf("error on starting service: %v", err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, os.Interrupt)
	<-interrupt

	log.Println("HTTP service stopped")
}
