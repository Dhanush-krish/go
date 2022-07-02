package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/index", serveIndex)
	mux.HandleFunc("/", serveRoot)

	//start http server
	err := http.ListenAndServe("127.0.0.1:9090", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Println("error in starting the server!!")
		os.Exit(1)
	}

}

func serveRoot(response http.ResponseWriter, request *http.Request) {

	log.Println("Handling the request for root")
	io.WriteString(response, "Hey you are in root path!!!")

}

func serveIndex(response http.ResponseWriter, request *http.Request) {

	log.Println("handling the request for index")
	io.WriteString(response, "Don't look for a nice index page coz, \n\nI'm a backend developer :)")

}
