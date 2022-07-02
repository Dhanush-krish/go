package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/v1/template/types", GetTemplateTypes)

	//start http server
	err := http.ListenAndServe("127.0.0.1:9090", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Println("error in starting the server!!")
		os.Exit(1)
	}

}

func GetTemplateTypes(response http.ResponseWriter, request *http.Request) {

	response.Header().Add("Content-Type", "application/json")

	filepath := "./default_template.json"
	databyte, err := ioutil.ReadFile(filepath)

	if err != nil {
		panic(err)
	}

	io.WriteString(response, string(databyte))

}
