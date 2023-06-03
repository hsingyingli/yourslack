package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/hsingyingli/yourslack-backend/db/sqlc"
	"github.com/hsingyingli/yourslack-backend/pkg/utils"
)

type CreateUserRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var request CreateUserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	user, err := server.store.CreateUser(ctx, db.CreateUserParams{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
