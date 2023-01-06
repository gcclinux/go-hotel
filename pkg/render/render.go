package render

import (
	"html/template"
	"log"
	"net/http"
)

// renderTemplate is where we will be rendering all templates through
func RenderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles("./templates/" + page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println("error parsing templates:", err)
		return
	}
}
