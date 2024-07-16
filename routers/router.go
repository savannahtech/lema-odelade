package routers

import (
	"accessment.com/microservice/db/postgres"
	"accessment.com/microservice/service"
	"github.com/gin-gonic/gin"
)

func Routers() {
	router := gin.Default()
	postgres.MigrateTables()
	servicesRoutes(router)
	router.Run(":8085")
}

func servicesRoutes(router *gin.Engine) {
	var route *gin.RouterGroup = router.Group("/api/services")

	route.GET("/repo", service.RepService.GetRepoDetails)
	route.GET("/commit", service.RepService.GetCommits)
}
