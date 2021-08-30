package model

import (
	"github.com/leesper/couchdb-golang"
)

var user_DB *couchdb.Database
var booking_DB *couchdb.Database
var room_DB *couchdb.Database
var arduino_DB *couchdb.Database

var template_DB *couchdb.Database

func init() {
	var err error
	user_DB, err = couchdb.NewDatabase("http://localhost:5984/ros_rocks_users")
	booking_DB, err = couchdb.NewDatabase("http://localhost:5984/ros_rocks_bookings")
	room_DB, err = couchdb.NewDatabase("http://localhost:5984/ros_rocks_rooms")
	arduino_DB, err = couchdb.NewDatabase("http://localhost:5984/ros_rocks_arduinos")
	template_DB, err = couchdb.NewDatabase("http://localhost:5984/ros_rocks_templates")
	if err != nil {
		print("Error has occured:")
		print(err)
	} else {
		print("The Databases have been initialized properly")
	}
}
