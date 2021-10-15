package server

import (
	"github.com/gin-gonic/gin"
	"github.com/joaomarcuslf/keyforge-name-of-the-day/handlers"
)

type Server struct {
	Port string
}

func NewServer(port string) *Server {
	return &Server{
		Port: port,
	}
}

func (a *Server) Run() {
	router := gin.New()

	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		handlers.Index(c)
	})

	router.POST("/tweet", func(c *gin.Context) {
		handlers.SendTweet(c)
	})

	router.Run(":" + a.Port)
}
