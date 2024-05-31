package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)

  //oAUTH
	r.GET("/auth/:provider/callback", s.authCallBackHandler)
	r.GET("/auth/logout/:provider", s.authLogoutHandler)
  r.GET("/auth/:provider", s.authHandler)

	return r
}
