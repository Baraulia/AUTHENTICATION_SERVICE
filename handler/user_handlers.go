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
	var user *model.User
	paramID := c.Param("id")
	varID, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err = h.service.GetUser(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
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
	users, err := h.service.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
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

	var user *model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	user, err := h.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// updateUser godoc
// @Summary update master user
// @Description updates by json master user
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
	var user model.User
	paramID := c.Param("id")
	varID, err := strconv.Atoi(paramID)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	usr, err := h.service.UpdateUser(user, varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, usr)
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
	varID, err := strconv.ParseInt(paramID, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = h.service.DeleteUserByID(int(varID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}else {
		c.JSON(http.StatusOK, gin.H{"message": "successful"})
	}
}
