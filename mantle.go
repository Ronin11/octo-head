package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	
	"github.com/julienschmidt/httprouter"
	"github.com/Ronin11/octo-tentacle/pkg/octo"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	fmt.Println("\n~~~~~~~ Starting Mantle ~~~~~~~")

	network := octo.JoinNetwork(os.Getenv("SERVER"), octo.NATSNetwork)

	err := octo.CreateNatsListener(network.GetServerAddress(), func(message string, subject string) {
		fmt.Printf("Subject: %s \tMessage: %s\n", subject, message)
	})
	if err != nil{
		panic(err)
	}
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}