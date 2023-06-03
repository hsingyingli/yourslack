package server

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type renewAccessTokenResponse struct {
	ID          int64     `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	AccessToken string    `json:"access_token"`
	ExpiredAt   time.Time `json:"expired_at"`
}

func (server *Server) renewAccessToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorMsg(err))
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(refreshToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorMsg(err))
		return
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorMsg(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	if session.RefreshToken != refreshToken {
		err := errors.New("mismatch refresh token")
		ctx.JSON(http.StatusUnauthorized, errorMsg(err))
	}

	user, err := server.store.GetUser(ctx, session.UID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(user, server.config.ACCESS_TOKEN_DURATION)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	rsp := renewAccessTokenResponse{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		AccessToken: accessToken,
		ExpiredAt:   accessPayload.ExpiredAt,
	}

	ctx.JSON(http.StatusAccepted, rsp)
}
