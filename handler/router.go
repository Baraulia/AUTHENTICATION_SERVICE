package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "stlab.itechart-group.com/go/food_delivery/authentication_service/docs"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/middleware"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/service"
)

type Handler struct {
	logger  logging.Logger
	service *service.Service
}

func NewHandler(logger logging.Logger, service *service.Service) *Handler {
	return &Handler{logger: logger, service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(
		middleware.CorsMiddleware,
	)

	user := router.Group("/users")
	{
		user.GET("/:id", h.getUser)
		user.DELETE("/:id", h.deleteUserByID)
		user.PUT("/:id", h.updateUser)
		user.GET("/", h.getUsers)
		user.POST("/staff", h.createStaff)
		user.POST("/customer", h.createCustomer)
		user.POST("/login", h.authUser)
	}

	return router
}
