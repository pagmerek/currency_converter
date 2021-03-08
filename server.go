package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/js/inputFieldChange.js", handleJs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func handleJs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "js/inputFieldChange.js")
}
func index(w http.ResponseWriter, r *http.Request) {
	rate := struct {
		Value float32
	}{
		Value: getExchangeRate("gbp"),
	}
	tpl.ExecuteTemplate(w, "index.gohtml", rate)
}
