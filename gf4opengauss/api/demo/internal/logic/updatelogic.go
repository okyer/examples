package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/okyer/examples/gf4opengauss/api/demo/internal/model"
	"github.com/okyer/examples/gf4opengauss/api/demo/internal/svc"
	"github.com/okyer/examples/gf4opengauss/api/demo/internal/types"

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

	// rs, err := l.svcCtx.DB.Update(l.ctx, "users", &user, "id", user.ID)
	rs, err := l.svcCtx.DB.Model("users").Data(user).Where("id = ?", user.ID).OmitEmpty().Update()
	if err != nil {
		return nil, err
	}

	rows, _ := rs.RowsAffected()

	resp = new(types.Response)
	resp.Message = fmt.Sprintf("success, affected rows: %d", rows)
	return
}
