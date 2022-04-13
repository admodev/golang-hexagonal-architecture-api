package server

import (
	"bctec/internal"
	"bctec/internal/platform/server/handler/health"
	"bctec/internal/platform/server/handler/reports"
	"bctec/internal/platform/server/handler/users"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// deps
	reportsRepository report.ReportsRepository
	usersRepository   report.UserRepository
}

func New(host string, port uint, reportsRepository report.ReportsRepository, usersRepository report.UserRepository) Server {
	srv := Server{
		engine:            gin.New(),
		httpAddr:          fmt.Sprintf("%s:%d", host, port),
		reportsRepository: reportsRepository,
		usersRepository:   usersRepository,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server is running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.HealthCheckHandler())
	s.engine.POST("/reports", reports.CreateHandler(s.reportsRepository))
	s.engine.POST("/users", users.CreateHandler(s.usersRepository))
}
