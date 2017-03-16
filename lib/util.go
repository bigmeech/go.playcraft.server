package go_playcraft_server

import "io/ioutil"
import "github.com/op/go-logging";

var log = logging.MustGetLogger("playcraft");

func checkError(e error){
	if e != nil {
		panic(e);
	}
}

func loadFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	checkError(err)

	log.Debug("Reading Configuration file from " + path);
	return data
}


func GetConfig(){
	loadFile("/config/local.yaml")
}
