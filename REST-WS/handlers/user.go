package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/JalMurH/DockerDevDeploy/rest-ws/models"
	"github.com/JalMurH/DockerDevDeploy/rest-ws/repository"
	"github.com/JalMurH/DockerDevDeploy/rest-ws/server"
	"github.com/segmentio/ksuid"
)

type SingUpReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SingUpResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func SingUpHandler(server server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req = SingUpReq{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var user = models.User{
			Email:    req.Email,
			Password: req.Password,
			Id:       id.String(),
		}
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SingUpResponse{
			Id:    user.Id,
			Email: user.Email,
		})

	}
}
