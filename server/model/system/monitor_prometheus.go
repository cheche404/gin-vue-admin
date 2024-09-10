package system

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type PrometheusInterface interface {
	GetName() string
	GetCode() string
	GetUrl() string
}

type Prometheus struct {
	global.GVA_MODEL
	Name string `json:"name"`
	Code string `json:"code"`
	Url  string `json:"url"`
}

func (Prometheus) TableName() string {
	return "prometheus"
}

func (prometheus *Prometheus) GetName() string {
	return prometheus.Name
}

func (prometheus *Prometheus) GetCode() string {
	return prometheus.Code
}

func (prometheus *Prometheus) GetUrl() string {
	return prometheus.Url
}
