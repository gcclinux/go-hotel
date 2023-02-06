package render

// render.go - lesson 30 - simple templating cache system
import (
	"bytes"
	"html/template"
	"log"
	"myapp/internal/config"

	"myapp/internal/models"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
)

//var functions = template.FuncMap{}

var app *config.AppConfig

// RenderTemplate sets the config for the templ8ate package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

// renderTemplates is where we will be rendering all templates through
func RenderTemplate(w http.ResponseWriter, r *http.Request, page string, templateData *models.TemplateData) {
	var templatecache map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		templatecache = app.TemplateCache
	} else {
		templatecache, _ = CreateTemplateCache()
	}
	// get the template cache from the app config

	// get requested template from cache
	myTemplate, ok := templatecache[page]
	if !ok {
		log.Fatal("Could not get template from template cash")
	}

	myBuffer := new(bytes.Buffer)
	templateData = AddDefaultData(templateData, r)
	_ = myTemplate.Execute(myBuffer, templateData)

	// render the template
	_, err := myBuffer.WriteTo(w)
	if err != nil {
		log.Println("error writing template to browser", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// type 1 make mp
	// myCache := make(map[string]*template.Template)

	// type 2 mmaking map and then just use curly brackets and make it an empty map.
	myCache := map[string]*template.Template{}

	//get all of the files name *.page.tmpl from ./template
	pages, err := filepath.Glob("./html-source/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		// strip of the path and leave only the file name it self
		name := filepath.Base(page)

		templateset, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./html-source/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateset, err = templateset.ParseGlob("./html-source/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = templateset
	}
	return myCache, nil
}
