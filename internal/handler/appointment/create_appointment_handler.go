package appointment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-member-api/internal/logic/appointment"
	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
)

// swagger:route post /appointment/create appointment CreateAppointment
//
// Create appointment information | 创建Appointment信息
//
// Create appointment information | 创建Appointment信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: AppointmentInfo
//
// Responses:
//  200: BaseMsgResp

func CreateAppointmentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppointmentInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := appointment.NewCreateAppointmentLogic(r.Context(), svcCtx)
		resp, err := l.CreateAppointment(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
