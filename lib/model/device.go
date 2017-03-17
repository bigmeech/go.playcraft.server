package go_playcraft_server

import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"

type Device struct {
	Id string
	Name string
	User
}

//adds a device
func (d *Device) Add(u User){
	session, err := mgo.Dial("localhost")
	reportError(err)
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("playcraft_db").C("device");
	err = c.Insert(u)
}

//reports error
func reportError(err error){
	if err != nil {
		panic(err)
	}
}

//deletes a device
func (d *Device) Delete(u User){}
