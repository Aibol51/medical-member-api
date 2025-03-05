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
			Page:        req.Page,
			PageSize:    req.PageSize,
			PatientName: req.PatientName,
			IdCardNumber: req.IdCardNumber,
			PhoneNumber: req.PhoneNumber,
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
        	PatientName: v.PatientName,
        	Gender: v.Gender,
        	Age: v.Age,
        	IdCardNumber: v.IdCardNumber,
        	PhoneNumber: v.PhoneNumber,
        	ChiefComplaint: v.ChiefComplaint,
        	PresentIllness: v.PresentIllness,
        	PastHistory: v.PastHistory,
        	SmokingHistory: v.SmokingHistory,
        	DrinkingHistory: v.DrinkingHistory,
        	AllergyHistory: v.AllergyHistory,
        	HeartRate: v.HeartRate,
        	BloodPressure: v.BloodPressure,
        	OxygenSaturation: v.OxygenSaturation,
        	BloodGlucose: v.BloodGlucose,
        	Weight: v.Weight,
        	WaistCircumference: v.WaistCircumference,
        	BodyFat: v.BodyFat,
        	Diagnosis: v.Diagnosis,
        	DietTherapy: v.DietTherapy,
        	ExerciseTherapy: v.ExerciseTherapy,
        	MedicationTherapy: v.MedicationTherapy,
        	TreatmentPlan: v.TreatmentPlan,
        	DoctorId: v.DoctorId,
        	AppointmentId: v.AppointmentId,
        	Remarks: v.Remarks,
        	UserId: v.UserId,
			})
	}
	return resp, nil
}
