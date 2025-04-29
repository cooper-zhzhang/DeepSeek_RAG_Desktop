package main

import (
	"dp_client/controller"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	obj := controller.NewTestGUI()
	obj.Run()

}
