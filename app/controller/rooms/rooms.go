package rooms

import (
	"errors"
	"fmt"
	"mus_projekt/app/model"
	"mus_projekt/app/model/query"
	"net/http"
	"strconv"
)

const BUILDING_FLOORS int = 13

func CheckByID(id string) (bool, error, *model.Room) {
	if len(id) != 16 {
		return false, errors.New("no id"), nil
	}
	room, err := model.GetRoomBy_RoomID(id)
	fmt.Println(room)
	if err != nil {
		return false, err, nil
	}
	return true, nil, room
}

func CreateNew(f int, c string) *model.Room {
	return model.CreateNew(f, c)
}

func CreateAPIData(r *model.Room) map[string]interface{} {
	m := make(map[string]interface{})
	fmt.Println("HAHAHAHAHAHA")
	m[query.ROOM_ARDUINO_AVAILBLE] = r.Arduino_Available()
	m[query.ROOM_ID] = r.ID()
	m[query.ROOM_NAME] = r.Room_Name()
	m[query.ROOM_NAME_MISC] = r.Room_Name_Misc()
	m[query.ROOM_LOCATION] = r.Location2M()
	m[query.ROOM_MAX_CAPACITY] = r.Max_Capacity()
	m[query.ROOM_BOOKABLE] = r.Bookable()
	m[query.ROOM_MAX_BOOKING_DURATION] = r.Max_Booking_Duration()
	m[query.SUCCESS] = true
	return m
}

func ServeAllRooms(f int, c string) map[string]interface{} {
	r := model.GetAllRooms(f, c)
	m := make(map[string]interface{})
	m["rooms"] = []map[string]interface{}{}
	fmt.Println(r)
	for _, k := range *r {
		m["rooms"] = append(m["rooms"].([]map[string]interface{}), CreateAPIData(k))
	}
	return m
}

func Put(r *model.Room, rq *http.Request) (map[string]interface{}, error) {

	room_name := rq.FormValue(query.ROOM_NAME)
	room_name_msc := rq.FormValue(query.ROOM_NAME_MISC)
	room_bookable := rq.FormValue(query.ROOM_BOOKABLE)
	room_location_name := rq.FormValue(query.LOCATION_ROOM_NAME)
	room_location_floor := rq.FormValue(query.LOCATION_FLOOR)
	room_location_corridor := rq.FormValue(query.LOCATION_ROOM_CORRIDOR)
	room_location_number := rq.FormValue(query.LOCATION_ROOM_NUMBER)
	room_max_capacity := rq.FormValue(query.ROOM_MAX_CAPACITY)
	room_max_duration := rq.FormValue(query.ROOM_MAX_BOOKING_DURATION)
	book, err0 := strconv.ParseBool(room_bookable)
	rlf, err1 := strconv.Atoi(room_location_floor)
	rlc, err2 := strconv.Atoi(room_location_corridor)
	rln, err3 := strconv.Atoi(room_location_number)
	rmx, err4 := strconv.Atoi(room_max_capacity)
	rmd, err5 := strconv.Atoi(room_max_duration)
	if err0 != nil || err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		fmt.Println(err0, err1, err2, err3, err4)
		return nil, errors.New("Number could not be parsed")
	}
	fmt.Println(rmx)
	r.Put(room_name,
		room_name_msc,
		book,
		rmd,
		room_location_name,
		rlf,
		rlc,
		rln,
		rmx)
	return CreateAPIData(r), nil
}

func CreateCalenderData(r *model.Room, yr int, week int) map[string]interface{} {
	m := make(map[string]interface{})
	b, _ := r.GetBookingsByWeek(yr, week)
	i := 0
	j := 0
	for i = 0; i < 5; i++ {
		var m_ = make(map[string]interface{})
		m["m"+strconv.Itoa(i)] = m_
		for j = 0; j < 14; j++ {
			p := make(map[string]interface{})
			m_["m"+strconv.Itoa(i)+"_"+strconv.Itoa(j)] = p
			p["bid"] = []string{}
			p["o"] = false
			p["b"] = true
			p["oc"] = 0
		}
	}

	for d, j := range *b {
		for h, s := range *j {
			for _, booking := range *s {
				for p := 0; p < booking.Duration()+1; p++ {
					i := m["m"+strconv.Itoa(d)].(map[string]interface{})["m"+strconv.Itoa(d)+"_"+strconv.Itoa(h+p)].(map[string]interface{})["oc"].(int)
					i = i + booking.GetCount()
					m["m"+strconv.Itoa(d)].(map[string]interface{})["m"+strconv.Itoa(d)+"_"+strconv.Itoa(h+p)].(map[string]interface{})["oc"] = i
					m["m"+strconv.Itoa(d)].(map[string]interface{})["m"+strconv.Itoa(d)+"_"+strconv.Itoa(h+p)].(map[string]interface{})["o"] = true
					q := m["m"+strconv.Itoa(d)].(map[string]interface{})["m"+strconv.Itoa(d)+"_"+strconv.Itoa(h+p)].(map[string]interface{})["bid"].([]string)
					q = append(q, booking.ID())
					m["m"+strconv.Itoa(d)].(map[string]interface{})["m"+strconv.Itoa(d)+"_"+strconv.Itoa(h+p)].(map[string]interface{})["bid"] = q
				}
			}
		}
	}
	m[query.ROOM_MAX_CAPACITY] = r.Max_Capacity()
	m[query.ROOM_BOOKABLE] = r.Bookable()
	m[query.BOOKING_ROOM_ID] = r.ID()
	m[query.ROOM_MAX_BOOKING_DURATION] = r.Max_Booking_Duration()

	return m
}

func GetISOWeekFromString(y string, w string) (int, int, error) {
	if len(y) == 0 || len(w) == 0 {
		return -1, -1, errors.New("No week or year provided")
	}
	yi, _ := strconv.Atoi(y)
	wi, _ := strconv.Atoi(w)
	return yi, wi, nil

}
