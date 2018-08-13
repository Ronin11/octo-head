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

func sprinkler0on(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	messenger := octo.CreateMessenger("sprinkler.backyard.0.input", network)
	duration := time.Second
	time.Sleep(duration)
	messenger.Write(`{"Name":"Action Description","State":{"sprinklerIsOn": true,"Duration":"SomeDuration"},"onDone":{"name": "ON DONE"}}`)
	messenger.Subscribe(func(message string){
		fmt.Println("RESPONSE: ", message)
	})
	fmt.Fprintf(w, "OK")
}
func sprinkler0off(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	messenger := octo.CreateMessenger("sprinkler.backyard.0.input", network)
	duration := time.Second
	time.Sleep(duration)
	messenger.Write(`{"Name":"Action Description","State":{"sprinklerIsOn": false,"Duration":"SomeDuration"},"onDone":{"name": "ON DONE"}}`)
	messenger.Subscribe(func(message string){
		fmt.Println("RESPONSE: ", message)
	})
	fmt.Fprintf(w, "OK")
}

func sprinkler1on(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	messenger := octo.CreateMessenger("sprinkler.backyard.1.input", network)
	duration := time.Second
	time.Sleep(duration)
	messenger.Write(`{"Name":"Action Description","State":{"sprinklerIsOn": true,"Duration":"SomeDuration"},"onDone":{"name": "ON DONE"}}`)
	messenger.Subscribe(func(message string){
		fmt.Println("RESPONSE: ", message)
	})
	fmt.Fprintf(w, "OK")
}
func sprinkler1off(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	messenger := octo.CreateMessenger("sprinkler.backyard.1.input", network)
	duration := time.Second
	time.Sleep(duration)
	messenger.Write(`{"Name":"Action Description","State":{"sprinklerIsOn": false,"Duration":"SomeDuration"},"onDone":{"name": "ON DONE"}}`)
	messenger.Subscribe(func(message string){
		fmt.Println("RESPONSE: ", message)
	})
	fmt.Fprintf(w, "OK")
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
	router.GET("/sprinkler0on", sprinkler0on)
	router.GET("/sprinkler0off", sprinkler0off)
	router.GET("/sprinkler1on", sprinkler1on)
	router.GET("/sprinkler1off", sprinkler1off)

	log.Fatal(http.ListenAndServe(":8080", router))
}