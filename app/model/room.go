package model

import (
	"errors"
	"fmt"
	"log"
	"mus_projekt/app/model/query"
	"mus_projekt/app/model/utils"
	"strconv"
	"time"
)

/*

	Manages the Room Info

	Datamanagement with the Database and Admin Access auth data

*/
//binary checks possible
const __ORIENATION_NORTH = int(1) //0b00001
const __ORIENATION_EAST = int(2)  //0b00010
const __ORIENATION_SOUTH = int(3) //0b00011
const __ORIENATION_WEST = int(4)  //0b00100

const __ORIENATION_NORTH_S = "N"
const __ORIENATION_EAST_S = "E"
const __ORIENATION_SOUTH_S = "S"
const __ORIENATION_WEST_S = "W"

var t0 = []int{8, 8, 9, 10, 11, 12, 13, 14, 15, 15, 16, 17, 18, 19}
var t1 = []int{0, 50, 45, 35, 30, 20, 15, 5, 0, 50, 45, 35, 30, 20}
var t2 = []int{8, 9, 10, 11, 12, 13, 14, 14, 15, 16, 17, 18, 19, 20}
var t3 = []int{45, 35, 30, 20, 15, 5, 0, 50, 45, 35, 30, 20, 15, 5}

type RoomLocation struct {
	room_floor_level int `json:"room_floor_level"` //Saves the floor level for the room
	room_corridor    int `json:"room_corridor"`    //The Direction for the facing of the Room Area
	room_number      int `json:"room_number"`      //The Room Number for the Room on the different floor

	room_name string //non database stored information about the room location, will be generated upon request
}

type Room struct {
	_id  string `json:"_id"`
	_rev string `json:"_rev`

	room_id string `json:"room_id"`

	room_name      string `json:"room_name"`
	room_name_misc string `json:"room_name_misc"`

	microcontroller_available bool `json:"arduino_attached"`

	room_bookable             bool `json:"room_bookable"`
	room_booking_max_duration int  `json:"room_booking_max_duration"`

	room_booking_ids []string `json:"room_booking_ids"` //All the bookings the Room is occupied with

	room_maximum_capacity int `json:"room_maximum_capacity"`

	room_location RoomLocation `json:"room_location"`
}

func GetRoomBy_RoomID(id string) (*Room, error) {
	query := `{"selector":{"room_id":"` + id + `"}}`
	m, err := room_DB.QueryJSON(query)
	if len(m) == 0 || err != nil {
		return nil, err
	}
	return m2r(m[0]), nil
}

func CreateNew(f int, c string) *Room {
	r := Room{
		_id:                       "-",
		_rev:                      "-",
		room_id:                   utils.GenerateID(),
		room_name:                 "-",
		room_name_misc:            "-",
		microcontroller_available: false,
		room_bookable:             false,
		room_booking_ids:          []string{},
		room_maximum_capacity:     0,
		room_booking_max_duration: 0,
		room_location: RoomLocation{
			room_floor_level: f,
			room_corridor:    D2i(c),
			room_number:      0,
			room_name:        "-",
		},
	}
	r.storeNew()
	return &r
}

func (r *Room) storeNew() (bool, error) {
	m := r2m(r)
	delete(m, query.ROOM_DB_ID)
	delete(m, query.ROOM_DB_REV)
	id, rev, err := room_DB.Save(m, nil)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	r._id = id
	r._rev = rev
	return true, nil
}

func (r *Room) Save() (bool, error) {
	m := r2m(r)
	id, rev, err := room_DB.Save(m, nil)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	r._id = id
	r._rev = rev
	return true, nil
}

func r2m(r *Room) map[string]interface{} {
	m := make(map[string]interface{})
	fmt.Println("\n\n\n")
	fmt.Println(r)
	m[query.ROOM_DB_ID] = r._id
	m[query.ROOM_ID] = r.room_id
	m[query.ROOM_DB_REV] = r._rev
	m[query.ROOM_NAME] = r.room_name
	m[query.ROOM_MAX_BOOKING_DURATION] = r.room_booking_max_duration
	m[query.ROOM_BOOKABLE] = r.room_bookable
	m[query.ROOM_ARDUINO_AVAILBLE] = r.microcontroller_available
	m[query.ROOM_NAME_MISC] = r.room_name_misc
	m[query.ROOM_BOOKING_IDS] = r.room_booking_ids
	m[query.ROOM_MAX_CAPACITY] = r.room_maximum_capacity
	l := make(map[string]interface{})
	l[query.LOCATION_FLOOR] = r.room_location.room_floor_level
	l[query.LOCATION_ROOM_CORRIDOR] = r.room_location.room_corridor
	l[query.LOCATION_ROOM_NUMBER] = r.room_location.room_number
	m[query.ROOM_LOCATION] = l
	return m
}

func m2r(m map[string]interface{}) *Room {
	r := Room{
		_id:                       m[query.ROOM_DB_ID].(string),
		_rev:                      m[query.ROOM_DB_REV].(string),
		room_id:                   m[query.ROOM_ID].(string),
		microcontroller_available: m[query.ROOM_ARDUINO_AVAILBLE].(bool),
		room_maximum_capacity:     int(m[query.ROOM_MAX_CAPACITY].(float64) + 0.5),
		room_booking_max_duration: int(m[query.ROOM_MAX_BOOKING_DURATION].(float64) + 0.5),
		room_name:                 m[query.ROOM_NAME].(string),
		room_bookable:             m[query.ROOM_BOOKABLE].(bool),
		room_name_misc:            m[query.ROOM_NAME_MISC].(string),
		room_booking_ids:          getBookingIDs(m[query.ROOM_BOOKING_IDS].([]interface{})),
		room_location: RoomLocation{
			room_floor_level: int(m[query.ROOM_LOCATION].(map[string]interface{})[query.LOCATION_FLOOR].(float64) + 0.5), //correct rounding to int
			room_corridor:    int(m[query.ROOM_LOCATION].(map[string]interface{})[query.LOCATION_ROOM_CORRIDOR].(float64) + 0.5),
			room_number:      int(m[query.ROOM_LOCATION].(map[string]interface{})[query.LOCATION_ROOM_NUMBER].(float64) + 0.5),
		},
	}
	r.room_location.room_name = generateLocationName(&r.room_location)
	return &r
}

func generateLocationName(l *RoomLocation) string {
	d := i2D(l.room_corridor)
	f := i2F(l.room_floor_level)
	n := i2Rn(l.room_number)
	return d + "-" + f + "-" + n
}

func getBookingIDs(m []interface{}) []string {
	s := []string{}
	for _, k := range m {
		s = append(s, k.(string))
	}
	return s
}

func (r *Room) Arduino_Available() bool {
	return r.microcontroller_available
}

func (r *Room) ID() string {
	return r.room_id
}

func (r *Room) Room_Name() string {
	return r.room_name
}

func (r *Room) Room_Name_Misc() string {
	return r.room_name_misc
}

func (r *Room) Location2M() map[string]interface{} {
	m := make(map[string]interface{})
	m[query.LOCATION_FLOOR] = r.room_location.room_floor_level
	m[query.LOCATION_ROOM_NUMBER] = r.room_location.room_number
	m[query.LOCATION_ROOM_CORRIDOR] = r.room_location.room_corridor
	m[query.LOCATION_ROOM_NAME] = generateLocationName(&r.room_location)
	return m
}

func (r *Room) Max_Capacity() int {
	return r.room_maximum_capacity
}

func (r *Room) Bookable() bool {
	return r.room_bookable
}

func (r *Room) Remove() error {
	err := room_DB.Delete(r._id)
	return err
}

func (r *Room) Max_Booking_Duration() int {
	return r.room_booking_max_duration
}

func GetAllRooms(f int, c string) *[]*Room {
	r := []*Room{}
	w := D2i(c)
	query := `{"selector":{"` + query.ROOM_LOCATION + `":{"` + query.LOCATION_FLOOR + `":` + strconv.Itoa(f) + `, "` + query.LOCATION_ROOM_CORRIDOR + `": ` + strconv.Itoa(w) + `}}}`
	m, _ := room_DB.QueryJSON(query)
	for _, room := range m {
		r = append(r, m2r(room))
	}
	return &r
}

func GetAllRoomsByInt(f int, c int) *[]*Room {
	r := []*Room{}
	query := `{"selector":{"` + query.ROOM_LOCATION + `":{"` + query.LOCATION_FLOOR + `":` + strconv.Itoa(f) + `, "` + query.LOCATION_ROOM_CORRIDOR + `": ` + strconv.Itoa(c) + `}}}`
	m, _ := room_DB.QueryJSON(query)
	for _, room := range m {
		r = append(r, m2r(room))
	}
	return &r
}

func (r *Room) Put(
	room_name string,
	room_name_msc string,
	room_bookable bool,
	room_duration int,
	room_location_name string,
	room_location_floor int,
	room_location_corridor int,
	room_location_number int,
	room_max_capacity int) {
	r.room_booking_max_duration = room_duration
	r.room_name = room_name
	r.room_name_misc = room_name_msc
	r.room_bookable = room_bookable
	r.room_maximum_capacity = room_max_capacity
	r.room_location.room_name = room_location_name
	r.room_location.room_number = room_location_number
	r.room_location.room_floor_level = room_location_floor
	r.room_location.room_corridor = room_location_corridor
	r.Save()
}

func (r *Room) RoomNumber() int {
	return r.room_location.room_number
}

func (r *Room) Location() *RoomLocation {
	return &r.room_location
}

func (l *RoomLocation) Name() string {
	return l.room_name
}

func (r *Room) GetBookingsByWeek(y int, w int) (*[]*[]*[]*Booking, error) {
	return GetBookingsByRoomID_Week(y, w, r.room_id)
}

func (r *Room) AttachNewBooking(b *Booking) {
	r.room_booking_ids = append(r.room_booking_ids, b.booking_id)
	r.Save()
}

func (r *Room) RemoveBookingByID(id string) error {
	for i, j := range r.room_booking_ids {
		if j == id {
			r.room_booking_ids[i] = r.room_booking_ids[len(r.room_booking_ids)-1]
			return nil
		}
	}
	r.Save()
	return errors.New("The id was not found")
}

func (r *Room) RemoveArduino() {
	if r == nil {
		return
	}
	r.microcontroller_available = false
}

func (r *Room) GetCurrentBookings() (*[]*Booking, error) {
	t := time.Now()
	y, w := t.ISOWeek()
	d := t.Day()
	ts := getTimeSlot(t.Hour(), t.Minute())
	b, err := GetBookingsBy_RoomID_YWDH(r.room_id, y, w, d, ts)
	if err != nil {
		return b, err
	}
	return b, nil
}

func getTimeSlot(h int, m int) int {
	for i, _ := range t0 {
		if (h == t0[i] && m >= t1[i]) || (h == t2[i] && m <= t3[i]) {
			fmt.Println("HAHAAAAAAA")
			return i
		}
	}
	return 0
}
