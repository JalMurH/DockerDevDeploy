package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/JalMurH/DockerDevDeploy/rest-ws/handlers"
	"github.com/JalMurH/DockerDevDeploy/rest-ws/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	JWT_Secret := os.Getenv("JWT_Secert")
	DBURL := os.Getenv("DBURL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:      PORT,
		JWTSecret: JWT_Secret,
		DBURL:     DBURL,
	})
	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)

}
