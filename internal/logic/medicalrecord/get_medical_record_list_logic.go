package medicalrecord

import (
	"context"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMedicalRecordListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMedicalRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMedicalRecordListLogic {
	return &GetMedicalRecordListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMedicalRecordListLogic) GetMedicalRecordList(req *types.MedicalRecordListReq) (resp *types.MedicalRecordListResp, err error) {
	data, err := l.svcCtx.MmsRpc.GetMedicalRecordList(l.ctx,
		&mms.MedicalRecordListReq{
			Page:          req.Page,
			PageSize:      req.PageSize,
			PatientName:   req.PatientName,
			PhoneNumber:   req.PhoneNumber,
			UserId:        l.ctx.Value("userId").(string),
			AppointmentId: req.AppointmentId,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.MedicalRecordListResp{}
	resp.Msg = "successful"
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.MedicalRecordInfo{
				BaseUUIDInfo: types.BaseUUIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				PatientName:        v.PatientName,
				PhoneNumber:        v.PhoneNumber,
				Gender:             v.Gender,
				Age:                v.Age,
				VisitTime:          v.VisitTime,
				Diagnosis:          v.Diagnosis,
				TreatmentPlan:      v.TreatmentPlan,
				Prescription:       v.Prescription,
				ExaminationResults: v.ExaminationResults,
				DoctorAdvice:       v.DoctorAdvice,
				DoctorId:           v.DoctorId,
				Department:         v.Department,
				AppointmentId:      v.AppointmentId,
				Remarks:            v.Remarks,
				UserId:             v.UserId,
			})
	}
	return resp, nil
}
