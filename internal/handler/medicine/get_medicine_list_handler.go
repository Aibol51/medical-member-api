package medicine

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-member-api/internal/logic/medicine"
	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
)

// swagger:route post /medicine/list medicine GetMedicineList
//
// Get medicine list | 获取Medicine信息列表
//
// Get medicine list | 获取Medicine信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: MedicineListReq
//
// Responses:
//  200: MedicineListResp

func GetMedicineListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MedicineListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := medicine.NewGetMedicineListLogic(r.Context(), svcCtx)
		resp, err := l.GetMedicineList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
