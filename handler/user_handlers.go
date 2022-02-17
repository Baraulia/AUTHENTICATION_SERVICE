package handler

import (
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// getUserByID godoc
// @Summary show master user by id
// @Description get string by ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {string} string
// @Failure 404 {object} model.User
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /user/{id} [get]
func (h *Handler) getUser(c *gin.Context) {
	paramID := c.Param("id")
	varID, err := strconv.Atoi(paramID)
	if err != nil {
		h.logger.Warnf("Handler getUser (reading param):%s", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	user, err := h.service.AppUser.GetUser(varID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// getUsers godoc
// @Summary show list master user
// @Description get users
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {array} model.User
// @Failure 400 {string} string
// @Failure 404 {object} model.User
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /user/ [get]
func (h *Handler) getUsers(c *gin.Context) {
	var page = 0
	var limit = 0
	if c.Query("page") != "" {
		paramPage, err := strconv.Atoi(c.Query("page"))
		if err != nil || paramPage < 0 {
			h.logger.Warnf("No url request:%s", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid url query"})
			return
		}
		page = paramPage
	}
	if c.Query("limit") != "" {
		paramLimit, err := strconv.Atoi(c.Query("limit"))
		if err != nil || paramLimit < 0 {
			h.logger.Warnf("No url request:%s", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid url query"})
			return
		}
		limit = paramLimit
	}

	users, err := h.service.AppUser.GetUsers(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)

}

// createUser godoc
// @Summary create master user
// @Description add by json master user
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body model.MUser true "User ID"
// @Success 200 {object} model.MUser
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /user/ [post]
func (h *Handler) createUser(c *gin.Context) {
	var input model.CreateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Warnf("Handler createUser (binding JSON):%s", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	validationErrors := validateStruct(input)
	if len(validationErrors) != 0 {
		h.logger.Warnf("Incorrect data came from the request:%s", validationErrors)
		c.JSON(http.StatusBadRequest, validationErrors)
		return
	}
	user, err := h.service.AppUser.CreateUser(&input)
	if err != nil {
		if err.Error() == "createUser: error while scanning for user:pq: duplicate key value violates unique constraint \"users_email_key\"" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "User with such an email already exists"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusCreated, user)
}

// updateUser godoc
// @Summary update master user
// @Description update by json master user
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body model.MUser true "User ID"
// @Success 200 {object} model.MUser
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /user/ [put]
func (h *Handler) updateUser(c *gin.Context) {
	var input model.UpdateUser
	paramID := c.Query("id")
	varID, err := strconv.Atoi(paramID)
	if err != nil {
		h.logger.Warnf("Handler updateUser (reading param):%s", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid url query"})
		return
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Warnf("Handler updateUser (binding JSON):%s", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	validationErrors := validateStruct(input)
	if len(validationErrors) != 0 {
		h.logger.Warnf("Incorrect data came from the request:%s", validationErrors)
		c.JSON(http.StatusBadRequest, validationErrors)
		return
	}
	id, err := h.service.AppUser.UpdateUser(&input, varID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// deleteUserByID godoc
// @Summary delete a master user by id
// @Description delete user by ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID" Format(int64)
// @Success 200 {object} model.MUser
// @Failure 400 {string} string
// @Failure 404 {object} model.MUser
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /user/{id} [delete]
func (h *Handler) deleteUserByID(c *gin.Context) {
	paramID := c.Param("id")
	varID, err := strconv.Atoi(paramID)
	if err != nil {
		h.logger.Warnf("Handler deleteUserByID (reading param):%s", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}
	id, err := h.service.AppUser.DeleteUserByID(int(varID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}
}

func (h *Handler) grpcFunc(c *gin.Context) {
	var input map[string]string
	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Warnf("Handler createUser (binding JSON):%s", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	resp, err := h.service.AppUser.GrpcExample(input)
	if err != nil {
		return
	} else {
		c.JSON(http.StatusOK, resp)
	}

}
