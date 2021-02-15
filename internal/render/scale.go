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

type Display struct {
	Picture string
	AudioLeft Listen
	AudioRight Listen
	m string
}

type Listen struct {
	Label string
	Source string
}

func GenerateDisplay(scalearp string, pitch string, displayKey string, octave string) Display {
	switch scalearp + " " + pitch {
		case "Scale Major":
			return Display{
				Picture: "img/scale/major/" + strings.ToLower(displayKey) + octave + ".png",
				AudioLeft: Listen{
					Label: "Listen to Major Scale",
					Source: "mp3/scale/major/" + strings.ToLower(displayKey) + octave + ".mp3",
				},
				AudioRight: Listen{
					Label: "Listen to Drone",
					Source: "mp3/drone/" + strings.ToLower(displayKey) + octave + ".mp3",
				},
			}
		case "Scale Minor":
			return Display{
				Picture: "img/scale/minor/" + strings.ToLower(displayKey) + octave + ".png",
				AudioLeft: Listen{
					Label: "Listen to Harmonic Minor Scale",
					Source: "mp3/scale/minor/" + strings.ToLower(displayKey) + octave + ".mp3",
				},
				AudioRight: Listen{
					Label: "Listen to Melodic Minor Scale",
					Source: "mp3/scale/minor/" + strings.ToLower(displayKey) + octave + "m.mp3",
				},
			}
		case "Arpeggio Major":
			return Display{
				Picture: "img/arps/major/" + strings.ToLower(displayKey) + octave + ".png",
				AudioLeft: Listen{
					Label: "Listen to Major Arpeggio",
					Source: "img/arps/major/" + strings.ToLower(displayKey) + octave + ".mp3",
				},
				AudioRight: Listen{
					Label: "Listen to Drone",
					Source: "mp3/drone/" + strings.ToLower(displayKey) + octave + ".mp3",
				},
			}
		case "Arpeggio Minor":
			return Display{
				Picture: "img/arps/minor/" + strings.ToLower(displayKey) + octave + ".png",
				AudioLeft: Listen{
					Label: "Listen to Minor Arpeggio",
					Source: "mp3/arps/minor/" + strings.ToLower(displayKey) + octave + ".mp3",
				},
				AudioRight: Listen{
					Label: "Listen to Drone",
					Source: "mp3/drone/" + strings.ToLower(displayKey) + octave + ".mp3",
				},
			}
		default:
			return Display{}
	}
}

func ChooseKey(pitch string, key string) string {
	key = strings.Replace(key, "#", "s", 1)
	keys := strings.Split(key, "/")
	if pitch == "Minor" || len(keys) == 1 {
		return keys[0]
	}
	return keys[1]
}

