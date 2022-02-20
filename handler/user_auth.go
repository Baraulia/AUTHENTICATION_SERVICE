package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/model"
	"strconv"
)

// authUser godoc
// @Summary authUser
// @Description check auth information
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param input body model.AuthUser true "User"
// @Success 200 {object} auth_proto.GeneratedTokens
// @Failure 400 {string} string
// @Failure 401 {string} string
// @Router /users/login [post]
func (h *Handler) authUser(ctx *gin.Context) {
	h.logger.Info("Working authUser")
	var input model.AuthUser
	if err := ctx.BindJSON(&input); err != nil {
		h.logger.Errorf("authUser: error while decoding request:%s", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input body"})
		return
	}
	validationErrors := validateStruct(input)
	if len(validationErrors) != 0 {
		h.logger.Warnf("Incorrect data came from the request:%s", validationErrors)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Wrong email or password entered"})
		return
	}
	tokens, id, err := h.service.AppUser.AuthUser(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong email or password entered"})
	} else {
		ctx.Header("id", strconv.Itoa(id))
		ctx.JSON(http.StatusOK, tokens)
	}
}
