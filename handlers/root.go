package handlers

import (
	"net/http"

	"github.com/unrolled/render"
)

// Root is the handler defined for "/"
func Root(ren *render.Render) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ren.HTML(w, http.StatusOK, "index", struct{}{}, render.HTMLOptions{Layout: "layout"})
		// ren.Render(w, http.StatusOK, "index", struct{}{}, "layout")
	})
}
