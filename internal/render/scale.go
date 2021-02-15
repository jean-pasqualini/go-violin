package render

import "strings"

// SetDefaultScaleOptions provide the defaults for rendering scales.
func SetDefaultScaleOptions() ([]HtmlSelectOption, []HtmlSelectOption, []HtmlSelectOption, []HtmlSelectOption) {

	scaleOptions := []HtmlSelectOption{
		{Name: "Scalearp", Value: "Scale", Text: "Scales", IsChecked: true},
		{Name: "Scalearp", Value: "Arpeggio", Text: "Arpeggios"},
	}

	pitchOptions := []HtmlSelectOption{
		{Name: "Pitch", Value: "Major", Text: "Major", IsChecked: true},
		{Name: "Pitch", Value: "Minor", Text: "Minor"},
	}

	keyOptions := []HtmlSelectOption{
		{Name: "Key", Value: "A", Text: "A", IsChecked: true},
		{Name: "Key", Value: "Bb", Text: "Bb"},
		{Name: "Key", Value: "B", Text: "B"},
		{Name: "Key", Value: "C", Text: "C"},
		{Name: "Key", Value: "C#/Db", Text: "C#/Db"},
		{Name: "Key", Value: "D", Text: "D"},
		{Name: "Key", Value: "Eb", Text: "Eb"},
		{Name: "Key", Value: "E", Text: "E"},
		{Name: "Key", Value: "F", Text: "F"},
		{Name: "Key", Value: "F#/Gb", Text: "F#/Gb"},
		{Name: "Key", Value: "G", Text: "G"},
		{Name: "Key", Value: "G#/Ab", Text: "G#/Ab"},
	}

	octaveOptions := []HtmlSelectOption{
		{Name: "Octave", Value: "1", Text: "1 Octave", IsChecked: true},
		{Name: "Octave", Value: "2", Text: "2 Octave"},
	}

	return scaleOptions, pitchOptions, keyOptions, octaveOptions
}

// ChangeSharpToS WE DON'T KNOW WHY YET.
func ChangeSharpToS(path string) string {
	if strings.Contains(path, "#") {
		path = path[:len(path)-1]
		path += "s"
	}
	return path
}

func GenerateScalePaths(scalearp string, pitch string, displayKey string, octave string) (string, string, string) {
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
	imgPath = ChangeSharpToS(imgPath)
	audioPath = ChangeSharpToS(audioPath)

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
		audioPath2 = ChangeSharpToS(audioPath2)
		switch octave {
		case "1":
			audioPath2 += "1.mp3"
		case "2":
			audioPath2 += "2.mp3"
		}
	}

	return imgPath, audioPath, audioPath2
}

func GenerateScaleLabels(pitch string, scalearp string) (string, string) {
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

	return leftlabel, rightlabel
}

