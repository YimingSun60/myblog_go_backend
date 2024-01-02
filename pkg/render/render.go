package render

import (
	"fmt"
	"html/template"
	"myblog"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	//使用layout template
	parsedTemplate, err1 := template.ParseFiles(go_backend.GlobalTemplatePath+tmpl, go_backend.GlobalTemplatePath+"/base.layout.gohtml")
	if err1 != nil {
		panic(fmt.Errorf("error parsing template: %w", err1))
		return
	}
	//fmt.Println("The Name is", parsedTemplate.Name())
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		panic(fmt.Errorf("error parsing template: %w", err))
		return
	}
}
