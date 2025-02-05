package service

import (
    "context"

    "github.com/suyuan32/simple-admin-member-api/internal/svc"
    "github.com/suyuan32/simple-admin-member-api/internal/types"
    "github.com/suyuan32/simple-admin-member-rpc/types/mms"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetServiceByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetServiceByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServiceByIdLogic {
	return &GetServiceByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetServiceByIdLogic) GetServiceById(req *types.IDReq) (resp *types.ServiceInfoResp, err error) {
	data, err := l.svcCtx.MmsRpc.GetServiceById(l.ctx, &mms.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.ServiceInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  "successful",
		},
		Data: types.ServiceInfo{
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
        	DescriptionZh: data.DescriptionZh,
        	DescriptionEn: data.DescriptionEn,
        	DescriptionRu: data.DescriptionRu,
        	DescriptionKk: data.DescriptionKk,
        	CoverUrl: data.CoverUrl,
		},
	}, nil
}

