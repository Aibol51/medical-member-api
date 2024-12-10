package news

import (
	"context"

    "github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNewsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNewsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNewsListLogic {
	return &GetNewsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNewsListLogic) GetNewsList(req *types.NewsListReq) (resp *types.NewsListResp, err error) {
	data, err := l.svcCtx.MmsRpc.GetNewsList(l.ctx,
		&mms.NewsListReq{
			Page:        req.Page,
			PageSize:    req.PageSize,
			TitleZh: req.TitleZh,
			TitleEn: req.TitleEn,
			TitleRu: req.TitleRu,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.NewsListResp{}
	resp.Msg = "successful"
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.NewsInfo{
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
        	ContentZh: v.ContentZh,
        	ContentEn: v.ContentEn,
        	ContentRu: v.ContentRu,
        	ContentKk: v.ContentKk,
        	CoverUrl: v.CoverUrl,
			})
	}
	return resp, nil
}
