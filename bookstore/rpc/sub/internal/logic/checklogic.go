package logic

import (
	"context"

	"bookstore/rpc/sub/internal/svc"
	"bookstore/rpc/sub/sub"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *sub.SubReq) (*sub.SubResp, error) {
	// todo: add your logic here and delete this line

	return &sub.SubResp{}, nil
}
