package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (server *Server) setupRouter() {
	server.router.GET("/health", health)

	server.router.Use(GinMiddleware(server.config.ALLOW_ORIGIN))
	v1 := server.router.Group("/v1")
	{
		user := v1.Group("/user")
		{
			// create user
			user.POST("", server.createUser)
			// login user
			user.POST("/login", server.login)
			// logout user
			user.POST("/logout", server.logout)
			//renew access token
			user.POST("/renew_access", server.renewAccessToken)
		}

		// TODO: auth
		authorized := v1.Group("/auth")
		authorized.Use(authMiddleware(server.tokenMaker))
		{
			// get user
			authorized.GET("/user/me", server.getUser)
			// update user
			authorized.PATCH("/user/me", server.updateUser)

			// delete user
			authorized.DELETE("/user/me", server.deleteUser)
		}
	}
}
