package appointment

import (
	"context"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppointmentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAppointmentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppointmentListLogic {
	return &GetAppointmentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAppointmentListLogic) GetAppointmentList(req *types.AppointmentListReq) (resp *types.AppointmentListResp, err error) {
	data, err := l.svcCtx.MmsRpc.GetAppointmentList(l.ctx,
		&mms.AppointmentListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			UserId:   l.ctx.Value("userId").(string),
		})
	if err != nil {
		return nil, err
	}
	resp = &types.AppointmentListResp{}
	resp.Msg = "successful"
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.AppointmentInfo{
				BaseUUIDInfo: types.BaseUUIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				PatientName:     v.PatientName,
				PhoneNumber:     v.PhoneNumber,
				IdCard:          v.IdCard,
				Gender:          v.Gender,
				Age:             v.Age,
				AppointmentTime: v.AppointmentTime,
				Symptoms:        v.Symptoms,
				Status:          v.Status,
				Remarks:         v.Remarks,
				UserId:          v.UserId,
			})
	}
	return resp, nil
}
