package handler

import (
	"github.com/gin-gonic/gin"
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

	router.Use(
		middleware.CorsMiddleware,
	)

	user := router.Group("/users")
	{
		user.GET("/:id", h.getUser)
		user.GET("/", h.getUsers)
		user.POST("/", h.createUser)
		user.PUT("/:id", h.updateUser)
		user.DELETE("/:id", h.deleteUserByID)
		user.POST("/login", h.authUser)
		user.GET("/grpc", h.grpcFunc)
	}

	return router
}
