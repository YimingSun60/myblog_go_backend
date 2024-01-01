package main

import (
	"fmt"
	"path/filepath"
)

var globalTemplatePath string

func initialize() {
	absPath, absPath_err := filepath.Abs("../myblog/templates")
	if absPath_err != nil {
		fmt.Println("Error when get abs path: ", absPath_err)
		return
	}
	globalTemplatePath = absPath
	fmt.Println(globalTemplatePath)
}
