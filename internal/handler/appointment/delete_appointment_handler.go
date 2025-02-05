package appointment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-member-api/internal/logic/appointment"
	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
)

// swagger:route post /appointment/delete appointment DeleteAppointment
//
// Delete appointment information | 删除Appointment信息
//
// Delete appointment information | 删除Appointment信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UUIDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteAppointmentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UUIDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := appointment.NewDeleteAppointmentLogic(r.Context(), svcCtx)
		resp, err := l.DeleteAppointment(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
