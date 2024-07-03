// Package main for starting a bouncer server instance
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/prettyboiiii/bouncer/internal/bouncer"
	"github.com/prettyboiiii/bouncer/internal/janitor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	s := grpc.NewServer()
	_ = bouncer.NewInstance(s)

	port := getServerPort()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer janitor.CheckClose(lis.Close, "listener")

	reflection.Register(s)

	log.Printf("serving on port: %d", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getServerPort() uint16 {
	var port uint16
	if parsed, err := strconv.ParseUint(os.Getenv("SERVER_PORT"), 10, 16); err == nil {
		port = uint16(parsed)
	} else {
		log.Fatalf("failed to parse port: %v", err)
	}
	return port
}
