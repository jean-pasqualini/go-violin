package render

import "fmt"

func GenerateDuetPaths(duet string) (string, string, string, string) {
	DuetImgPath := fmt.Sprintf("img/duet/%smajor.png", duet)
	DuetAudioBoth := fmt.Sprintf("mp3/duet/%smajorduetboth.mp3", duet)
	DuetAudio1 := fmt.Sprintf("mp3/duet/%smajorduetpt1.mp3", duet)
	DuetAudio2 := fmt.Sprintf("mp3/duet/%smajorduetpt2.mp3", duet)
	return DuetImgPath, DuetAudioBoth, DuetAudio1, DuetAudio2
}
