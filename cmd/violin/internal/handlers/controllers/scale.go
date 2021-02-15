package controllers

import (
	"github.com/jean-pasqualini/goviolin/internal/render"
	"log"
	"net/http"
)

type ScaleController struct {
	Log *log.Logger
}

// ScaleShow handles get/post calls for the scale page.
func (b *ScaleController) Scale(responseWriter http.ResponseWriter, request *http.Request) {
	b.Log.Printf("%s %s -> %s", request.Method, request.URL.Path, request.RemoteAddr)

	scaleOptions, pitchOptions, keyOptions, octaveOptions := render.SetDefaultScaleOptions()

	// TODO: Write a function to handle errors and missing data
	pitch, octave, scalearp, key := getRequestData(request)

	keyOptions = render.BindValueToHtmlSelectOptions(key, keyOptions)
	scaleOptions = render.BindValueToHtmlSelectOptions(scalearp, scaleOptions)
	pitchOptions = render.BindValueToHtmlSelectOptions(pitch, pitchOptions)
	octaveOptions = render.BindValueToHtmlSelectOptions(octave, octaveOptions)

	display := render.GenerateDisplay(scalearp, pitch, render.ChooseKey(pitch, key), octave)

	pageVars := render.PageVars{
		Title:         "Practice Scales and Arpeggios",
		Scalearp:      scalearp,
		Pitch:         pitch,
		ScaleImgPath:  display.Picture,
		GifPath:       "img/major/gif/a1.gif",
		AudioPath:     display.AudioLeft.Source,
		AudioPath2:    display.AudioRight.Source,
		LeftLabel:     display.AudioLeft.Label,
		RightLabel:    display.AudioRight.Label,
		ScaleOptions:  scaleOptions,
		PitchOptions:  pitchOptions,
		KeyOptions:    keyOptions,
		OctaveOptions: octaveOptions,
	}
	if err := render.Render(responseWriter, "scale.html", pageVars); err != nil {
		b.Log.Printf("%s %s -> %s : ERROR : %v", request.Method, request.URL.Path, request.RemoteAddr, err)
		return
	}
}

func getRequestData(request *http.Request) (string, string, string, string) {
	var pitch = "Major"
	var octave = "1"
	var scalearp = "Scale"
	var key = "A"
	if request.Method == "POST" {
		request.ParseForm()
		pitch = request.Form["Pitch"][0]
		octave = request.Form["Octave"][0]
		scalearp = request.Form["Scalearp"][0]
		key = request.Form["Key"][0]
	}
	return pitch, octave, scalearp, key
}