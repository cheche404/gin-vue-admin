package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PrometheusApi struct{}

func (s *PrometheusApi) GetPrometheusList(c *gin.Context) {
	var pageInfo systemReq.SearchPrometheusParams
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := prometheusService.GetPrometheusInfoList(pageInfo.Prometheus, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (s *PrometheusApi) GetPrometheusById(c *gin.Context) {
	var idInfo request.GetById
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(idInfo, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	prometheus, err := prometheusService.GetPrometheusById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.MonitorPrometheusResponse{Prometheus: prometheus}, "获取成功", c)
}

func (a *PrometheusApi) CreatePrometheus(c *gin.Context) {
	var prometheus system.Prometheus
	err := c.ShouldBindJSON(&prometheus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	errVerify := utils.Verify(prometheus, utils.PrometheusVerify)
	if errVerify != nil {
		response.FailWithMessage(errVerify.Error(), c)
		return
	}

	errAdd := prometheusService.AddPrometheus(prometheus)
	if errAdd != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage(errVerify.Error(), c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

func (a *PrometheusApi) DeletePrometheus(c *gin.Context) {
	var prometheus system.Prometheus
	err := c.ShouldBindJSON(&prometheus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	errDelete := prometheusService.DeletePrometheus(prometheus)
	if errDelete != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(errDelete.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (a *PrometheusApi) UpdatePrometheus(c *gin.Context) {
	var prometheus system.Prometheus
	err := c.ShouldBindJSON(&prometheus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	errVerify := utils.Verify(prometheus, utils.PrometheusVerify)
	if errVerify != nil {
		response.FailWithMessage(errVerify.Error(), c)
		return
	}

	errUpdate := prometheusService.UpdatePrometheus(prometheus)
	if errUpdate != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(errUpdate.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}
