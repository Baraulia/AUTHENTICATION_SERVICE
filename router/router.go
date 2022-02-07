package router
import (
	"github.com/Baraulia/AUTHENTICATION_SERVICE/service"
	"github.com/gin-gonic/gin"
)

func NewRoutes() *gin.Engine {

	router := gin.Default()
	forUser := router.Group("/api")

	// register router from each controller service
	service.RoutesUser(forUser)

	return router
}
