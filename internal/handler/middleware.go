package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kozhamseitova/ustaz-hub-api/api"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userCtx"
)

type UserCtx struct {
	UserId   int64
	UserRole string
}

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == " " {
		err := errors.New("authorization header is not set")
		c.AbortWithStatusJSON(http.StatusUnauthorized, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		err := errors.New("authorization header incorrect format")
		c.AbortWithStatusJSON(http.StatusUnauthorized, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	userId, userRole, err := h.services.VerifyToken(headerParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, &api.Error{
			Code:    -3,
			Message: err.Error(),
		})
		return
	}

	user := UserCtx{
		UserId:   userId,
		UserRole: userRole,
	}

	c.Set(userCtx, user)
}
