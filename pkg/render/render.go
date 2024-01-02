package render

import (
	"bytes"
	"html/template"
	"log"
	"myblog"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)

	if err != nil {
		log.Fatal(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	//get all page template
	pages, err := filepath.Glob(go_backend.GlobalTemplatePath + "*.page.gohtml")
	if err != nil {
		return myCache, err
	}
	// range through pages
	for _, page := range pages {
		//get file name
		name := filepath.Base(page)
		//create template set
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		//add template set to cache
		matches, err := filepath.Glob(go_backend.GlobalTemplatePath + "*.layout.gohtml")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(go_backend.GlobalTemplatePath + "*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
