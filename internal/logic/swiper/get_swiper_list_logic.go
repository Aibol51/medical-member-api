package swiper

import (
	"context"

    "github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSwiperListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSwiperListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSwiperListLogic {
	return &GetSwiperListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSwiperListLogic) GetSwiperList(req *types.SwiperListReq) (resp *types.SwiperListResp, err error) {
	data, err := l.svcCtx.MmsRpc.GetSwiperList(l.ctx,
		&mms.SwiperListReq{
			Page:        req.Page,
			PageSize:    req.PageSize,
			TitleZh: req.TitleZh,
			TitleEn: req.TitleEn,
			TitleRu: req.TitleRu,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.SwiperListResp{}
	resp.Msg = "successful"
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.SwiperInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
        	Status: v.Status,
        	Sort: v.Sort,
        	TitleZh: v.TitleZh,
        	TitleEn: v.TitleEn,
        	TitleRu: v.TitleRu,
        	TitleKk: v.TitleKk,
        	BannerZh: v.BannerZh,
        	BannerEn: v.BannerEn,
        	BannerRu: v.BannerRu,
        	BannerKk: v.BannerKk,
        	ContentZh: v.ContentZh,
        	ContentEn: v.ContentEn,
        	ContentRu: v.ContentRu,
        	ContentKk: v.ContentKk,
        	JumpUrl: v.JumpUrl,
			})
	}
	return resp, nil
}
