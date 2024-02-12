// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/JubaerHossain/go-boilerplate/greet/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/users/create",
				Handler: CreateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/users/id/:userId",
				Handler: GetUserHandler(serverCtx),
			},
		},
	)
}
