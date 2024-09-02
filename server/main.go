package main

import (
	"net/http"
	"server/utils"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/paquetes", utils.RecibirPaquetes)
	mux.HandleFunc("/mensaje", utils.RecibirMensaje)
	// aca no hice nd
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
