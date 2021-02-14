package handlers

import (
	"github.com/jean-pasqualini/goviolin/internal/render"
	"log"
	"net/http"
)

// Base represents the base handlers
type Base struct {
	log *log.Logger
}

// Handler for / renders the home.html
func (b *Base) Home(w http.ResponseWriter, req *http.Request) {
	pageVars := render.PageVars{
		Title: "GoViolin",
	}

	if err := render.Render(w, "home.html", pageVars); err != nil {
		b.log.Println(err)
		return
	}
}