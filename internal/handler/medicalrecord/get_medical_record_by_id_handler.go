package medicalrecord

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-member-api/internal/logic/medicalrecord"
	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
)

// swagger:route post /medical_record medicalrecord GetMedicalRecordById
//
// Get medical record by ID | 通过ID获取MedicalRecord信息
//
// Get medical record by ID | 通过ID获取MedicalRecord信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UUIDReq
//
// Responses:
//  200: MedicalRecordInfoResp

func GetMedicalRecordByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UUIDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := medicalrecord.NewGetMedicalRecordByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetMedicalRecordById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
