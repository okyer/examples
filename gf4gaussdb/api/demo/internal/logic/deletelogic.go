package logic

import (
	"context"
	"fmt"

	"github.com/okyer/examples/gf4gaussdb/api/demo/internal/svc"
	"github.com/okyer/examples/gf4gaussdb/api/demo/internal/types"

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
	// rs, err := l.svcCtx.DB.Delete(l.ctx, "users", model.User{ID: uint(req.Id)})
	rs, err := l.svcCtx.DB.Model("users").Where("id = ?", req.Id).Delete()
	if err != nil {
		return nil, err
	}

	rows, _ := rs.RowsAffected()

	resp = new(types.Response)
	resp.Message = fmt.Sprintf("success, affected rows: %d", rows)
	return
}
