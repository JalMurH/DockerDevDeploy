package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	db "github.com/JalMurH/DockerDevDeploy/rest-ws/DB"
	"github.com/JalMurH/DockerDevDeploy/rest-ws/repository"
	websockets "github.com/JalMurH/DockerDevDeploy/rest-ws/webSockets"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Config struct { //se crea la estructura de la configuracion del server
	Port      string
	JWTSecret string
	DBURL     string
}

type Server interface { //al crear una interface Server garantizamos que las funciones que hagan uso de la funcion config seran consideradas Server gracias a la herencia que go meneja de manera implicita
	Config() *Config
	Hub() *websockets.Hub
}

type Broker struct { //la estructura brocker sera una instancia de Server debido a que usa Config y debera tener un router que hace parte de la libreria gorilla mux
	config *Config
	router *mux.Router
	hub    *websockets.Hub
}

func (b *Broker) Config() *Config { // resiver function que es Config y asi brocker es considerado herencia de server
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) { //validaciones de configuracion encesarias para crear un servidor
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("secret Key is required")
	}
	if config.DBURL == "" {
		return nil, errors.New("data Base URL is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
		hub:    websockets.NewHub(),
	}

	return broker, nil
}

func (b *Broker) Hub() *websockets.Hub {
	return b.hub
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) { // configuracion del souter para acceder a la base de datos
	b.router = mux.NewRouter()
	binder(b, b.router)
	handler := cors.Default().Handler(b.router)
	repo, err := db.NewPostgresRepo(b.config.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	go b.hub.Run()
	repository.SetRepo(repo)
	log.Println("Staring server on port:", b.Config().Port)
	if err := http.ListenAndServe(b.config.Port, handler); err != nil {
		log.Fatal("Exit ListenAndServe", err)
	}
}

func NewPostgresRepo() {
	panic("unimplemented")
}
