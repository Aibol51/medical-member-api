package appointment

import (
	"context"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppointmentByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAppointmentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppointmentByIdLogic {
	return &GetAppointmentByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAppointmentByIdLogic) GetAppointmentById(req *types.UUIDReq) (resp *types.AppointmentInfoResp, err error) {
	l.Infof("ctx: %v; userId: %v;", l.ctx, l.ctx.Value("userId"))
	data, err := l.svcCtx.MmsRpc.GetAppointmentById(l.ctx, &mms.UUIDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.AppointmentInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  "successful",
		},
		Data: types.AppointmentInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			PatientName:     data.PatientName,
			PhoneNumber:     data.PhoneNumber,
			IdCard:          data.IdCard,
			Gender:          data.Gender,
			Age:             data.Age,
			AppointmentTime: data.AppointmentTime,
			Symptoms:        data.Symptoms,
			Status:          data.Status,
			Remarks:         data.Remarks,
			UserId:          data.UserId,
		},
	}, nil
}
