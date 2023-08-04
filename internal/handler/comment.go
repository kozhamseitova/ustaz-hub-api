package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kozhamseitova/ustaz-hub-api/api"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
	"log"
	"net/http"
	"strconv"
)

// createComment creates a new comment
// @Summary Create comment
// @Description Create a new comment
// @Tags comments
// @Accept json
// @Produce json
// @Param req body entity.CreateComment true "Request body"
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 401 {object} api.Error
// @Failure 500 {object} api.Error
// @Security ApiKeyAuth
// @Router /comments [post]
func (h *Handler) createComment(ctx *gin.Context) {
	var req entity.CreateComment

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	commentId, err := h.services.Comment.CreateComment(ctx, &req)
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
		Data:    commentId,
	})
}

// getAllProductsComments gets all comments for a product
// @Summary Get all comments for a product
// @Description Get all comments associated with a specific product
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 401 {object} api.Error
// @Failure 500 {object} api.Error
// @Security ApiKeyAuth
// @Router /products/{id}/comments [get]
func (h *Handler) getAllProductsComments(ctx *gin.Context) {
	productId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	comments, err := h.services.Comment.GetAllProductsComments(ctx, int64(productId))

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
		Data:    comments,
	})
}

// getAllPostsComments gets all comments for a post
// @Summary Get all comments for a post
// @Description Get all comments associated with a specific post
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 401 {object} api.Error
// @Failure 500 {object} api.Error
// @Security ApiKeyAuth
// @Router /posts/{id}/comments [get]
func (h *Handler) getAllPostsComments(ctx *gin.Context) {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	comments, err := h.services.Comment.GetAllPostsComments(ctx, int64(postId))

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
		Data:    comments,
	})
}

// getCommentsByParentId gets comments by parent ID
// @Summary Get comments by parent ID
// @Description Get all comments with a specific parent ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Parent ID"
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 401 {object} api.Error
// @Failure 500 {object} api.Error
// @Security ApiKeyAuth
// @Router /comments/{id}/comments [get]
func (h *Handler) getCommentsByParentId(ctx *gin.Context) {
	parentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	comments, err := h.services.Comment.GetCommentsByParentId(ctx, int64(parentId))

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
		Data:    comments,
	})
}
