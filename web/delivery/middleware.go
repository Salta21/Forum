package delivery

import (
	"context"
	"net/http"
)

type ctx string

var userIdCtx ctx = "userId"

func (h *Handler) middleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("accessToken")
		if err != nil {
			next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), userIdCtx, 0)))
			return
		}
		userId, err := h.services.GetUserIdByToken(cookie.Value)
		if err != nil {
			next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), userIdCtx, 0)))
			return
		}
		next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), userIdCtx, userId)))
	}
}
