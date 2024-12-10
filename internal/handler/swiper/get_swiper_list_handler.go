package swiper

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-member-api/internal/logic/swiper"
	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
)

// swagger:route post /swiper/list swiper GetSwiperList
//
// Get swiper list | 获取Swiper信息列表
//
// Get swiper list | 获取Swiper信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SwiperListReq
//
// Responses:
//  200: SwiperListResp

func GetSwiperListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SwiperListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := swiper.NewGetSwiperListLogic(r.Context(), svcCtx)
		resp, err := l.GetSwiperList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
