package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			flag := true
			fmt.Println("Cheking Auth")
			if flag {
				hf(w, req)
			} else {
				return
			}
		}
	}
}

func Logging() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(req.URL.Path, time.Since(start))
			}()
			hf(w, req)
		}
	}
}
