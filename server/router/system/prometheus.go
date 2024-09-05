package system

import (
	"github.com/gin-gonic/gin"
)

type PrometheusRoute struct{}

func (s *PrometheusRoute) InitPrometheusRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	prometheusRoute := Router.Group("prometheus")
	{
		prometheusRoute.POST("addPrometheus", prometheusApi.CreatePrometheus)
		prometheusRoute.PUT("updatePrometheus", prometheusApi.UpdatePrometheus)
		prometheusRoute.DELETE("deletePrometheus", prometheusApi.DeletePrometheus)
	}
	return prometheusRoute
}
