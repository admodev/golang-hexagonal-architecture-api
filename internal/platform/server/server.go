package server

import (
	report "bctec/internal"
	"bctec/internal/platform/server/handler/health"
	"bctec/internal/platform/server/handler/reports"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// deps
	reportsRepository report.ReportsRepository
}

func New(host string, port uint, reportsRepository report.ReportsRepository) Server {
	srv := Server{
		engine:            gin.New(),
		httpAddr:          fmt.Sprintf("%s:%d", host, port),
		reportsRepository: reportsRepository,
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
}
