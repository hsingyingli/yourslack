package server

import (
	"log"

	"github.com/gin-gonic/gin"
	db "github.com/hsingyingli/yourslack-backend/db/sqlc"
	"github.com/hsingyingli/yourslack-backend/pkg/token"
	"github.com/hsingyingli/yourslack-backend/pkg/utils"
)

type Server struct {
	config     utils.Config
	tokenMaker token.Maker
	store      db.Store
	router     *gin.Engine
}

func NewServer(config utils.Config, store db.Store) *Server {
	router := gin.Default()
	tokenMaker, err := token.NewPasetoMaker(config.SYMMETRICKEY)

	if err != nil {
		log.Fatal(err)
	}

	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
		store:      store,
		router:     router,
	}

	server.setupRouter()

	return server
}

func (server *Server) Start(port string) error {
	return server.router.Run(port)
}

func errorMsg(err error) gin.H {
	return gin.H{"error": err.Error()}
}
