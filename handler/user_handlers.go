package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/model"
	"strconv"
)

// getUserByID godoc
// @Summary getUser
// @Description get user by ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} model.ResponseUser
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users/{id} [get]
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

type listUsers struct {
	Data []model.ResponseUser
}

// getUsers godoc
// @Summary getUsers
// @Description get list of users
// @Tags User
// @Accept  json
// @Produce  json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} listUsers
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users/ [get]
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

	users, pages, err := h.service.AppUser.GetUsers(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.Header("pages", strconv.Itoa(pages))
	c.JSON(http.StatusOK, listUsers{Data: users})

}

// createUser godoc
// @Summary createUser
// @Description create new user
// @Tags User
// @Accept  json
// @Produce  json
// @Param input body model.CreateUser true "User"
// @Success 201 {object} auth_proto.GeneratedTokens
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users/ [post]
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
	tokens, id, err := h.service.AppUser.CreateUser(&input)
	if err != nil {
		if err.Error() == "createUser: error while scanning for user:pq: duplicate key value violates unique constraint \"users_email_key\"" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "User with such an email already exists"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.Header("id", strconv.Itoa(id))
	c.JSON(http.StatusCreated, tokens)
}

// updateUser godoc
// @Summary updateUser
// @Description change user password
// @Tags User
// @Accept  json
// @Produce  json
// @Param input body model.UpdateUser true "User"
// @Param id query int true "Id"
// @Success 204
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users/ [put]
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
	err = h.service.AppUser.UpdateUser(&input, varID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// deleteUserByID godoc
// @Summary deleteUserByID
// @Description delete user by ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID" Format(int64)
// @Success 200  {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users/{id} [delete]
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
	var input string
	input = c.Query("token")
	proto, err := h.service.AppUser.GrpcExample(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, proto)
	}

}
