package api

import (
	db "simple_bank_app/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requestsfor out banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP serverand setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

// starts the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
