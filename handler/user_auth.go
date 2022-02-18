package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/model"
)

func (h *Handler) authUser(c *gin.Context) {
	h.logger.Info("Working authUser")

	var input model.AuthUser
	if err := c.BindJSON(&input); err != nil {
		h.logger.Errorf("authUser: error while decoding request:%s", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input body"})
		return
	}
	validationErrors := validateStruct(input)
	if len(validationErrors) != 0 {
		h.logger.Warnf("Incorrect data came from the request:%s", validationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong email or password entered"})
		return
	}
	id, err := h.service.AppUser.AuthUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong email or password entered"})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}
}
