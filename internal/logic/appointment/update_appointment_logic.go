package appointment

import (
	"context"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAppointmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAppointmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAppointmentLogic {
	return &UpdateAppointmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAppointmentLogic) UpdateAppointment(req *types.AppointmentInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.MmsRpc.UpdateAppointment(l.ctx,
		&mms.AppointmentInfo{
			Id:          req.Id,
        	PatientName: req.PatientName,
        	PhoneNumber: req.PhoneNumber,
        	IdCard: req.IdCard,
        	Gender: req.Gender,
        	Age: req.Age,
        	AppointmentTime: req.AppointmentTime,
        	Symptoms: req.Symptoms,
        	Status: req.Status,
        	Remarks: req.Remarks,
        	UserId: req.UserId,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
