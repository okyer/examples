package logic

import (
	"context"

	"github.com/okyer/examples/gorm4gaussdb/api/demo/internal/model"
	"github.com/okyer/examples/gorm4gaussdb/api/demo/internal/svc"
	"github.com/okyer/examples/gorm4gaussdb/api/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.DelUser) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	if err := l.svcCtx.DB.Delete(&model.User{}, req.Id).Error; err != nil {
		return nil, err
	}

	resp = new(types.Response)
	resp.Message = "success"
	return
}
