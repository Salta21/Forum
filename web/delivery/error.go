package delivery

import (
	"html/template"
	"net/http"

	"forum/web/models"
)

func (h *Handler) errorPage(writer http.ResponseWriter, request *http.Request, code int) {
	userId := request.Context().Value(userIdCtx).(int)

	e := models.ErrorPage{
		Message: http.StatusText(code),
		Code:    code,
		User:    userId,
	}

	path := "./web/templates/error.html"
	t, err := template.ParseFiles(path)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), 500)
		return
	}
	writer.WriteHeader(code)
	err = t.Execute(writer, e)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), 500)
		return
	}
}
