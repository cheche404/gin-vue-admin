package system

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type PrometheusInterface interface {
	GetName() string
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
