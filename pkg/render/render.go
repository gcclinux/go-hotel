package render

// render.go - lesson 30 - simple templating cache system
import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// renderTemplate is where we will be rendering all templates through
func RenderTemplate(w http.ResponseWriter, page string) {
	// create  a template cache

	templatecache, err := createTemplateCache()
	if err != nil {
		log.Fatal("error creating createTemplateCache()", err)
		return
	}

	// get requested template from cache

	myTemplate, ok := templatecache[page]
	if !ok {
		log.Fatal(err)
	}

	myBuffer := new(bytes.Buffer)
	err = myTemplate.Execute(myBuffer, nil)
	if err != nil {
		log.Println("error comes from the map:", err)
	}

	// render the template

	_, err = myBuffer.WriteTo(w)
	if err != nil {
		log.Println("error comes from the myBuffer:", err)
	}

}

func createTemplateCache() (map[string]*template.Template, error) {
	// type 1 make mp
	// myCache := make(map[string]*template.Template)

	// type 2 mmaking map and then just use curly brackets and make it an empty map.
	myCache := map[string]*template.Template{}

	//get all of the files name *.page.tmpl from ./template
	pages, err := filepath.Glob("./templates/*.page.tmpl")
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

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateset, err = templateset.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = templateset
	}
	return myCache, nil
}
