package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/xor0x/bookings/internal/config"
	"github.com/xor0x/bookings/internal/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// Render is a function that handles the rendering of a template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache{
		// get the template cache from the app config
		tc = app.TemplateCache
	} else{
		tc, _ = CreateTemplateCache()
	}
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	err := t.Execute(buf, td)

	if err != nil {
		log.Println(err)
	}
	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	// get all the templates from the folder
	pages, err := filepath.Glob("./templates/*.tmpl")
	if err != nil {
		return cache, err
	}
	// range through all files ending with *page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		// parse the template
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}
		cache[name] = ts
	}
	return cache, nil
}