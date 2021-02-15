package controllers

import (
	"github.com/jean-pasqualini/goviolin/internal/render"
	"log"
	"net/http"
)

type DuetController struct {
	Log *log.Logger
}

// DuetGET handles GET/POST calls for the duet page
func (b *DuetController) Duet(responseWriter http.ResponseWriter, request *http.Request) {
	b.Log.Printf("%s %s -> %s", request.Method, request.URL.Path, request.RemoteAddr)

	var duet = "G"
	if request.Method == "POST" {
		request.ParseForm()
		duet = request.Form["Duet"][0]
	}

	// define default duet options
	dOptions := render.BindValueToHtmlSelectOptions(
		duet,
		[]render.HtmlSelectOption{
			{"Duet", "g", false, true, "G Major"},
			{"Duet", "d", false, false, "D Major"},
			{"Duet", "a", false, false, "A Major"},
		},
	)

	DuetImgPath, DuetAudioBoth, DuetAudio1, DuetAudio2 := render.GenerateDuetPaths(duet)
	pageVars := render.PageVars{
		Title:         "Practice DuetGET",
		Key:           "G Major",
		DuetImgPath:   DuetImgPath,
		DuetAudioBoth: DuetAudioBoth,
		DuetAudio1:    DuetAudio1,
		DuetAudio2:    DuetAudio2,
		DuetOptions:   dOptions,
	}
	if err := render.Render(responseWriter, "duets.html", pageVars); err != nil {
		b.Log.Printf("%s %s -> %s : ERROR : %v", request.Method, request.URL.Path, request.RemoteAddr, err)
		return
	}
}


