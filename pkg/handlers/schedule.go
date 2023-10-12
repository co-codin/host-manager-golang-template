package handlers

import (
	"net/http"

	"host-manager/pkg/config"
	"host-manager/pkg/helpers"
)

// ListEntries lists schedule entries
func (repo *DBRepo) ListEntries(app config.AppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := helpers.RenderPage(w, r, "schedule", nil, nil)
		if err != nil {
			printTemplateError(w, err)
		}
	}
}
