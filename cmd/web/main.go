package main

import (
	"fmt"
	"log"
	"myblog/pkg/config"
	"myblog/pkg/handlers"
	"myblog/pkg/render"
	"net/http"
)

const portNumber = ":8081"
const assetsPath = "/assets/"

func main() {

	config.SetTemplatePath()
	var a config.AppConfig
	tc, err := render.CreateTemplateCache()
	fmt.Println(tc)
	if err != nil {
		log.Fatal("Error when create template cache: ", err)
		return
	}
	a.TemplateCache = tc
	render.NewTemplates(&a)

	fmt.Println("Start the service")
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/album", handlers.Album)
	http.Handle(assetsPath, http.StripPrefix(assetsPath, http.FileServer(http.Dir(config.TemplatePath+"/assets"))))
	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
		return
	}
	fmt.Println("End the service")

}
