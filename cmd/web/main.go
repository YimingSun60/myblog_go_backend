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

const cssPath = "/css/"

func main() {

	config.SetTemplatePath()
	var a config.AppConfig
	var repo *handlers.Repository

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Error when create template cache: ", err)
		return
	}
	a.TemplateCache = tc
	//True for production, false for development
	a.UseCache = false
	repo = handlers.NewRepo(&a)
	handlers.NewHandlers(repo)

	render.NewTemplates(&a)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/album", handlers.Repo.Album)

	http.Handle(assetsPath, http.StripPrefix(assetsPath, http.FileServer(http.Dir(config.TemplatePath+"/assets"))))

	http.Handle(cssPath, http.StripPrefix(cssPath, http.FileServer(http.Dir(config.TemplatePath+"/css"))))
	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
		return
	}
	fmt.Println("End the service")

}
