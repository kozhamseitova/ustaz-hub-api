package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kozhamseitova/ustaz-hub-api/api"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createUser(ctx *gin.Context) {
	var req entity.User

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = h.services.User.CreateUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) loginUser(ctx *gin.Context) {
	var req api.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	token, err := h.services.User.Login(ctx, req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    http.StatusOK,
		Message: "success",
		Data:    token,
	})
}

func (h *Handler) updateUser(ctx *gin.Context) {
	action := "update"

	var req entity.User

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	userContext, ok := ctx.Get(userCtx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    http.StatusInternalServerError,
			Message: "user not found",
		})
		return
	}

	user, ok := userContext.(UserCtx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    http.StatusInternalServerError,
			Message: "user is not of the expected type",
		})
		return
	}

	isAllowed := h.services.User.CheckPermission(user.UserId, req.ID, user.UserRole, action)

	if !isAllowed {
		ctx.JSON(http.StatusUnauthorized, &api.Error{
			Code:    http.StatusUnauthorized,
			Message: "permission denied",
		})
		return
	}

	err = h.services.User.UpdateUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    0,
		Message: "successfully updated",
		Data:    user.UserId,
	})

}

func (h *Handler) deleteUser(ctx *gin.Context) {
	action := "delete"

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	userId := int64(id)

	userContext, ok := ctx.Get(userCtx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    http.StatusInternalServerError,
			Message: "user not found",
		})
		return
	}

	user, ok := userContext.(UserCtx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    http.StatusInternalServerError,
			Message: "user is not of the expected type",
		})
		return
	}

	isAllowed := h.services.User.CheckPermission(user.UserId, userId, user.UserRole, action)

	if !isAllowed {
		ctx.JSON(http.StatusUnauthorized, &api.Error{
			Code:    http.StatusUnauthorized,
			Message: "permission denied",
		})
		return
	}

	err = h.services.User.DeleteUser(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    0,
		Message: "successfully deleted",
		Data:    userId,
	})

}
