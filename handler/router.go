package handler

import (
	"github.com/Baraulia/AUTHENTICATION_SERVICE/middleware"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/service"
	"github.com/gin-gonic/gin"
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
