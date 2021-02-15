package handlers

import (
	"github.com/jean-pasqualini/goviolin/cmd/violin/internal/handlers/controllers"
	"log"
	"net/http"
)

// NewMux constructs and mux with all route predefined.
func NewMux(log *log.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	// Serve everything in the css folder, the img folder and mp3 folder as a file.
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	mux.Handle("/mp3/", http.StripPrefix("/mp3/", http.FileServer(http.Dir("mp3"))))

	// When navigating to /home it should serve the home page.
	mux.HandleFunc("/", (&controllers.HomeController{Log: log}).Home)
	mux.HandleFunc("/scale", (&controllers.ScaleController{Log: log}).Scale)
	mux.HandleFunc("/duets", (&controllers.DuetController{Log: log}).Duet)

	return mux
}