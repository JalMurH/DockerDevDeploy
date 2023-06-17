package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HaldleRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Ruta raiz")
}
func HandleHome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Home end-point")
}

func PostRequest(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var metaData MetaData
	err := decoder.Decode(&metaData)
	if err != nil {
		fmt.Fprintf(w, "Error %v\n", err)
		return
	}
	fmt.Fprintf(w, "PayLoad %v\n", metaData)
}

func UserPostRequest(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error %v\n", err)
		return
	}
	response, err := user.ToJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Conten-Type", "aplications/json")
	w.Write(response)
}
