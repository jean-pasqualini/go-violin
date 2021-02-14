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

	//populate the default ScaleOptions, PitchOptions, KeyOptions, OctaveOptions for scales and arpeggios
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

	// Populate the default ScaleOptions, PitchOptions, KeyOptions, OctaveOptions for scales and arpeggios
	sOptions, pOptions, kOptions, oOptions := render.SetDefaultScaleOptions()

	request.ParseForm() //r is url.Values which is a map[string][]string

	var svalues []string
	for _, values := range request.Form { // range over map
		for _, value := range values { // range over []string
			svalues = append(svalues, value) // stick each value in a slice I know the name of
		}
	}

	scalearp, key, pitch, octave, leftlabel, rightlabel := "", "", "", "", "", ""

	// the slice of values return by the request can be arranged in any order
	// so identify selected scale / arpeggio, pitch, key and octave and store values in variables for later use.
	for i := 0; i < 4; i++ {
		switch svalues[i] {
		case "Major":
			pitch = svalues[i]
		case "Minor":
			pitch = svalues[i]
		case "1":
			octave = svalues[i]
		case "2":
			octave = svalues[i]
		case "Scale":
			scalearp = svalues[i]
		case "Arpeggio":
			scalearp = svalues[i]
		default:
			key = svalues[i]
		}
	}

	// Update options based on the user's selection

	// Set key options - set isChecked true for selected key and false for all other keys
	kOptions = render.SetKeyOptions(key)

	// Set scale options
	if scalearp == "Scale" {
		// if scale is selected set scale isChecked to true and arpeggio isChecked to false
		sOptions = []render.ScaleOptions{
			{"Scalearp", "Scale", false, true, "Scales"},
			{"Scalearp", "Arpeggio", false, false, "Arpeggios"},
		}
	} else {
		// if arpeggio is selected set arpeggio isChecked to true and scale isChecked to false
		sOptions = []render.ScaleOptions{
			{"Scalearp", "Scale", false, false, "Scales"},
			{"Scalearp", "Arpeggio", false, true, "Arpeggios"},
		}
	}

	// Set pitch options
	if pitch == "Major" {
		pOptions = []render.ScaleOptions{ // if major was selected, set major isChecked to true and minor isChecked to false
			{"Pitch", "Major", false, true, "Major"},
			{"Pitch", "Minor", false, false, "Minor"},
		}
	} else {
		pOptions = []render.ScaleOptions{ // if minor was selected, set minor isChecked to true and major isChecked to false
			{"Pitch", "Major", false, false, "Major"},
			{"Pitch", "Minor", false, true, "Minor"},
		}
	}

	// Set octave options
	if octave == "1" {
		oOptions = []render.ScaleOptions{
			{"Octave", "1", false, true, "1 Octave"},
			{"Octave", "2", false, false, "2 Octave"},
		}
	} else {
		oOptions = []render.ScaleOptions{
			{"Octave", "1", false, false, "1 Octave"},
			{"Octave", "2", false, true, "2 Octave"},
		}
	}

	// work out what the actual key is and set its value
	if pitch == "Major" {
		// for major scales if the key is longer than 2 characters, we only care about the last 2 characters
		if len(key) > 2 { // only select last two characters for keys which contain two possible names e.g. C#/Db
			key = key[3:]
		}
	} else { // pitch is minor
		// for minor scales if the key is longer than 2 characters, we only care about the first 2 characters
		if len(key) > 2 { // only select first two characters for keys which contain two possible names e.g. C#/Db
			key = key[:2]
		}
	}

	// Set the labels, Major have a scale and a drone, while minor have melodic and harmonic minor scales
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

	audioPath += strings.ToLower(key)
	imgPath += strings.ToLower(key)
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
		audioPath2 += strings.ToLower(key)
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
		Key:           key,
		Pitch:         pitch,
		ScaleImgPath:  imgPath,
		GifPath:       "img/major/gif/a1.gif",
		AudioPath:     audioPath,
		AudioPath2:    audioPath2,
		LeftLabel:     leftlabel,
		RightLabel:    rightlabel,
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

// DuetGET handles GET calls for the duet page
func (b *Base) DuetGET(responseWriter http.ResponseWriter, request *http.Request) {
	b.log.Printf("%s %s -> %s", request.Method, request.URL.Path, request.RemoteAddr)

	dOptions := []render.ScaleOptions{
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
	dOptions := []render.ScaleOptions{
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
		dOptions = []render.ScaleOptions{
			{"Duet", "G Major", false, false, "G Major"},
			{"Duet", "D Major", false, true, "D Major"},
			{"Duet", "A Major", false, false, "A Major"},
		}
		DuetImgPath = "img/duet/dmajor.png"
		DuetAudioBoth = "mp3/duet/dmajorduetboth.mp3"
		DuetAudio1 = "mp3/duet/dmajorduetpt1.mp3"
		DuetAudio2 = "mp3/duet/dmajorduetpt2.mp3"
	case "G Major":
		dOptions = []render.ScaleOptions{
			{"Duet", "G Major", false, true, "G Major"},
			{"Duet", "D Major", false, false, "D Major"},
			{"Duet", "A Major", false, false, "A Major"},
		}
		DuetImgPath = "img/duet/gmajor.png"
		DuetAudioBoth = "mp3/duet/gmajorduetboth.mp3"
		DuetAudio1 = "mp3/duet/gmajorduetpt1.mp3"
		DuetAudio2 = "mp3/duet/gmajorduetpt2.mp3"

	case "A Major":
		dOptions = []render.ScaleOptions{
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