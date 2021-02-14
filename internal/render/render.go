package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// HtmlSelectOption represents the options for generating the content.
type HtmlSelectOption struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

// PageVars represents the input for generating a web page.
type PageVars struct {
	Title         string
	Scalearp      string
	Key           string
	Pitch         string
	DuetImgPath   string
	ScaleImgPath  string
	GifPath       string
	AudioPath     string
	AudioPath2    string
	DuetAudioBoth string
	DuetAudio1    string
	DuetAudio2    string
	LeftLabel     string
	RightLabel    string
	ScaleOptions  []HtmlSelectOption
	DuetOptions   []HtmlSelectOption
	PitchOptions  []HtmlSelectOption
	KeyOptions    []HtmlSelectOption
	OctaveOptions []HtmlSelectOption
}

func BindValueToHtmlSelectOptions(value string, options []HtmlSelectOption) []HtmlSelectOption {
	for key, option := range options {
		if value == option.Value {
			options[key].IsChecked = true
		}
	}

	return options
}

// Render genereates the html for any given web page.
func Render(w http.ResponseWriter, tmpl string, pageVars PageVars) error {

	// Prefix the name passed in with templates.
	tmpl = fmt.Sprintf("templates/%s", tmpl)
	t, err := template.ParseFiles(tmpl)

	// Parse the template file held in the templates folder.
	if err != nil {
		return err
	}

	// execute the template and pass in the variables to fill the gaps.
	if err := t.Execute(w, pageVars); err != nil {
		return err
	}

	return nil
}