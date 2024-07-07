// Package test TODO: docs
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/prettyboiiii/bouncer/internal/janitor"
	"github.com/prettyboiiii/bouncer/pkg/bouncer"
	"github.com/prettyboiiii/bouncer/pkg/client"
	"google.golang.org/grpc"
)

func main() {
	c, err := client.Dial("localhost:8080", client.GRPCWrapper(grpc.WithInsecure()))
	if err != nil {
		log.Fatalln(err)
	}
	defer janitor.CheckClose(c, "client")

	b := bouncer.NewBouncer(context.Background(), c)

	fmt.Println(b)
}
