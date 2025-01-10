package medicalrecord

import (
    "context"

    "github.com/suyuan32/simple-admin-member-api/internal/svc"
    "github.com/suyuan32/simple-admin-member-api/internal/types"
    "github.com/suyuan32/simple-admin-member-rpc/types/mms"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetMedicalRecordByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMedicalRecordByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMedicalRecordByIdLogic {
	return &GetMedicalRecordByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMedicalRecordByIdLogic) GetMedicalRecordById(req *types.UUIDReq) (resp *types.MedicalRecordInfoResp, err error) {
	data, err := l.svcCtx.MmsRpc.GetMedicalRecordById(l.ctx, &mms.UUIDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.MedicalRecordInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  "successful",
		},
		Data: types.MedicalRecordInfo{
            BaseUUIDInfo: types.BaseUUIDInfo{
                Id:        data.Id,
                CreatedAt: data.CreatedAt,
                UpdatedAt: data.UpdatedAt,
            },
        	PatientName: data.PatientName,
        	PhoneNumber: data.PhoneNumber,
        	Gender: data.Gender,
        	Age: data.Age,
        	VisitTime: data.VisitTime,
        	Diagnosis: data.Diagnosis,
        	TreatmentPlan: data.TreatmentPlan,
        	Prescription: data.Prescription,
        	ExaminationResults: data.ExaminationResults,
        	DoctorAdvice: data.DoctorAdvice,
        	DoctorId: data.DoctorId,
        	Department: data.Department,
        	AppointmentId: data.AppointmentId,
        	Remarks: data.Remarks,
        	UserId: data.UserId,
		},
	}, nil
}

