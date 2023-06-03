package server

import (
	"github.com/gin-gonic/gin"
	db "github.com/hsingyingli/yourslack-backend/db/sqlc"
)

type Server struct {
	router *gin.Engine
	store  db.Store
}

func NewServer(store db.Store) *Server {
	router := gin.Default()

	server := &Server{
		router: router,
		store:  store,
	}

	server.setUpRouter()

	return server
}

func (server *Server) Start(port string) error {
	return server.router.Run(port)
}

func errorMsg(err error) gin.H {
	return gin.H{"error": err.Error()}
}
