package main

import (
	hlogger "MasteringGoProgramming/4.Object-Oriented_Patterns/Singleton/Logger"
	"fmt"
	"net/http"
)

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra web service")

	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance()
	fmt.Fprint(w, "Welcome to the Hydra software system")

	logger.Println("Received an http Get request on root url")
}
