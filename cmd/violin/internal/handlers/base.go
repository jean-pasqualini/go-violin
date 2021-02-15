package handlers

import (
	"github.com/jean-pasqualini/goviolin/internal/render"
	"log"
	"net/http"
	"strings"
)

// Base represents the base handlers
type Base struct {
	log *log.Logger
}

// Handler for / renders the home.html
func (b *Base) Home(responseWriter http.ResponseWriter, request *http.Request) {
	b.log.Printf("%s %s -> %s", request.Method, request.URL.Path, request.RemoteAddr)

	pageVars := render.PageVars{
		Title: "GoViolin",
	}

	if err := render.Render(responseWriter, "home.html", pageVars); err != nil {
		b.log.Printf("%s %s -> %s : ERROR : %v", request.Method, request.URL.Path, request.RemoteAddr, err)
		return
	}
}

// ScaleShow handles get/post calls for the scale page.
func (b *Base) Scale(responseWriter http.ResponseWriter, request *http.Request) {
	b.log.Printf("%s %s -> %s", request.Method, request.URL.Path, request.RemoteAddr)

	// TODO: Write a function to handle errors and missing data
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

	scaleOptions, pitchOptions, keyOptions, octaveOptions := render.SetDefaultScaleOptions()
	keyOptions = render.BindValueToHtmlSelectOptions(key, keyOptions)
	scaleOptions = render.BindValueToHtmlSelectOptions(scalearp, scaleOptions)
	pitchOptions = render.BindValueToHtmlSelectOptions(pitch, pitchOptions)
	octaveOptions = render.BindValueToHtmlSelectOptions(octave, octaveOptions)

	displayKey := generateDisplayKey(pitch, key)

	// Set the labels, Major have a scale and a drone, while minor have melodic and harmonic minor scales
	leftAudiolabel, rightAudiolabel := render.GenerateScaleLabels(pitch, scalearp)

	// Intialise paths to the associated images and mp3s
	imgPath, leftAudioPath, rightAudioPath := render.GenerateScalePaths(scalearp, pitch, displayKey, octave)

	pageVars := render.PageVars{
		Title:         "Practice Scales and Arpeggios",
		Scalearp:      scalearp,
		Key:           displayKey,
		Pitch:         pitch,
		ScaleImgPath:  imgPath,
		GifPath:       "img/major/gif/a1.gif",
		AudioPath:     leftAudioPath,
		AudioPath2:    rightAudioPath,
		LeftLabel:     leftAudiolabel,
		RightLabel:    rightAudiolabel,
		ScaleOptions:  scaleOptions,
		PitchOptions:  pitchOptions,
		KeyOptions:    keyOptions,
		OctaveOptions: octaveOptions,
	}
	if err := render.Render(responseWriter, "scale.html", pageVars); err != nil {
		b.log.Printf("%s %s -> %s : ERROR : %v", request.Method, request.URL.Path, request.RemoteAddr, err)
		return
	}
}

func generateDisplayKey(pitch string, key string) string {
	keys := strings.Split(key, "/")
	if pitch == "Minor" || len(keys) == 1 {
		return keys[0]
	}

	return keys[1]
}

// DuetGET handles GET/POST calls for the duet page
func (b *Base) Duet(responseWriter http.ResponseWriter, request *http.Request) {
	b.log.Printf("%s %s -> %s", request.Method, request.URL.Path, request.RemoteAddr)

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
		b.log.Printf("%s %s -> %s : ERROR : %v", request.Method, request.URL.Path, request.RemoteAddr, err)
		return
	}
}


