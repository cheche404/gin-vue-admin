package system

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

type PrometheusService struct{}

var PrometheusServiceApp = new(PrometheusService)

func (prometheusService *PrometheusService) AddPrometheus(prometheus system.Prometheus) error {
	if !errors.Is(global.GVA_DB.Where("code = ?", prometheus.Code).First(&system.Prometheus{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复code，请修改code")
	}
	return global.GVA_DB.Create(&prometheus).Error
}

func (prometheusService *PrometheusService) DeletePrometheus(prometheus system.Prometheus) error {
	return global.GVA_DB.Delete(&prometheus).Error
}

func (prometheusService *PrometheusService) UpdatePrometheus(prometheus system.Prometheus) error {
	return global.GVA_DB.Updates(&prometheus).Error
}
