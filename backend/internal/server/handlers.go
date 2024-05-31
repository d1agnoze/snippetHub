package server

import (
	"context"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

// AUTH
const (
	authRedirectURL = "http://localhost:5173/"
)

func (s *Server) authCallBackHandler(c *gin.Context) {
	provider := c.Param("provider")
	c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info(user.Name)
	c.Redirect(http.StatusFound, "http://localhost:5173/")
}

func (s *Server) authLogoutHandler(c *gin.Context) {
	gothic.Logout(c.Writer, c.Request)
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173/")
}

func (s *Server) authHandler(c *gin.Context) {
	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		log.Info(gothUser)
		c.JSON(http.StatusOK, gothUser)
	} else {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}

//================================
