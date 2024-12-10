package swiper

import (
    "context"

    "github.com/suyuan32/simple-admin-member-api/internal/svc"
    "github.com/suyuan32/simple-admin-member-api/internal/types"
    "github.com/suyuan32/simple-admin-member-rpc/types/mms"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetSwiperByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSwiperByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSwiperByIdLogic {
	return &GetSwiperByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSwiperByIdLogic) GetSwiperById(req *types.IDReq) (resp *types.SwiperInfoResp, err error) {
	data, err := l.svcCtx.MmsRpc.GetSwiperById(l.ctx, &mms.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.SwiperInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  "successful",
		},
		Data: types.SwiperInfo{
            BaseIDInfo: types.BaseIDInfo{
                Id:        data.Id,
                CreatedAt: data.CreatedAt,
                UpdatedAt: data.UpdatedAt,
            },
        	Status: data.Status,
        	Sort: data.Sort,
        	TitleZh: data.TitleZh,
        	TitleEn: data.TitleEn,
        	TitleRu: data.TitleRu,
        	TitleKk: data.TitleKk,
        	BannerZh: data.BannerZh,
        	BannerEn: data.BannerEn,
        	BannerRu: data.BannerRu,
        	BannerKk: data.BannerKk,
        	ContentZh: data.ContentZh,
        	ContentEn: data.ContentEn,
        	ContentRu: data.ContentRu,
        	ContentKk: data.ContentKk,
        	JumpUrl: data.JumpUrl,
		},
	}, nil
}

