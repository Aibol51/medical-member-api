package expert

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-member-api/internal/logic/expert"
	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
)

// swagger:route post /expert/list expert GetExpertList
//
// Get expert list | 获取Expert信息列表
//
// Get expert list | 获取Expert信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ExpertListReq
//
// Responses:
//  200: ExpertListResp

func GetExpertListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExpertListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := expert.NewGetExpertListLogic(r.Context(), svcCtx)
		resp, err := l.GetExpertList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
