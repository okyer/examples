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

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.User) (resp *types.Response, err error) {
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
		Name:      req.Name,
		Email:     email,
		Age:       uint8(req.Age),
		Birthday:  birthday,
		CreatedAt: time.Now(),
	}

	// rs, err := l.svcCtx.DB.Insert(l.ctx, "users", &user)
	rs, err := l.svcCtx.DB.Model("users").Data(user).OmitEmpty().Insert()
	if err != nil {
		return nil, err
	}

	id, _ := rs.LastInsertId()

	resp = new(types.Response)
	resp.Message = fmt.Sprintf("success, insert id: %d", id)
	return
}
