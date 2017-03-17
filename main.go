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

	configRoutes(mux)

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


func userHandler(rw http.ResponseWriter, req *http.Request){
	defaultHandler(rw , req)
}

func teamHandler(rw http.ResponseWriter, req *http.Request){
	defaultHandler(rw , req)
}

func stratgegyHandler(rw http.ResponseWriter, req *http.Request){
	defaultHandler(rw , req)
}

func catchAll(rw http.ResponseWriter, req *http.Request){
	defaultHandler(rw , req);
}


//Configure routes
func configRoutes(mux bone.Mux){

	//INDEX
	mux.Get("/", http.HandlerFunc(defaultHandler))
	mux.Get("/api", http.HandlerFunc(defaultHandler))

	//USER API
	mux.Get("/api/user/", http.HandlerFunc(userHandler))
	mux.Get("/api/user/:id", http.HandlerFunc(userHandler))

	mux.Put("/api/user/", http.HandlerFunc(userHandler))
	mux.Put("/api/user/:id", http.HandlerFunc(userHandler))

	mux.Post("/api/user", http.HandlerFunc(userHandler))
	mux.Post("/api/user/:id", http.HandlerFunc(userHandler))

	mux.Delete("/api/user", http.HandlerFunc(userHandler))
	mux.Delete("/api/user/:id", http.HandlerFunc(userHandler))

	//TEAM API
	mux.Get("/api/team/", http.HandlerFunc(teamHandler))
	mux.Get("/api/team/:id", http.HandlerFunc(teamHandler))

	mux.Put("/api/team/", http.HandlerFunc(teamHandler))
	mux.Put("/api/team/:id", http.HandlerFunc(teamHandler))

	mux.Post("/api/team", http.HandlerFunc(teamHandler))
	mux.Post("/api/team/:id", http.HandlerFunc(teamHandler))

	mux.Delete("/api/team", http.HandlerFunc(teamHandler))
	mux.Delete("/api/team/:id", http.HandlerFunc(teamHandler))

	//STRATEGY API
	mux.Get("/api/strategy/", http.HandlerFunc(stratgegyHandler))
	mux.Get("/api/strategy/:id", http.HandlerFunc(stratgegyHandler))

	mux.Put("/api/strategy/", http.HandlerFunc(stratgegyHandler))
	mux.Put("/api/strategy/:id", http.HandlerFunc(stratgegyHandler))

	mux.Post("/api/strategy", http.HandlerFunc(stratgegyHandler))
	mux.Post("/api/strategy/:id", http.HandlerFunc(stratgegyHandler))

	mux.Delete("/api/strategy", http.HandlerFunc(stratgegyHandler))
	mux.Delete("/api/strategy/:id", http.HandlerFunc(stratgegyHandler))
}
