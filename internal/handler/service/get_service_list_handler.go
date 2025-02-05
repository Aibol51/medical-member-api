package service

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-member-api/internal/logic/service"
	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
)

// swagger:route post /service/list service GetServiceList
//
// Get service list | 获取Service信息列表
//
// Get service list | 获取Service信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ServiceListReq
//
// Responses:
//  200: ServiceListResp

func GetServiceListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ServiceListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := service.NewGetServiceListLogic(r.Context(), svcCtx)
		resp, err := l.GetServiceList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
