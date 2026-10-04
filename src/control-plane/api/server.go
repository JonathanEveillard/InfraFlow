package api

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Server Health: Running!")
	if _, err := io.WriteString(w, "Server Health: Running!"); err != nil {
		log.Fatalln(err)
	}
}

func getProxmox(w http.ResponseWriter, r *http.Request) {

}

func Server() {

	http.HandleFunc("/health", getHealth)
	fmt.Printf("Infra Flow server at :3333\n")

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
