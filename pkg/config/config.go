package config

import (
	"fmt"
	"html/template"
	"path/filepath"
)

var TemplatePath string

type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
}

func SetTemplatePath() {
	absPath, absPath_err := filepath.Abs("../templates/")
	if absPath_err != nil {
		fmt.Println("Error when get abs path: ", absPath_err)
		return
	}
	TemplatePath = absPath
	fmt.Println(TemplatePath)
}
