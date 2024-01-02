package go_backend

import (
	"fmt"
	"path/filepath"
)

var GlobalTemplatePath string

func Initialize() {
	absPath, absPath_err := filepath.Abs("../templates")
	if absPath_err != nil {
		fmt.Println("Error when get abs path: ", absPath_err)
		return
	}
	GlobalTemplatePath = absPath
	fmt.Println(GlobalTemplatePath)
}
