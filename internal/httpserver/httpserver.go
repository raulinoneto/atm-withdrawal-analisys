package httpserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/raulinoneto/atm-withdrawal-analisys/tools/logger"
)

type (
	Handler        func(http.ResponseWriter, *http.Request) error
	MiddlewareFunc func(Handler) Handler
	Options        struct {
		Middlewares []MiddlewareFunc
		Routes      []Route
		Port        string
		Host        string
		Logger      *logger.Logger
	}
)

type Server struct {
	middlewares   []MiddlewareFunc
	routes        []Route
	mux           *http.ServeMux
	server        *http.Server
	logger        *logger.Logger
	serverRunning bool
}

func New(opt *Options) *Server {
	if opt.Logger == nil {
		opt.Logger = logger.New(context.Background())
	}
	return &Server{
		middlewares: opt.Middlewares,
		routes:      opt.Routes,
		server: &http.Server{
			Addr: opt.Host + ":" + opt.Port,
		},
		logger:        opt.Logger,
		serverRunning: false,
	}
}

func (s *Server) quitSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()
	s.logger.Fatal("Waiting for the application to shutdown gracefully")
	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Fatal(err)
	}
	s.logger.Fatal("Application shutdown")
}

func (s *Server) configureRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range s.routes {
		s.setHandler(mux, route)
	}
	return mux
}

func (s *Server) setHandler(mux *http.ServeMux, route Route) {
	handler := route.Handler
	for _, middleware := range s.middlewares {
		handler = middleware(handler)
	}
	for _, middleware := range route.Middlewares {
		handler = middleware(handler)
	}
	if route.methods == nil {
		route.methods = make(map[string]Handler)
	}
	route.methods[route.Method] = route.Handler
	mux.HandleFunc(route.Path, route.verifyHttpMethod())
}

func (s *Server) serve() {
	if err := s.server.ListenAndServe(); err != nil {
		s.logger.Error("Server could not start: " + err.Error())
	}
	s.serverRunning = true
}

func (s *Server) Run() {
	s.server.Handler = s.configureRoutes()
	go s.serve()
	s.quitSignal()
}

func (s *Server) IsRunning() bool {
	return s.serverRunning
}

type Route struct {
	Path        string
	Method      string
	Handler     Handler
	Middlewares []MiddlewareFunc
	methods     map[string]Handler
}

func (route *Route) verifyHttpMethod() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		method, ok := route.methods[req.Method]
		if !ok {
			res.WriteHeader(http.StatusMethodNotAllowed)
			_, err := fmt.Fprint(res, "Method not allowed")
			if err != nil {
				log.Println(err.Error())
			}
			return
		}

		if err := method(res, req); err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, err = fmt.Fprint(res, "Internal Server Error: "+err.Error())
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}
