package delivery

import (
	"fmt"
	"html/template"
	"net/http"
)

func (h *Handler) render(w http.ResponseWriter, filename string, data interface{}) error {
	path := fmt.Sprintf("./web/templates/%s", filename)
	t, err := template.ParseFiles(path)
	if err != nil {
		return fmt.Errorf("in parsefiles function: %w", err)
	}
	err = t.Execute(w, data)
	if err != nil {
		return fmt.Errorf("in execute function: %w", err)
	}
	return nil
}
