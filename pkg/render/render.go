package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/rohanv8676/bookings/pkg/config"
	"github.com/rohanv8676/bookings/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package.
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(templateData *models.TemplateData) *models.TemplateData {
	return templateData
}

// Renders templates using html templates.
func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Create a template cache, by getting the template cache  from app.config
	// templateCache := app.TemplateCache

	// templateCache, err := CreateTemplateCache()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Get requested template from cache
	templatePointer, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	templateData = AddDefaultData(templateData)

	err := templatePointer.Execute(buf, templateData)
	if err != nil {
		log.Println(err)
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	// parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")
	// err := parsedTemplate.Execute(w, nil)

	// if err != nil {
	// 	fmt.Println("error parsing template:", err)
	// 	return
	// }
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// Get all of the filesnames *.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// Range through all files ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}

	return myCache, nil
}

// // cache of templates
// var templateCache = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	// check to see if we already have the template in our cache
// 	_, inMap := templateCache[t]

// 	if !inMap {
// 		//need to create the template
// 		log.Println("Creating template and adding to cache.")
// 		err = createTemplateCache(t)

// 		if err != nil {
// 			log.Println(err)
// 		}

// 	} else {
// 		// we have the template in cache.
// 		log.Println("Using cached template.")
// 	}

// 	tmpl = templateCache[t]
// 	err = tmpl.Execute(w, nil)

// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.html",
// 	}

// 	// parse the template
// 	tmpl, err := template.ParseFiles(templates...)

// 	if err != nil {
// 		return err
// 	}

// 	// add template to map
// 	templateCache[t] = tmpl

// 	return nil

// }
