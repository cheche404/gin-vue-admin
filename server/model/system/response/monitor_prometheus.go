package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

type MonitorPrometheusResponse struct {
	Prometheus system.Prometheus `json:"prometheus"`
}
