package controllers

import (
	"github.com/jean-pasqualini/goviolin/internal/render"
	"log"
	"net/http"
)

type HomeController struct {
	Log *log.Logger
}

// Handler for / renders the home.html
func (b *HomeController) Home(responseWriter http.ResponseWriter, request *http.Request) {
	b.Log.Printf("%s %s -> %s", request.Method, request.URL.Path, request.RemoteAddr)

	pageVars := render.PageVars{
		Title: "GoViolin",
	}

	if err := render.Render(responseWriter, "home.html", pageVars); err != nil {
		b.Log.Printf("%s %s -> %s : ERROR : %v", request.Method, request.URL.Path, request.RemoteAddr, err)
		return
	}
}
