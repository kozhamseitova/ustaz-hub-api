package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kozhamseitova/ustaz-hub-api/api"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
	"log"
	"net/http"
	"strconv"
)

// createUser registration new user
//
//	@Summary      Create user
//	@Description  Create new user
//	@Tags         auth
//	@Accept       json
//	@Produce      json
//	@Param req body entity.User true "req body"
//
//	@Success      201
//	@Failure      400  {object}  api.Error
//	@Failure      500  {object}  api.Error
//	@Router       /auth/user-register [post]
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

// loginUser authentication for user
//
// @Summary      User login
// @Description  Authenticate user and generate JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param req body api.LoginRequest true "req body"
//
// @Success      200  {object}  api.Ok
// @Failure      400  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /auth/user-login [post]
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

// updateUser updates user profile
//
// @Summary      Update user profile
// @Description  Update user information
// @Tags         users
// @Accept       json
// @Produce      json
// @Param req body entity.User true "req body"
//
// @Security     ApiKeyAuth
// @Success      200  {object}  api.Ok
// @Failure      400  {object}  api.Error
// @Failure      401  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /users [put]
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

// deleteUser deletes a user
//
// @Summary      Delete user
// @Description  Delete a user account
// @Tags         users
// @Accept       json
// @Produce      json
// @Param id path int true "User ID" Format(int64)
//
// @Security     ApiKeyAuth
// @Success      200  {object}  api.Ok
// @Failure      400  {object}  api.Error
// @Failure      401  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /users/{id} [delete]
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
