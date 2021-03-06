package rest

import (
	"github.com/gin-gonic/gin"
	"ubiwhere/controller"
)

func SetupRouter() *gin.Engine {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	ubiwhere := router.Group("/ubiwhere")
	{
		ubiwhere.GET("/read/:n", getNMetrics)
		ubiwhere.GET("/read/:n/vars", getNMetricsVars)
		ubiwhere.GET("/avg", getAvgVars)
	}

	return router
}

func getNMetrics(c *gin.Context) {
	controller.GetNMetrics(c)
}

func getNMetricsVars(c *gin.Context) {
	controller.GetNMetricsVars(c)
}

func getAvgVars(c *gin.Context) {
	controller.GetAvgVars(c)
}
