package expert

import (
    "context"

    "github.com/suyuan32/simple-admin-member-api/internal/svc"
    "github.com/suyuan32/simple-admin-member-api/internal/types"
    "github.com/suyuan32/simple-admin-member-rpc/types/mms"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetExpertByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetExpertByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExpertByIdLogic {
	return &GetExpertByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExpertByIdLogic) GetExpertById(req *types.IDReq) (resp *types.ExpertInfoResp, err error) {
	data, err := l.svcCtx.MmsRpc.GetExpertById(l.ctx, &mms.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.ExpertInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  "successful",
		},
		Data: types.ExpertInfo{
            BaseIDInfo: types.BaseIDInfo{
                Id:        data.Id,
                CreatedAt: data.CreatedAt,
                UpdatedAt: data.UpdatedAt,
            },
        	Status: data.Status,
        	Sort: data.Sort,
        	NameZh: data.NameZh,
        	NameEn: data.NameEn,
        	NameRu: data.NameRu,
        	NameKk: data.NameKk,
        	ContentZh: data.ContentZh,
        	ContentEn: data.ContentEn,
        	ContentRu: data.ContentRu,
        	ContentKk: data.ContentKk,
        	CoverUrl: data.CoverUrl,
		},
	}, nil
}

