package routes

import (
	"github.com/AlissonBarbosa/shylockgo-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
  router.GET("/projects", controllers.GetProjects)
}
