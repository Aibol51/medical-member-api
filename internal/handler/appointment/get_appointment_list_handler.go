package appointment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-member-api/internal/logic/appointment"
	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
)

// swagger:route post /appointment/list appointment GetAppointmentList
//
// Get appointment list | 获取Appointment信息列表
//
// Get appointment list | 获取Appointment信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: AppointmentListReq
//
// Responses:
//  200: AppointmentListResp

func GetAppointmentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppointmentListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := appointment.NewGetAppointmentListLogic(r.Context(), svcCtx)
		resp, err := l.GetAppointmentList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
