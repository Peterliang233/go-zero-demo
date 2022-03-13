package handler

import (
	"net/http"

	"book/service/search/api/internal/logic"
	"book/service/search/api/internal/svc"
	"book/service/search/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func userHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegistryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.User(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
