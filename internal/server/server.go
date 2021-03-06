package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/sqlite-go/internal/handlers"
)

type Server struct {
	e   *gin.Engine
	h   *handlers.Handler
	srv http.Server
}

func New(handler *handlers.Handler) *Server {
	return &Server{
		h: handler,
		e: gin.Default(),
	}
}

func (s *Server) SetupRoutes() *gin.Engine {
	s.e.POST("/createUser", s.h.CreateUser)
	s.e.GET("/readUser/:id", s.h.ReadUser)
	s.e.GET("/readUsers", s.h.ReadUsers)
	return s.e
}

func (s *Server) Start() {

	s.srv = http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: s.SetupRoutes(),
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		if err := s.srv.Close(); err != nil {
			log.Println("failed to shutdown server", err)
		}
	}()

	if err := s.srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("server closed after interruption")
		} else {
			log.Println("unexpected server shutdown. err:", err)
		}
	}
}

func (s *Server) Stop() error {
	return s.srv.Close()
}
