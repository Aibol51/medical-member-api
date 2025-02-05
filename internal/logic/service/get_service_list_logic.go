package service

import (
	"context"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServiceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetServiceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServiceListLogic {
	return &GetServiceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetServiceListLogic) GetServiceList(req *types.ServiceListReq) (resp *types.ServiceListResp, err error) {
	data, err := l.svcCtx.MmsRpc.GetServiceList(l.ctx,
		&mms.ServiceListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			NameZh:   req.NameZh,
			NameEn:   req.NameEn,
			NameRu:   req.NameRu,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.ServiceListResp{}
	resp.Msg = "successful"
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.ServiceInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Status:   v.Status,
				Sort:     v.Sort,
				NameZh:   v.NameZh,
				NameEn:   v.NameEn,
				NameRu:   v.NameRu,
				NameKk:   v.NameKk,
				CoverUrl: v.CoverUrl,
			})
	}
	return resp, nil
}
