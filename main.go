package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"
const assetsPath = "/assets/"

func main() {
	fmt.Println("Start the service")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.Handle(assetsPath, http.StripPrefix(assetsPath, http.FileServer(http.Dir("templates/assets"))))
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
		return
	}

}
