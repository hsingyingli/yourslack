package server

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/hsingyingli/yourslack-backend/db/sqlc"
	"github.com/hsingyingli/yourslack-backend/pkg/token"
	"github.com/hsingyingli/yourslack-backend/pkg/utils"
)

type CreateUserRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type CreateUserResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
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

	args := db.CreateUserParams{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	rsp := CreateUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	ctx.JSON(http.StatusOK, rsp)
}

type GetUserResponse struct {
	User db.User
}

func (server *Server) getUser(ctx *gin.Context) {
	payload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	user, err := server.store.GetUser(ctx, int64(payload.UID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	rsp := GetUserResponse{
		User: user,
	}
	ctx.JSON(http.StatusOK, rsp)
	return
}

func (server *Server) deleteUser(ctx *gin.Context) {
	payload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	err := server.store.DeleteUser(ctx, payload.UID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user has been deleted"})
}

type UpdateUserRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UpdateUserResponse struct {
	User db.User
}

func (server *Server) updateUser(ctx *gin.Context) {
	payload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	var req UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	args := db.UpdateUserParams{
		ID:        payload.UID,
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		UpdatedAt: time.Now(),
	}

	updatedUser, err := server.store.UpdateUser(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	rsp := UpdateUserResponse{
		User: updatedUser,
	}

	ctx.JSON(http.StatusAccepted, rsp)
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	ACCESS_TOKEN string    `json:"access_token"`
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	ExpiredAt    time.Time `json:"expired_at"`
}

func (server *Server) login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	user, err := server.store.GetUserByEmail(ctx, req.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorMsg(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	err = utils.CheckPassword(user.Password, req.Password)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorMsg(err))
		return
	}

	accessToken, accessTokenPayload, err := server.tokenMaker.CreateToken(user, server.config.ACCESS_TOKEN_DURATION)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	refreshToken, refreshTokenPayload, err := server.tokenMaker.CreateToken(user, server.config.REFRESH_TOKEN_DURATION)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	_, err = server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshTokenPayload.ID,
		UID:          user.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		ExpiredAt:    refreshTokenPayload.ExpiredAt,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	rsp := LoginResponse{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		ACCESS_TOKEN: accessToken,
		ExpiredAt:    accessTokenPayload.ExpiredAt,
	}

	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie("refresh_token", refreshToken, server.config.COOKIE_MAXAGE, "/", "localhost", false, true)
	ctx.JSON(http.StatusAccepted, rsp)
}

func (server *Server) logout(ctx *gin.Context) {
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
