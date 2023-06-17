package main

import (
	"fmt"
	"net/http"
)

func HaldleRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Ruta raiz")
}
func HandleHome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Home end-point")
}
