package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	pb "github.com/darshankapadiya19/rest-protobuf/proto/gen"
	"github.com/gorilla/mux"
	"google.golang.org/protobuf/proto"
)

type grpcHaloHandler struct{}

// grpcHaloHandler is a struct that implements the http.Handler interface
func (h *grpcHaloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received with content of length: %d", r.ContentLength)
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
	}

	// Notice the declaration of the variable here
	// Alternatively we can do the following
	// 		req := &pb.HelloRequest{}
	// 		proto.Unmarshal(data, req)
	var req pb.HelloRequest
	err = proto.Unmarshal(data, &req)
	if err != nil {
		log.Fatalf("Unable to unmarshal message from request : %v", err)
	}
	log.Printf("Halo request form %s", req.Name)

	resp := &pb.HelloResponse{
		Message: "Halo " + req.Name + ", Majama?",
	}

	response, err := proto.Marshal(resp)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	log.Printf("Sending response with content of length: %d", len(response))
	w.Write(response)
}

func grpcHelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("==========================================================")
	log.Printf("[Protobuf] Request received with content of length: %d", r.ContentLength)
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
	}

	req := &pb.HelloRequest{}
	proto.Unmarshal(data, req)
	log.Printf("Hello request form %s", req.Name)

	resp := &pb.HelloResponse{
		Message: "Hello " + req.Name + ", How's it going?",
	}
	response, err := proto.Marshal(resp)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	log.Printf("Sending response with content of length: %d", len(response))
	w.Write(response)
}

func jsonHelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("==========================================================")
	log.Printf("[Json] Request received with content of length: %d", r.ContentLength)
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
	}

	req := &pb.HelloRequest{}
	err = json.Unmarshal(data, req)
	if err != nil {
		log.Fatalf("Unable to unmarshal message from request : %v", err)
	}
	log.Printf("Hello request form %s", req.Name)

	resp := &pb.HelloResponse{
		Message: "Hello " + req.Name + ", How's it going?",
	}
	response, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	log.Printf("Sending response with content of length: %d", len(response))
	w.Write(response)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/grpc_hello", grpcHelloHandler).Methods("POST")
	router.HandleFunc("/json_hello", jsonHelloHandler).Methods("POST")

	haloHandler := &grpcHaloHandler{}
	router.Handle("/halo", haloHandler).Methods("POST")

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Server is running and listening on port 8080")
	log.Fatal(server.ListenAndServe())
}
