package main

import (
	"github.com/op/go-logging"
	"github.com/go-zoo/bone"
	"net/http"
)

var log = logging.MustGetLogger("playcraft");

func TeamHandler(rw http.ResponseWriter, req *http.Request){
	teamId := bone.GetValue(req, "id");
	rw.Write([]byte("Returning result for team id: " + teamId))
}

func defaultHandler(rw http.ResponseWriter, req *http.Request){
	rw.Write([]byte("Nothing to see in " + req.RequestURI))
}

func main(){
	mux := bone.New()
	mux.Get("/", http.HandlerFunc(defaultHandler))
	mux.Get("/api", http.HandlerFunc(defaultHandler))
	mux.Get("/api/team/:id", http.HandlerFunc(TeamHandler))

	//loading config
	util.LoadConfig()

	// starting api server
	log.Debug("Starting go.playcraft.server at 8080");
	http.ListenAndServe(":8080", mux)
}
