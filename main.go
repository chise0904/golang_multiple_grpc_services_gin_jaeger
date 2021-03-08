package main

import (
	"golang_multiple_grpc_services_gin_jaeger/grpcServer"
	"golang_multiple_grpc_services_gin_jaeger/httpServer"
)

func main() {

	ch := make(chan struct{})

	go grpcServer.Run()
	go httpServer.Run()

	<-ch
}
