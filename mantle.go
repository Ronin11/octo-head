package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"net/http"
	
	"github.com/julienschmidt/httprouter"
	"github.com/Ronin11/octo-tentacle/pkg/octo"
)

// func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//     fmt.Fprint(w, "Welcome!\n")
// }

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//     fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

var network *octo.Network

func discovery(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var services []string
	messenger := octo.CreateMessenger("discovery", network)
	messenger.Write("?")
	messenger.Subscribe(func(message string){
		services = append(services, message)
	})
	time.Sleep(time.Second)

	fmt.Fprintf(w, "DISCOVERY:, %s\n", services)
}

func main() {
	fmt.Println("\n~~~~~~~ Starting Mantle ~~~~~~~")

	network = octo.JoinNetwork(os.Getenv("SERVER"), octo.NATSNetwork)

	// err := octo.CreateListener(network, func(message string, subject string) {
	// 	fmt.Printf("Subject: %s \tMessage: %s\n", subject, message)
	// })
	// if err != nil{
	// 	panic(err)
	// }

	router := httprouter.New()
	router.GET("/discovery", discovery)
	// router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}