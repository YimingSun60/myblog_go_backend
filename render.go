package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {

	parsedTemplate, err1 := template.ParseFiles(globalTemplatePath + tmpl)
	if err1 != nil {
		panic(fmt.Errorf("error parsing template: %w", err1))
		return
	}
	fmt.Println("The Name is", parsedTemplate.Name())
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		panic(fmt.Errorf("error parsing template: %w", err))
		return
	}
}
