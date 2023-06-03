package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/hsingyingli/yourslack-backend/db/sqlc"
	"github.com/hsingyingli/yourslack-backend/pkg/utils"
)

type CreateUserRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type CreateUserResponse struct {
	ID       int64     `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	UpdateAt time.Time `json:"updated_at"`
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

	rsp := CreateUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		UpdateAt: user.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}
