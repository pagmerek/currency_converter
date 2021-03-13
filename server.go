package main

import (
	"fmt"
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
	fmt.Println("Started server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func handleJs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "js/inputFieldChange.js")
}
func index(w http.ResponseWriter, r *http.Request) {
	rate := struct {
		Value   float32
		Success bool
	}{
		Value:   getExchangeRate("gbp"),
		Success: true,
	}
	fmt.Println(rate.Success, rate.Value)

	if rate.Value < 0 {
		rate.Success = false
	}
	fmt.Println(rate.Success)
	tpl.ExecuteTemplate(w, "index.gohtml", rate)
}
