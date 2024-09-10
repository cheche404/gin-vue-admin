import service from '@/utils/request'

// @Tags api
// @Summary 分页获取prometheus列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取prometheus列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
// {
//  page     int
//	pageSize int
// }
export const getPrometheusList = (data) => {
    return service({
        url: '/prometheus/getPrometheusList',
        method: 'post',
        data
    })
}

export const getPrometheusById = (data) => {
    return service({
        url: '/prometheus/getPrometheusById',
        method: 'post',
        data
    })
}