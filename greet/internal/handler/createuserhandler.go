package handler

import (
	"fmt"
	"net/http"

	"github.com/JubaerHossain/go-boilerplate/greet/internal/logic"
	"github.com/JubaerHossain/go-boilerplate/greet/internal/svc"
	"github.com/JubaerHossain/go-boilerplate/greet/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		fmt.Println("Request: ", req)
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateUserLogic(r.Context(), svcCtx)
		err := l.CreateUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
