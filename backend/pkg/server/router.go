package server

func (server *Server) setUpRouter() {
	v1 := server.router.Group("v1")
	{
		v1.GET("/user")
		v1.POST("/user", server.createUser)
		v1.PATCH("/user")
		v1.DELETE("/user")
	}
}
