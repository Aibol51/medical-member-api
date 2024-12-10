package appointment

import (
	"context"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAppointmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAppointmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAppointmentLogic {
	return &DeleteAppointmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAppointmentLogic) DeleteAppointment(req *types.UUIDsReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.MmsRpc.DeleteAppointment(l.ctx, &mms.UUIDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
