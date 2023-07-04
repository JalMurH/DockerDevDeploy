package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/JalMurH/DockerDevDeploy/rest-ws/models"
	"github.com/JalMurH/DockerDevDeploy/rest-ws/repository"
	"github.com/JalMurH/DockerDevDeploy/rest-ws/server"
	"github.com/golang-jwt/jwt"
	"github.com/segmentio/ksuid"
)

type InserPostReq struct {
	PostContent string `json:"post_content"`
}

type PostResponse struct {
	Id          string `json:"id"`
	PostContent string `json:"post_content"`
}

func InserPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			var postReq = InserPostReq{}
			if err := json.NewDecoder(r.Body).Decode(&postReq); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			id, err := ksuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			post := models.Post{
				Id:          id.String(),
				PostContent: postReq.PostContent,
				UserId:      claims.UserId,
			}
			err = repository.InsertPost(r.Context(), &post)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(PostResponse{
				Id:          post.Id,
				PostContent: post.PostContent,
			})
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
