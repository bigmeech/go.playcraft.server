package main

import (
	logging "github.com/op/go-logging"
	bone "github.com/go-zoo/bone"
	http "net/http"
	flag "flag"
	playcraft "github.com/bigmeech/go.playcraft.server/lib"
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

	DB_USERNAME := flag.String("db_username", "default_username", "Username to use for database login")
	DB_PASSWORD := flag.String("db_password", "default_password", "Password to use for database login")
	flag.Parse();

	log.Info(*DB_USERNAME, *DB_PASSWORD)

	//loading config
	playcraft.GetConfig()

	// starting api server
	log.Debug("Starting go.playcraft.server at 8080");
	http.ListenAndServe(":8080", mux)
}
