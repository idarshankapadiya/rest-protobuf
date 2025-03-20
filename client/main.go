package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	pb "github.com/darshankapadiya19/rest-protobuf/proto/gen"
	"google.golang.org/protobuf/proto"
)

func sendGRPCRequest(req *pb.HelloRequest, endpoint string) (*pb.HelloResponse, error) {
	log.Println(" ==========================================================")
	log.Printf("[GRPC] Sending request to localhost:8080")
	// Encoding data using proto.Marshal using protobuf definitions
	request, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	log.Printf("rq: %s, length: %d", string(request), len(request))
	response, err := http.Post("http://localhost:8080"+endpoint, "application/x-binary", bytes.NewReader(request))
	if err != nil {
		log.Printf("Error sending request to localhost:8080: %v", err)
		return nil, err
	}
	log.Printf("Received response from localhost:8080")

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	resp := &pb.HelloResponse{}

	// Decoding data using proto.Unmarshal using protobuf definitions
	err = proto.Unmarshal(responseBytes, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func sendJsonRequest(req *pb.HelloRequest, endpoint string) (*pb.HelloResponse, error) {
	log.Println(" ==========================================================")
	log.Printf("[Json] Sending request to localhost:8080")

	// Encoding data usting json.Marshal
	request, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	log.Printf("rq: %s, length: %d", string(request), len(request))
	response, err := http.Post("http://localhost:8080"+endpoint, "application/json", bytes.NewReader(request))
	if err != nil {
		log.Printf("Error sending request to localhost:8080: %v", err)
		return nil, err
	}
	log.Printf("Received response from localhost:8080")

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	resp := &pb.HelloResponse{}

	// Decoding data using json.Unmarshal
	err = json.Unmarshal(responseBytes, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func main() {
	request := &pb.HelloRequest{Name: "Darshan"}

	helloGRPCRequest(request)
	helloJSONRequest(request)
	// haloGRPCRequest(request)
}

func helloGRPCRequest(request *pb.HelloRequest) {
	response, err := sendGRPCRequest(request, "/grpc_hello")
	if err != nil {
		log.Fatalf("Error sending request: %s", err.Error())
	}
	log.Printf("Response from server: %v", response)
}

func helloJSONRequest(request *pb.HelloRequest) {
	response, err := sendJsonRequest(request, "/json_hello")
	if err != nil {
		log.Fatalf("Error sending request: %s", err.Error())
	}
	log.Printf("Response from server: %v", response)
}

func haloGRPCRequest(request *pb.HelloRequest) {
	response, err := sendGRPCRequest(request, "/halo")
	if err != nil {
		log.Fatalf("Error sending request: %s", err.Error())
	}
	log.Printf("Response from server: %v", response)
}
