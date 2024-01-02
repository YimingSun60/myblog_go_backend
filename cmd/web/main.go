package main

import (
	"fmt"
	"myblog"
	"myblog/pkg/handlers"
	"net/http"
)

const portNumber = ":8081"
const assetsPath = "/assets/"

func main() {
	go_backend.Initialize()
	fmt.Println("Start the service")
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.Handle(assetsPath, http.StripPrefix(assetsPath, http.FileServer(http.Dir(go_backend.GlobalTemplatePath+"/assets"))))
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
		return
	}
	fmt.Println("End the service")

}
