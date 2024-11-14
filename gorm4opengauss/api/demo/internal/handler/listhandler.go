package handler

import (
	"net/http"

	"github.com/okyer/examples/gorm4opengauss/api/demo/internal/logic"
	"github.com/okyer/examples/gorm4opengauss/api/demo/internal/svc"
	"github.com/okyer/examples/gorm4opengauss/api/demo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func listHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListUser
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
