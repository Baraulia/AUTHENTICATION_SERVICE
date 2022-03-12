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
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /users/{id} [get]
func (h *Handler) getUser(ctx *gin.Context) {
	paramID := ctx.Param("id")
	varID, err := strconv.Atoi(paramID)
	if err != nil || varID <= 0 {
		h.logger.Warnf("Handler getUser (reading param):%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
		return
	}
	user, err := h.service.AppUser.GetUser(varID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
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
// @Param input body model.ResponseFilters false "Filters"
// @Success 200 {object} listUsers
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /users/ [get]
func (h *Handler) getUsers(ctx *gin.Context) {
	var page = 0
	var limit = 0
	var filters model.ResponseFilters
	if ctx.Query("page") != "" {
		paramPage, err := strconv.Atoi(ctx.Query("page"))
		if err != nil || paramPage < 0 {
			h.logger.Warnf("No url request:%s", err)
			ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Invalid url query"})
			return
		}
		page = paramPage
	}
	if ctx.Query("limit") != "" {
		paramLimit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil || paramLimit < 0 {
			h.logger.Warnf("No url request:%s", err)
			ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Invalid url query"})
			return
		}
		limit = paramLimit
	}
	if ctx.Request.Body != nil {
		if err := ctx.ShouldBindJSON(&filters); err != nil {
			h.logger.Warnf("Handler getUsers (binding JSON):%s", err)
			ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request body"})
			return
		}

	}
	users, pages, err := h.service.AppUser.GetUsers(page, limit, &filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.Header("pages", strconv.Itoa(pages))
	ctx.JSON(http.StatusOK, listUsers{Data: users})

}

// createCustomer godoc
// @Summary createCustomer
// @Description create new customer
// @Tags User
// @Accept  json
// @Produce  json
// @Param input body model.CreateCustomer true "User"
// @Success 201 {object} authProto.GeneratedTokens
// @Failure 400 {object} model.ErrorResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} model.ErrorResponse
// @Router /users/customer [post]
func (h *Handler) createCustomer(ctx *gin.Context) {
	var input model.CreateCustomer
	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.logger.Warnf("Handler createUser (binding JSON):%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
		return
	}
	validationErrors := ValidateStruct(input)
	if len(validationErrors) != 0 {
		h.logger.Warnf("Incorrect data came from the request:%s", validationErrors)
		ctx.JSON(http.StatusBadRequest, validationErrors)
		return
	}
	tokens, id, err := h.service.AppUser.CreateCustomer(&input)
	if err != nil {
		if err.Error() == "createCustomer: error while scanning for user:pq: duplicate key value violates unique constraint \"users_email_key\"" {
			ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "User with such an email already exists"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
			return
		}
	}
	ctx.Header("id", strconv.Itoa(id))
	ctx.JSON(http.StatusCreated, tokens)
}

// createStaff godoc
// @Summary createStaff
// @Description create new restaurant or courier manager or courier
// @Tags User
// @Accept  json
// @Produce  json
// @Param input body model.CreateStaff true "User"
// @Success 201 {string} string
// @Failure 400 {object} model.ErrorResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} model.ErrorResponse
// @Router /users/staff [post]
func (h *Handler) createStaff(ctx *gin.Context) {
	var input model.CreateStaff
	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.logger.Warnf("Handler createUser (binding JSON):%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
		return
	}
	validationErrors := ValidateStruct(input)
	if len(validationErrors) != 0 {
		h.logger.Warnf("Incorrect data came from the request:%s", validationErrors)
		ctx.JSON(http.StatusBadRequest, validationErrors)
		return
	}
	err := h.service.AppUser.CheckInputRole(input.Role)
	if err != nil {
		h.logger.Warnf("Incorrect role came from the request:%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Incorrect role came from the request"})
		return
	}
	id, err := h.service.AppUser.CreateStaff(&input)
	if err != nil {
		if err.Error() == "createStaff: error while scanning for user:pq: duplicate key value violates unique constraint users_email_key" {
			ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "User with such an email already exists"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
			return
		}
	}
	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

// updateUser godoc
// @Summary updateUser
// @Description change user password
// @Tags User
// @Accept  json
// @Produce  json
// @Param input body model.UpdateUser true "User"
// @Param id path int true "User ID" Format(int64)
// @Success 204
// @Failure 400 {object} model.ErrorResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} model.ErrorResponse
// @Router /users/{id} [put]
func (h *Handler) updateUser(ctx *gin.Context) {
	var input model.UpdateUser
	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.logger.Warnf("Handler updateUser (binding JSON):%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
		return
	}
	validationErrors := ValidateStruct(input)
	if len(validationErrors) != 0 {
		h.logger.Warnf("Incorrect data came from the request:%s", validationErrors)
		ctx.JSON(http.StatusBadRequest, validationErrors)
		return
	}
	err := h.service.AppUser.UpdateUser(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

// deleteUserByID godoc
// @Summary deleteUserByID
// @Description delete user by ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID" Format(int64)
// @Success 200  {string} string
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /users/{id} [delete]
func (h *Handler) deleteUserByID(ctx *gin.Context) {
	paramID := ctx.Param("id")
	varID, err := strconv.Atoi(paramID)
	if err != nil || varID <= 0 {
		h.logger.Warnf("Handler deleteUserByID (reading param):%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Invalid id"})
		return
	}
	id, err := h.service.AppUser.DeleteUserByID(int(varID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}
}
