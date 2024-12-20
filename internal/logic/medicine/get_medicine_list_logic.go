package medicine

import (
	"context"

    "github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMedicineListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMedicineListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMedicineListLogic {
	return &GetMedicineListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMedicineListLogic) GetMedicineList(req *types.MedicineListReq) (resp *types.MedicineListResp, err error) {
	data, err := l.svcCtx.MmsRpc.GetMedicineList(l.ctx,
		&mms.MedicineListReq{
			Page:        req.Page,
			PageSize:    req.PageSize,
			NameZh: req.NameZh,
			NameEn: req.NameEn,
			NameRu: req.NameRu,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.MedicineListResp{}
	resp.Msg = "successful"
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.MedicineInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
        	Status: v.Status,
        	Sort: v.Sort,
        	NameZh: v.NameZh,
        	NameEn: v.NameEn,
        	NameRu: v.NameRu,
        	NameKk: v.NameKk,
        	Quantity: v.Quantity,
        	DescriptionZh: v.DescriptionZh,
        	DescriptionEn: v.DescriptionEn,
        	DescriptionRu: v.DescriptionRu,
        	DescriptionKk: v.DescriptionKk,
        	Remarks: v.Remarks,
        	Images: v.Images,
			})
	}
	return resp, nil
}
