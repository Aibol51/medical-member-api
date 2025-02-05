package expert

import (
	"context"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExpertListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetExpertListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExpertListLogic {
	return &GetExpertListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExpertListLogic) GetExpertList(req *types.ExpertListReq) (resp *types.ExpertListResp, err error) {
	data, err := l.svcCtx.MmsRpc.GetExpertList(l.ctx,
		&mms.ExpertListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			NameZh:   req.NameZh,
			NameEn:   req.NameEn,
			NameRu:   req.NameRu,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.ExpertListResp{}
	resp.Msg = "successful"
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.ExpertInfo{
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
