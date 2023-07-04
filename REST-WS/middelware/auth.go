package middelware

import (
	"net/http"
	"strings"

	"github.com/JalMurH/DockerDevDeploy/rest-ws/models"
	"github.com/JalMurH/DockerDevDeploy/rest-ws/server"
	"github.com/golang-jwt/jwt"
)

var (
	NoAuthNeeded = []string{
		"login",
		"singup",
	}
)

func shouldCheckToken(route string) bool {
	for _, p := range NoAuthNeeded {
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

func CheckAuthMiddleWare(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !shouldCheckToken(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}
			tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
			_, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTSecret), nil
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
