package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kozhamseitova/ustaz-hub-api/api"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
	"log"
	"net/http"
	"strconv"
)

// createProduct creates a new product
// @Summary Create product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param req body api.CreateProductRequest true "Request body"
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 401 {object} api.Error
// @Failure 500 {object} api.Error
// @Security ApiKeyAuth
// @Router /products [post]
func (h *Handler) createProduct(ctx *gin.Context) {
	action := "create"
	var req api.CreateProductRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var product = ConvertCreateProductRequestToProduct(&req)

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

	productId, err := h.services.Product.CreateProduct(ctx, product)
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
		Data:    productId,
	})

}

// getProductById gets a product by ID
// @Summary Get product by ID
// @Description Get a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 401 {object} api.Error
// @Failure 500 {object} api.Error
// @Security ApiKeyAuth
// @Router /products/{id} [get]
func (h *Handler) getProductById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	product, err := h.services.Product.GetProductById(ctx, int64(id))

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
		Data:    product,
	})
}

// getProductsByUserId gets products by user ID
// @Summary Get products by user ID
// @Description Get all products created by a specific user
// @Tags products
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 401 {object} api.Error
// @Failure 500 {object} api.Error
// @Security ApiKeyAuth
// @Router /users/{user_id}/products [get]
func (h *Handler) getProductsByUserId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	products, err := h.services.Product.GetProductsByUserId(ctx, int64(id))

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
		Data:    products,
	})
}

// getAllProducts gets all products
// @Summary Get all products
// @Description Get all products in the system
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 401 {object} api.Error
// @Failure 500 {object} api.Error
// @Security ApiKeyAuth
// @Router /products [get]
func (h *Handler) getAllProducts(ctx *gin.Context) {
	products, err := h.services.Product.GetAllProducts(ctx)

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
		Data:    products,
	})
}

// updateProduct updates a product
// @Summary Update product
// @Description Update a product
// @Tags products
// @Accept json
// @Produce json
// @Param req body api.CreateProductRequest true "Request body"
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 401 {object} api.Error
// @Failure 500 {object} api.Error
// @Security ApiKeyAuth
// @Router /products [put]
func (h *Handler) updateProduct(ctx *gin.Context) {
	action := "update"

	var req api.CreateProductRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var product = ConvertCreateProductRequestToProduct(&req)

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

	err = h.services.Product.UpdateProduct(ctx, product)
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

// deleteProduct deletes a product
// @Summary Delete product
// @Description Delete a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 401 {object} api.Error
// @Failure 500 {object} api.Error
// @Security ApiKeyAuth
// @Router /products/{id} [delete]
func (h *Handler) deleteProduct(ctx *gin.Context) {
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

	err = h.services.Product.DeleteProduct(ctx, int64(id))

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

func ConvertCreateProductRequestToProduct(req *api.CreateProductRequest) *entity.Product {
	return &entity.Product{
		ID:          req.ID,
		User:        entity.User{ID: req.UserId},
		Title:       req.Title,
		Description: req.Description,
		FilePath:    req.FilePath,
		UploadedAt:  req.UploadedAt,
	}
}
