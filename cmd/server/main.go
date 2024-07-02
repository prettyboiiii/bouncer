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
)

func main() {
	s := grpc.NewServer()
	_ = bouncer.NewInstance(s)

	var port uint16
	if parsed, err := strconv.ParseUint(os.Getenv("SERVER_PORT"), 10, 16); err == nil {
		port = uint16(parsed)
	} else {
		log.Fatalf("failed to parse port: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer janitor.CheckClose(lis.Close, "listener")

	log.Printf("serving on port: %d", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
