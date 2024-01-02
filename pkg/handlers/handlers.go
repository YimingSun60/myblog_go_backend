package handlers

import (
	"myblog/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "/home.page.gohtml")
}

func Album(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "/album.page.gohtml")
}
