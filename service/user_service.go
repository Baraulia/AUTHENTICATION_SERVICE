package service

import (
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RoutesUser(rg *gin.RouterGroup) {
	user := rg.Group("/user")

	user.GET("/:id",  getUser)
	user.GET("/", getUsers)
	user.POST("/", createUser)
	user.PUT("/:id", updateUser)
	user.DELETE("/:id", deleteUserByID)
}

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
func getUser(c *gin.Context) {
	var user model.User
	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err = repository.GetUserByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if (model.User{}) == user {
		c.JSON(http.StatusNotFound, user)
	} else {
		c.JSON(http.StatusOK, user)
	}
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
func getUsers(c *gin.Context) {

	var users []model.User
	users, err := repository.GetUserAll()
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
func createUser(c *gin.Context) {

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	user, err := repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err})
		return
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
func updateUser(c *gin.Context) {

	var user model.User
	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}
	user.ID = varID
	//usr, err := repository.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	//c.JSON(http.StatusOK, usr)
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
func deleteUserByID(c *gin.Context) {

	var user model.User

	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = repository.DeleteUserByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, user)
}
