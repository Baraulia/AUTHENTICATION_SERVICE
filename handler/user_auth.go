package handler

import (
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) authUser(c *gin.Context) {
	h.logger.Info("Working authUser")

	if c.Request.Method == "OPTIONS" {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Status(200)
		return
	}
	var input model.AuthUser
	if err := c.BindJSON(&input); err != nil {
		h.logger.Errorf("authUser: error while decoding request:%s", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	validationErrors := validateStruct(input)
	if len(validationErrors) != 0 {
		h.logger.Warnf("Incorrect data came from the request:%s", validationErrors)
		c.JSON(http.StatusBadRequest, validationErrors)
		return
	}
	id, err := h.service.AppUser.AuthUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "wrong email or password entered"})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}
}
