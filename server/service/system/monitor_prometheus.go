package system

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
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

func (prometheusService *PrometheusService) GetPrometheusInfoList(prometheus system.Prometheus, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.Prometheus{})
	var prometheusList []system.Prometheus

	if prometheus.Url != "" {
		db = db.Where("url LIKE ?", "%"+prometheus.Url+"%")
	}

	if prometheus.Name != "" {
		db = db.Where("name LIKE ?", "%"+prometheus.Name+"%")
	}

	if prometheus.Code != "" {
		db = db.Where("code LIKE ?", "%"+prometheus.Code+"%")
	}

	err = db.Count(&total).Error

	if err != nil {
		return prometheusList, total, err
	}

	db = db.Limit(limit).Offset(offset)
	OrderStr := "id desc"
	if order != "" {
		orderMap := make(map[string]bool, 5)
		orderMap["id"] = true
		orderMap["name"] = true
		orderMap["code"] = true
		orderMap["url"] = true
		if !orderMap[order] {
			err = fmt.Errorf("非法的排序字段: %v", order)
			return prometheusList, total, err
		}
		OrderStr = order
		if desc {
			OrderStr = order + " desc"
		}
	}
	err = db.Order(OrderStr).Find(&prometheusList).Error
	return prometheusList, total, err
}

func (prometheusService *PrometheusService) GetPrometheusById(id int) (prometheus system.Prometheus, err error) {
	err = global.GVA_DB.First(&prometheus, "id = ?", id).Error
	return
}
