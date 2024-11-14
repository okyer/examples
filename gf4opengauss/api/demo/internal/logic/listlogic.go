package logic

import (
	"context"

	"github.com/okyer/examples/gf4opengauss/api/demo/internal/svc"
	"github.com/okyer/examples/gf4opengauss/api/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListUser) (resp *types.ListUserResponse, err error) {
	// todo: add your logic here and delete this line
	where := map[string]interface{}{}
	if req.Name != "" {
		where["name"] = req.Name
	}
	if req.Email != "" {
		where["email"] = req.Email
	}

	// users := make([]model.User, 0)
	users, err := l.svcCtx.DB.Model("users").Where(where).Limit(0, 10).All()
	if err != nil {
		return nil, err
	}

	listUsers := make([]*types.UpdateUser, 0)
	for _, user := range users {
		listUsers = append(listUsers, &types.UpdateUser{
			Id:   user["id"].Int(),
			Name: user["name"].String(),
			// Email:    *user.Email,
			Age: user["age"].Int(),
			// Birthday: user.Birthday.String(),
		})
	}

	resp = new(types.ListUserResponse)
	resp.Message = "success"
	resp.Data = listUsers
	return
}
