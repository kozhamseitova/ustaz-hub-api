package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kozhamseitova/ustaz-hub-api/api"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
	"log"
	"net/http"
	"strconv"
)

// createPost creates a new post
// @Summary Create post
// @Description Create a new post
// @Tags posts
// @Accept json
// @Produce json
// @Param req body api.CreatePostRequest true "Request body"
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 401 {object} api.Error
// @Failure 500 {object} api.Error
// @Security ApiKeyAuth
// @Router /posts [post]
func (h *Handler) createPost(ctx *gin.Context) {
	action := "create"

	var req api.CreatePostRequest

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

	isAllowed := h.services.User.CheckPermission(user.UserId, req.UserId, user.UserRole, action)

	if !isAllowed {
		ctx.JSON(http.StatusUnauthorized, &api.Error{
			Code:    http.StatusUnauthorized,
			Message: "permission denied",
		})
		return
	}

	var post = ConvertCreatePostRequestToPost(&req)

	postId, err := h.services.Post.CreatePost(ctx, post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    http.StatusOK,
		Message: "successfully created",
		Data:    postId,
	})

}

// getPostById gets a post by ID
// @Summary Get post by ID
// @Description Get a post by its ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 401 {object} api.Error
// @Failure 500 {object} api.Error
// @Security ApiKeyAuth
// @Router /posts/{id} [get]
func (h *Handler) getPostById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	post, err := h.services.Post.GetPostById(ctx, int64(id))

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
		Data:    post,
	})
}

// getPostsByUserId gets posts by user ID
// @Summary Get posts by user ID
// @Description Get all posts created by a specific user
// @Tags posts
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 401 {object} api.Error
// @Failure 500 {object} api.Error
// @Security ApiKeyAuth
// @Router /users/{user_id}/posts [get]
func (h *Handler) getPostsByUserId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	posts, err := h.services.Post.GetPostsByUserId(ctx, int64(id))

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
		Data:    posts,
	})
}

// getAllPosts gets all posts
// @Summary Get all posts
// @Description Get all posts in the system
// @Tags posts
// @Accept json
// @Produce json
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 401 {object} api.Error
// @Failure 500 {object} api.Error
// @Security ApiKeyAuth
// @Router /posts [get]
func (h *Handler) getAllPosts(ctx *gin.Context) {
	posts, err := h.services.Post.GetAllPosts(ctx)

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
		Data:    posts,
	})
}

// updatePost updates a post
// @Summary Update post
// @Description Update a post
// @Tags posts
// @Accept json
// @Produce json
// @Param req body api.CreatePostRequest true "Request body"
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 401 {object} api.Error
// @Failure 500 {object} api.Error
// @Security ApiKeyAuth
// @Router /posts [put]
func (h *Handler) updatePost(ctx *gin.Context) {
	action := "update"
	var req api.CreatePostRequest

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

	isAllowed := h.services.User.CheckPermission(user.UserId, req.UserId, user.UserRole, action)

	if !isAllowed {
		ctx.JSON(http.StatusUnauthorized, &api.Error{
			Code:    http.StatusUnauthorized,
			Message: "permission denied",
		})
		return
	}

	var post = ConvertCreatePostRequestToPost(&req)

	err = h.services.Post.UpdatePost(ctx, post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    http.StatusOK,
		Message: "successfully updated",
	})

}

// deletePost deletes a post
// @Summary Delete post
// @Description Delete a post
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 401 {object} api.Error
// @Failure 500 {object} api.Error
// @Security ApiKeyAuth
// @Router /posts/{id} [delete]
func (h *Handler) deletePost(ctx *gin.Context) {
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

	err = h.services.Post.DeletePost(ctx, int64(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    http.StatusOK,
		Message: "successfully deleted",
	})
}

func ConvertCreatePostRequestToPost(req *api.CreatePostRequest) *entity.Post {
	return &entity.Post{
		ID:       req.ID,
		User:     entity.User{ID: req.UserId},
		PostText: req.PostText,
	}
}
