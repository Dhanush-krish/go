package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// service interface
//service.go
type StringService interface {
	Length(string) int
}

//service implementation
//logic.go
type stringService struct{}

func (stringService) Length(str string) int {
	return len(str)
}

// rpc functions
//reqresp.go
type lengthRequest struct {
	S string `json:"s"`
}

type lengthResponse struct {
	V int `json:"v"`
}

//endpoints
// endpoints.go
func makeLengthEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(lengthRequest)
		v := svc.Length(req.S)
		return lengthResponse{v}, nil
	}
}

func main() {
	fmt.Println("Hello go kit")
	svc := stringService{}

	lengthHandler := httptransport.NewServer(
		makeLengthEndpoint(svc),
		decodeLengthRequest,
		encodeLengthRequest,
	)

	http.Handle("/count", lengthHandler)
	log.Fatal(http.ListenAndServe(":9090", nil))

}

func decodeLengthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request lengthRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil

}

func encodeLengthRequest(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
