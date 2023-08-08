package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/JalMurH/DockerDevDeploy/rest-ws/handlers"
	"github.com/JalMurH/DockerDevDeploy/rest-ws/middelware"
	"github.com/JalMurH/DockerDevDeploy/rest-ws/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

/*
En el proyecto se desarrolla un REST API Usando websockets para mantener conecion con el cliente en tiempo real
*/
func main() {
	err := godotenv.Load(".env") //primero se cargan las variables de entorno necesarias para el funcionamiento del proyecto ubicadas en el respectivo .env ubicado en el directorio "server" y se inicializan las variables DBURL(URL de la base de datos en este caso postgresSQL), el puerto donde va a a correr el servicio, y el jwtsecret psrs ls codificacion del json web token
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT") // se leen las constantes del .env
	JWT_Secret := os.Getenv("JWT_Secert")
	DBURL := os.Getenv("DBURL")

	s, err := server.NewServer(context.Background(), &server.Config{ //se inicia un nuevo server al cual se le pasa el contexto y su configuracion
		Port:      PORT,
		JWTSecret: JWT_Secret,
		DBURL:     DBURL,
	})
	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes) // con el server creado y configurado se inicia el servicio pasando como parametro las bind rutes que son nuestros handlers o entpoints
}

func BindRoutes(s server.Server, r *mux.Router) { //nuestro metodo de routeo
	api := r.PathPrefix("/api/v1").Subrouter()                         //los que usen api seran protejidos por el middleware
	api.Use(middelware.CheckAuthMiddleWare(s))                         //para cada un de las siguientes rutas va a usar el middleware para validar el jsonwebtoken
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet) //ruta principal o home
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/singup", handlers.SingUpHandler(s)).Methods(http.MethodPost)
	api.HandleFunc("/me", handlers.MeddleWareHandler(s)).Methods(http.MethodGet)           //para obtenerla data del usuario que envia el token
	api.HandleFunc("/post", handlers.InserPostHandler(s)).Methods(http.MethodPost)         //C
	r.HandleFunc("/post/{id}", handlers.GetPostByIdHandler(s)).Methods(http.MethodGet)     //R
	api.HandleFunc("/post/{id}", handlers.UpdatePostHandler(s)).Methods(http.MethodPut)    //U
	api.HandleFunc("/post/{id}", handlers.DeletePostHandler(s)).Methods(http.MethodDelete) //D
	r.HandleFunc("/posts", handlers.ListPostHandler(s)).Methods(http.MethodGet)            //lista de los posts existentes
	r.HandleFunc("/ws", s.Hub().HandleWebSocket)
}
