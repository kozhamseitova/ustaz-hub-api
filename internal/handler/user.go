package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kozhamseitova/ustaz-hub-api/api"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
	"log"
	"net/http"
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

	err = h.services.CreateUser(ctx, &req)
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

	token, err := h.services.Login(ctx, req.Username, req.Password)
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
	//var req entity.User
	//
	//err := ctx.ShouldBindJSON(&req)
	//if err != nil {
	//	log.Printf("bind json err: %s \n", err.Error())
	//	ctx.JSON(http.StatusBadRequest, &api.Error{
	//		Code:    http.StatusBadRequest,
	//		Message: err.Error(),
	//	})
	//	return
	//}

}

func (h *Handler) deleteUser(ctx *gin.Context) {

}
