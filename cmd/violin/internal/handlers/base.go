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

// ScaleGET handler GET calls for the scale page.
func (b *Base) ScaleGET(responseWriter http.ResponseWriter, request *http.Request) {
	b.log.Printf("%s %s -> %s", request.Method, request.URL.Path, request.RemoteAddr)

	//populate the default HtmlSelectOption, PitchOptions, KeyOptions, OctaveOptions for scales and arpeggios
	sOptions, pOptions, kOptions, oOptions := render.SetDefaultScaleOptions()

	// set default page variables
	pageVars := render.PageVars{
		Title:         "Practice Scales and Arpeggios", // default scale initially displayed is A Major
		Scalearp:      "Scale",
		Pitch:         "Major",
		Key:           "A",
		ScaleImgPath:  "img/scale/major/a1.png",
		GifPath:       "",
		AudioPath:     "mp3/scale/major/a1.mp3",
		AudioPath2:    "mp3/drone/a1.mp3",
		LeftLabel:     "Listen to Major scale",
		RightLabel:    "Listen to Drone",
		ScaleOptions:  sOptions,
		PitchOptions:  pOptions,
		KeyOptions:    kOptions,
		OctaveOptions: oOptions,
	}

	if err := render.Render(responseWriter, "scale.html", pageVars); err != nil {
		b.log.Printf("%s %s -> %s : ERROR : %v", request.Method, request.URL.Path, request.RemoteAddr, err)
		return
	}
}

// ScaleShow handles post calls for the scale page.
func (b *Base) ScalePOST(responseWriter http.ResponseWriter, request *http.Request) {
	b.log.Printf("%s %s -> %s", request.Method, request.URL.Path, request.RemoteAddr)

	// TODO: Write a function to handle errors and missing data
	request.ParseForm()
	pitch := request.Form["Pitch"][0]
	octave := request.Form["Octave"][0]
	scalearp := request.Form["Scalearp"][0]
	key := request.Form["Key"][0]

	scaleOptions, pitchOptions, keyOptions, octaveOptions := render.SetDefaultScaleOptions()
	keyOptions = render.BindValueToHtmlSelectOptions(key, keyOptions)
	scaleOptions = render.BindValueToHtmlSelectOptions(scalearp, scaleOptions)
	pitchOptions = render.BindValueToHtmlSelectOptions(pitch, pitchOptions)
	octaveOptions = render.BindValueToHtmlSelectOptions(octave, octaveOptions)

	displayKey := func(pitch string, key string) string {
		keys := strings.Split(key, "/")
		if pitch == "Minor" || len(keys) == 0 {
			return keys[0]
		}

		return keys[1]
	}(pitch, key)

	// Set the labels, Major have a scale and a drone, while minor have melodic and harmonic minor scales
	leftlabel, rightlabel := "", ""
	if pitch == "Major" {
		leftlabel = "Listen to Major "
		rightlabel = "Listen to Drone"
		if scalearp == "Scale" {
			leftlabel += "Scale"
		} else {
			leftlabel += "Arpeggio"
		}
	} else {
		if scalearp == "Arpeggio" {
			leftlabel += "Listen to Minor Arpeggio"
			rightlabel = "Listen to Drone"
		} else {
			leftlabel += "Listen to Harmonic Minor Scale"
			rightlabel += "Listen to Melodic Minor Scale"
		}
	}

	// Intialise paths to the associated images and mp3s
	imgPath, audioPath, audioPath2 := "img/", "mp3/", "mp3/"

	// Build paths to img and mp3 files that correspond to user selection
	if scalearp == "Scale" {
		imgPath += "scale/"
		audioPath += "scale/"

	} else {
		// if arpeggio is selected, add "arps/" to the img and mp3 paths
		imgPath += "arps/"
		audioPath += "arps/"
	}

	if pitch == "Major" {
		imgPath += "major/"
		audioPath += "major/"
	} else {
		imgPath += "minor/"
		audioPath += "minor/"
	}

	audioPath += strings.ToLower(displayKey)
	imgPath += strings.ToLower(displayKey)
	// if the img or audio path contain #, delete last character and replace it with s
	imgPath = render.ChangeSharpToS(imgPath)
	audioPath = render.ChangeSharpToS(audioPath)

	switch octave {
	case "1":
		imgPath += "1"
		audioPath += "1"
	case "2":
		imgPath += "2"
		audioPath += "2"
	}

	audioPath += ".mp3"
	imgPath += ".png"

	//generate audioPath2
	// audio path2 can either be a melodic minor scale or a drone note.
	// Set to melodic minor scale - if the first 16 characters of audio path are:
	if audioPath[:16] == "mp3/scale/minor/" {
		audioPath2 = audioPath                      // set audioPath2 to the original audioPath
		audioPath2 = audioPath2[:len(audioPath2)-4] // chop off the last 4 characters, this removes .mp3
		audioPath2 += "m.mp3"                       // then add m for melodic and the .mp3 suffix
	} else { // audioPath2 needs to be a drone note.
		audioPath2 += "drone/"
		audioPath2 += strings.ToLower(displayKey)
		// may have just added a # to the path, so use the function to change # to s
		audioPath2 = render.ChangeSharpToS(audioPath2)
		switch octave {
		case "1":
			audioPath2 += "1.mp3"
		case "2":
			audioPath2 += "2.mp3"
		}
	}

	pageVars := render.PageVars{
		Title:         "Practice Scales and Arpeggios",
		Scalearp:      scalearp,
		Key:           displayKey,
		Pitch:         pitch,
		ScaleImgPath:  imgPath,
		GifPath:       "img/major/gif/a1.gif",
		AudioPath:     audioPath,
		AudioPath2:    audioPath2,
		LeftLabel:     leftlabel,
		RightLabel:    rightlabel,
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

// DuetGET handles GET calls for the duet page
func (b *Base) DuetGET(responseWriter http.ResponseWriter, request *http.Request) {
	b.log.Printf("%s %s -> %s", request.Method, request.URL.Path, request.RemoteAddr)

	dOptions := []render.HtmlSelectOption{
		{"Duet", "G Major", false, true, "G Major"},
		{"Duet", "D Major", false, false, "D Major"},
		{"Duet", "A Major", false, false, "A Major"},
	}

	pageVars := render.PageVars{
		Title:         "Practice DuetGET",
		Key:           "G Major",
		DuetImgPath:   "img/duet/gmajor.png",
		DuetAudioBoth: "mp3/duet/gmajorduetboth.mp3",
		DuetAudio1:    "mp3/duet/gmajorduetpt1.mp3",
		DuetAudio2:    "mp3/duet/gmajorduetpt2.mp3",
		DuetOptions:   dOptions,
	}
	if err := render.Render(responseWriter, "duets.html", pageVars); err != nil {
		b.log.Printf("%s %s -> %s : ERROR : %v", request.Method, request.URL.Path, request.RemoteAddr, err)
		return
	}
}

// DuetGET handles POST calls for the duet page
func (b *Base) DuetPOST(responseWriter http.ResponseWriter, request *http.Request) {
	b.log.Printf("%s %s -> %s", request.Method, request.URL.Path, request.RemoteAddr)

	// define default duet options
	dOptions := []render.HtmlSelectOption{
		{"Duet", "G Major", false, true, "G Major"},
		{"Duet", "D Major", false, false, "D Major"},
		{"Duet", "A Major", false, false, "A Major"},
	}

	// Set a placeholder image path, this will be changed later.
	DuetImgPath := "img/duet/gmajor.png"
	DuetAudioBoth := "mp3/duet/gmajorduetboth.mp3"
	DuetAudio1 := "mp3/duet/gmajorduetpt1"
	DuetAudio2 := "mp3/duet/gmajorduetpt2"

	request.ParseForm() //r is url.Values which is a map[string][]string
	var dvalues []string
	for _, values := range request.Form { // range over map
		for _, value := range values { // range over []string
			dvalues = append(dvalues, value) // stick each value in a slice I know the name of
		}
	}

	switch dvalues[0] {
	case "D Major":
		dOptions = []render.HtmlSelectOption{
			{"Duet", "G Major", false, false, "G Major"},
			{"Duet", "D Major", false, true, "D Major"},
			{"Duet", "A Major", false, false, "A Major"},
		}
		DuetImgPath = "img/duet/dmajor.png"
		DuetAudioBoth = "mp3/duet/dmajorduetboth.mp3"
		DuetAudio1 = "mp3/duet/dmajorduetpt1.mp3"
		DuetAudio2 = "mp3/duet/dmajorduetpt2.mp3"
	case "G Major":
		dOptions = []render.HtmlSelectOption{
			{"Duet", "G Major", false, true, "G Major"},
			{"Duet", "D Major", false, false, "D Major"},
			{"Duet", "A Major", false, false, "A Major"},
		}
		DuetImgPath = "img/duet/gmajor.png"
		DuetAudioBoth = "mp3/duet/gmajorduetboth.mp3"
		DuetAudio1 = "mp3/duet/gmajorduetpt1.mp3"
		DuetAudio2 = "mp3/duet/gmajorduetpt2.mp3"

	case "A Major":
		dOptions = []render.HtmlSelectOption{
			{"Duet", "G Major", false, false, "G Major"},
			{"Duet", "D Major", false, false, "D Major"},
			{"Duet", "A Major", false, true, "A Major"},
		}
		DuetImgPath = "img/duet/amajor.png"
		DuetAudioBoth = "mp3/duet/amajorduetboth.mp3"
		DuetAudio1 = "mp3/duet/amajorduetpt1.mp3"
		DuetAudio2 = "mp3/duet/amajorduetpt2.mp3"
	}

	//	imgPath := "img/"

	// set default page variables
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
