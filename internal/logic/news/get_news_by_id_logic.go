package news

import (
    "context"

    "github.com/suyuan32/simple-admin-member-api/internal/svc"
    "github.com/suyuan32/simple-admin-member-api/internal/types"
    "github.com/suyuan32/simple-admin-member-rpc/types/mms"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetNewsByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNewsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNewsByIdLogic {
	return &GetNewsByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNewsByIdLogic) GetNewsById(req *types.IDReq) (resp *types.NewsInfoResp, err error) {
	data, err := l.svcCtx.MmsRpc.GetNewsById(l.ctx, &mms.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.NewsInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  "successful",
		},
		Data: types.NewsInfo{
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
        	ContentZh: data.ContentZh,
        	ContentEn: data.ContentEn,
        	ContentRu: data.ContentRu,
        	ContentKk: data.ContentKk,
        	CoverUrl: data.CoverUrl,
		},
	}, nil
}

