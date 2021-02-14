package render

import "strings"

// SetDefaultScaleOptions provide the defaults for rendering scales.
func SetDefaultScaleOptions() ([]HtmlSelectOption, []HtmlSelectOption, []HtmlSelectOption, []HtmlSelectOption) {

	// Set the default scaleOptions for scales and arpeggios.
	sOptions := []HtmlSelectOption{
		{"Scalearp", "Scale", false, true, "Scales"},
		{"Scalearp", "Arpeggio", false, false, "Arpeggios"},
	}

	// Set the default PitchOptions for scales and arpeggios.
	pOptions := []HtmlSelectOption{
		{"Pitch", "Major", false, true, "Major"},
		{"Pitch", "Minor", false, false, "Minor"},
	}

	// Set the default KeyOptions for scales and arpeggios.
	kOptions := []HtmlSelectOption{
		{"Key", "A", false, true, "A"},
		{"Key", "Bb", false, false, "Bb"},
		{"Key", "B", false, false, "B"},
		{"Key", "C", false, false, "C"},
		{"Key", "C#/Db", false, false, "C#/Db"},
		{"Key", "D", false, false, "D"},
		{"Key", "Eb", false, false, "Eb"},
		{"Key", "E", false, false, "E"},
		{"Key", "F", false, false, "F"},
		{"Key", "F#/Gb", false, false, "F#/Gb"},
		{"Key", "G", false, false, "G"},
		{"Key", "G#/Ab", false, false, "G#/Ab"},
	}

	// Set the default OctaveOptions for scales and arpeggios.
	oOptions := []HtmlSelectOption{
		{"Octave", "1", false, true, "1 Octave"},
		{"Octave", "2", false, false, "2 Octave"},
	}
	return sOptions, pOptions, kOptions, oOptions
}

// ChangeSharpToS WE DON'T KNOW WHY YET.
func ChangeSharpToS(path string) string {
	if strings.Contains(path, "#") {
		path = path[:len(path)-1]
		path += "s"
	}
	return path
}
