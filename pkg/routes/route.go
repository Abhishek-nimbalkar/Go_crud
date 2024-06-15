package routes

import (
	"example.com/crud/pkg/controller"
	"github.com/gin-gonic/gin"
)

// SetRoutes sets the routes for the application
func SetRoutes(routes *gin.Engine) {
	routes.GET("/", controller.ServerHealth)
	routes.POST("/user", controller.CreateUser)
	routes.GET("/user", controller.GetUsers)
	routes.GET("/user/:id", controller.GetUser)
	routes.PATCH("/user", controller.UpdateUser)

}
