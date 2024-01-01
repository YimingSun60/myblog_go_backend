package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8081"
const assetsPath = "/assets/"

func main() {
	initialize()
	fmt.Println("Start the service")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.Handle(assetsPath, http.StripPrefix(assetsPath, http.FileServer(http.Dir(globalTemplatePath+"/assets"))))
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
		return
	}

}
