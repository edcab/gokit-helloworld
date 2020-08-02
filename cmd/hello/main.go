package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/edcab/gokit-helloworld/pkg/helloworld"
	"github.com/edcab/gokit-helloworld/pkg/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()

	errChan := make(chan error)

	service := helloworld.NewService()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := server.Endpoints{
		Hello: server.MakeGetEndpoint(service),
	}

	// HTTP transport
	go func() {
		log.Println("helloWorld is listening on port:", *httpAddr)
		handler := server.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
