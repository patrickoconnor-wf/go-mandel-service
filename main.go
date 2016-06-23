package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/patrickoconnor-wf/go-mandel-service/handlers"
	"github.com/patrickoconnor-wf/go-mandel-service/util"
	"github.com/unrolled/render"
)

func processConfig() (util.Config, error) {
	config := util.Config{}
	err := envconfig.Process("mandelservice", &config)
	if err != nil {
		log.Printf("Configuration Error [%s] Returning Error", err)
	}
	return config, err
}

func main() {
	config, err := processConfig()
	if err != nil {
		log.Fatalf("Killing server due to configuration issue [%s]", err)
	}
	env, err := config.Env()
	if err != nil {
		log.Fatalf("Configuration Error [%s]", err)
	}
	log.Printf("%s", env)

	renderer := render.New(render.Options{Directory: "templates", Extensions: []string{".html"}, IsDevelopment: true})
	r := mux.NewRouter()

	r.Handle("/", handlers.Root(renderer)).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(r)
	hostStr := fmt.Sprintf(":%d", config.Port)
	n.Run(hostStr)
}
