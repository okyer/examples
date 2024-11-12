package handler

import (
	"net/http"

	"github.com/okyer/examples/gorm4gaussdb/api/demo/internal/logic"
	"github.com/okyer/examples/gorm4gaussdb/api/demo/internal/svc"
	"github.com/okyer/examples/gorm4gaussdb/api/demo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func updateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUser
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateLogic(r.Context(), svcCtx)
		resp, err := l.Update(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
