package main

// curl --request GET 127.0.0.1:9000
import (
	//"fmt"
	"rest-api/server"
	//"net/http"
	"os"
	"os/signal"
	"syscall"
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	serverDoneChan := make(chan os.Signal, 1)
	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)
	srv := server.New(":8080")

	go srv.ListenAndServe()
	log.Printf("Server started")
	<- serverDoneChan 
	srv.Shutdown(ctx)
	log.Printf("Server stopped")
}
