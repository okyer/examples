package logic

import (
	"context"
	"time"

	"github.com/okyer/examples/gorm4gaussdb/api/demo/internal/model"
	"github.com/okyer/examples/gorm4gaussdb/api/demo/internal/svc"
	"github.com/okyer/examples/gorm4gaussdb/api/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateUser) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	var (
		email    *string
		birthday *time.Time
	)

	if req.Email != "" {
		email = &req.Email
	}
	if req.Birthday != "" {
		t, err := time.Parse("2006-01-02", req.Birthday)
		if err == nil {
			birthday = &t
		}
	}

	user := model.User{
		ID:       uint(req.Id),
		Name:     req.Name,
		Email:    email,
		Age:      uint8(req.Age),
		Birthday: birthday,
	}

	if err := l.svcCtx.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	resp = new(types.Response)
	resp.Message = "success"
	return
}
