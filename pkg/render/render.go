package render

// render.go - lesson 30 - simple templating cache system
import (
	"bytes"
	"html/template"
	"log"
	"myapp/pkg/config"

	"myapp/pkg/models"
	"net/http"
	"path/filepath"
)

//var functions = template.FuncMap{}

var app *config.AppConfig

// RenderTemplate sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// renderTemplates is where we will be rendering all templates through
func RenderTemplate(w http.ResponseWriter, page string, templateData *models.TemplateData) {

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
	templateData = AddDefaultData(templateData)
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
	pages, err := filepath.Glob("./working-html/*.page.tmpl")
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

		matches, err := filepath.Glob("./working-html/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateset, err = templateset.ParseGlob("./working-html/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = templateset
	}
	return myCache, nil
}
