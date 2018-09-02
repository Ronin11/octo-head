package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"net/http"
	
	"github.com/julienschmidt/httprouter"

	"github.com/Ronin11/octo-tentacle/pkg/octo"
	"github.com/Ronin11/octo-tentacle/services/sprinkler"
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
	action := sprinklerService.SprinklerAction{}
	action.Channel = "sprinkler.backyard.0.input"
	action.Name = "Sprinkler Start"
	action.State = sprinklerService.SprinklerData{
		SprinklerIsOn: true,
	}
	octo.SendAction(action, network)
	fmt.Fprintf(w, "OK")
}
func sprinkler0off(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	action := sprinklerService.SprinklerAction{}
	action.Channel = "sprinkler.backyard.0.input"
	action.Name = "Sprinkler End"
	action.State = sprinklerService.SprinklerData{
		SprinklerIsOn: false,
	}
	octo.SendAction(action, network)
	fmt.Fprintf(w, "OK")
}

func sprinkler1on(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	action := sprinklerService.SprinklerAction{}
	action.Channel = "sprinkler.backyard.1.input"
	action.Name = "Sprinkler Start"
	action.State = sprinklerService.SprinklerData{
		SprinklerIsOn: true,
	}
	octo.SendAction(action, network)
	fmt.Fprintf(w, "OK")
}

func sprinkler1off(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	action := sprinklerService.SprinklerAction{}
	action.Channel = "sprinkler.backyard.1.input"
	action.Name = "Sprinkler End"
	action.State = sprinklerService.SprinklerData{
		SprinklerIsOn: false,
	}
	octo.SendAction(action, network)
	fmt.Fprintf(w, "OK")
}

func startSprinklers(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	finalAction := sprinklerService.SprinklerAction{}
	finalAction.Channel = "sprinkler.backyard.1.input"
	finalAction.Name = "Sprinkler End"
	finalAction.State = sprinklerService.SprinklerData{
		SprinklerIsOn: true,
		Duration: 10,
	}
	

	intialAction := sprinklerService.SprinklerAction{}
	intialAction.Channel = "sprinkler.backyard.0.input"
	intialAction.Name = "Sprinkler End"
	intialAction.State = sprinklerService.SprinklerData{
		SprinklerIsOn: true,
		Duration: 5,
	}
	intialAction.OnDone = finalAction
	octo.SendAction(intialAction, network)
	fmt.Fprintf(w, "OK")
}

func main() {
	fmt.Println("\n~~~~~~~ Starting Mantle ~~~~~~~")

	network = octo.JoinNetwork(os.Getenv("SERVER"), octo.NATSNetwork)

	err := octo.CreateListener(network, func(message string, subject string) {
		fmt.Printf("Subject: %s \tMessage: %s\n", subject, message)
	})
	if err != nil{
		panic(err)
	}

	router := httprouter.New()
	router.GET("/discovery", discovery)
	router.GET("/startSprinklers", startSprinklers)
	router.GET("/sprinkler0on", sprinkler0on)
	router.GET("/sprinkler0off", sprinkler0off)
	router.GET("/sprinkler1on", sprinkler1on)
	router.GET("/sprinkler1off", sprinkler1off)

	log.Fatal(http.ListenAndServe(":8080", router))
}