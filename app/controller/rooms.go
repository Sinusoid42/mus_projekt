package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"mus_projekt/app/auth"
	"mus_projekt/app/controller/bookings"
	"mus_projekt/app/controller/protocol"
	"mus_projekt/app/controller/rooms"
	"mus_projekt/app/controller/users"
	"mus_projekt/app/model"
	"mus_projekt/app/model/query"
	"net/http"
	"strconv"
	"time"
)

/**
	For accessing the rooms page or the rooms api page, this file contains all functionalities abstracting the http communication with a webclient

	The method RoomsAPI abstracts the functionalities for the api rooms webpage, where rooms can be removed, altered and added from the webinterface

	The RoomsHandler function is handling all communication when connected to a single calender app from within a webbrowser

	@author ben
 */

func RoomsApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("http:Rooms_Api")

	if r.Method == http.MethodGet {
		ExecuteTemplate("room_template.tmpl", w, r, nil)
		return
	}

	/**
	Method Put, when saving the settings of a room
	*/
	if r.Method == http.MethodPut {
		id := r.FormValue(query.ROOM_ID)
		a, err, room := rooms.CheckByID(id)

		if a && err == nil {

			//room service_id is there, the room can be found and saved
			fmt.Println(room)

			m, err := rooms.Put(room, r)
			if err != nil {
				m := make(map[string]interface{})
				m[query.SUCCESS] = false
				b, _ := json.MarshalIndent(m, "", "  ")
				w.Header().Set("Content-Type", "application/json")
				w.Write(b)
				return
			}
			fmt.Println(r)
			m[query.SUCCESS] = true
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}
		m := make(map[string]interface{})
		m[query.SUCCESS] = false
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}

	/**
	Method Post, when storing a new room
	*/
	if r.Method == http.MethodPost {
		id := r.FormValue(query.ROOM_ID)
		a, err, _ := rooms.CheckByID(id)
		fmt.Println(a, err)
		if err != nil || !a {
			//Create new Room
			f := r.FormValue(query.LOCATION_FLOOR)
			c := r.FormValue(query.LOCATION_ROOM_CORRIDOR)
			i, _ := strconv.Atoi(f)
			r := rooms.CreateNew(i, c)
			m := rooms.CreateAPIData(r)
			m[query.SUCCESS] = true
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}

	}

	/**
	Method Delete, Room deletion
	*/
	if r.Method == http.MethodDelete {
		fmt.Println("http:DELETE:RoomAPI")
		id := r.FormValue(query.ROOM_ID)
		a, err, room := rooms.CheckByID(id)
		if err == nil && a {
			room.Remove()
			m := make(map[string]interface{})
			m[query.SUCCESS] = true
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		} else {
			m := make(map[string]interface{})
			m[query.SUCCESS] = false
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}

	}

	/**
	Method Fetch, selection of a room query by corridor and floor for the api frontpage
	*/
	if r.Method == protocol.MethodFetch {
		fmt.Println("Testing")
		id := r.FormValue(query.ROOM_ID)
		a, err, room := rooms.CheckByID(id)
		if a && err == nil {
			fmt.Println("Hello world")
			m := rooms.CreateAPIData(room)
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}
		f := r.FormValue(query.LOCATION_FLOOR)
		c := r.FormValue(query.LOCATION_ROOM_CORRIDOR)
		i, _ := strconv.Atoi(f)
		m := rooms.ServeAllRooms(i, c)
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
}

func Rooms(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println("http:ROOMS:GET")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func RoomsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		fmt.Println("http:ROOMS:GET")
		var m map[string]interface{}
		m = make(map[string]interface{})

		session, _ := auth.GetCookieStore().Get(r, auth.GetSessionCookie())
		a := session.Values[auth.USER_COOKIE_AUTH]
		l := session.Values[auth.USER_ACCESS_LEVEL]
		id := session.Values[auth.USER_ID]

		v := mux.Vars(r)
		rid := v[query.ROOM_ID]

		j, er := model.Check_User_By_ID(id)

		if !j || er != nil {
			m = users.CreateApiData(nil)
			m[query.SUCCESS] = true
			m[query.AUTH] = false
			m[query.ADMIN_LEVEL] = 0

			yr, wk := time.Now().ISOWeek()
			m[query.BOOKING_ROOM_ID] = rid
			m[query.CALENDER_YEAR] = yr
			m[query.CALENDER_WEEK] = strconv.Itoa(yr) + "-W" + strconv.Itoa(wk)
			html_templates.ExecuteTemplate(w, "rooms.tmpl", m)
			return
		}
		u, _ := model.GetUserByID(id.(string))
		m = users.CreateApiData(u)
		m[query.SUCCESS] = true
		m[query.AUTH] = a.(bool)
		m[query.ADMIN_LEVEL] = l.(int)
		m[query.BOOKING_ROOM_ID] = rid
		yr, wk := time.Now().ISOWeek()
		m[query.CALENDER_YEAR] = yr
		m[query.CALENDER_WEEK] = strconv.Itoa(yr) + "-W" + strconv.Itoa(wk)
		m[query.CALENDER_YEAR] = time.Now().Year()

		api_rooms, api_arduinos, api_templates := users.CheckAccess(u)
		m[query.HTML_API_ROOMS_ACCESS] = api_rooms
		m[query.HTML_API_ARDUINOS_ACCESS] = api_arduinos
		m[query.HTML_API_TEMPLATES_ACCESS] = api_templates

		ExecuteTemplate("rooms.tmpl", w, r, m)
	}

	if r.Method == protocol.MethodFetch {
		session, _ := auth.GetCookieStore().Get(r, auth.GetSessionCookie())
		a := session.Values[auth.USER_COOKIE_AUTH]
		l := session.Values[auth.USER_ACCESS_LEVEL]
		id := session.Values[auth.USER_ID]

		v := mux.Vars(r)
		rid := v[query.ROOM_ID]

		yr := r.FormValue(query.CALENDER_YEAR)
		mt := r.FormValue(query.CALENDER_WEEK)
		rm, err := model.GetRoomBy_RoomID(rid)

		if err != nil {
			m := make(map[string]interface{})
			m[query.SUCCESS] = false
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}
		yi, wi, err := rooms.GetISOWeekFromString(yr, mt)
		if err != nil {
			m := make(map[string]interface{})
			m[query.SUCCESS] = false
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}
		m := rooms.CreateCalenderData(rm, yi, wi)
		if a == nil || l == nil || id == nil || len(id.(string)) == 0 {
			m[query.AUTH] = true
			m[query.ADMIN_LEVEL] = 0
			m[query.SUCCESS] = true
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}
		u, e := model.GetUserByID(id.(string))
		m[query.AUTH] = false
		m[query.ADMIN_LEVEL] = 0
		if a != nil && a.(bool) && e == nil {
			m[query.AUTH] = true
			m[query.ADMIN_LEVEL] = u.GetAccessLevel()
			session.Values[auth.USER_ACCESS_LEVEL] = u.GetAccessLevel()
		}
		m[query.SUCCESS] = true
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}

	if r.Method == http.MethodDelete {
		v := mux.Vars(r)
		rid := v[query.ROOM_ID]
		rm, err := model.GetRoomBy_RoomID(rid)
		session, _ := auth.GetCookieStore().Get(r, auth.GetSessionCookie())
		fmt.Println(session)

		a := session.Values[auth.USER_COOKIE_AUTH]
		//l := session.Values[auth.USER_ACCESS_LEVEL]
		id := session.Values[auth.USER_ID]

		bid := r.FormValue(query.BOOKING_ID)

		if err != nil || len(bid) != 16 || a == nil || (a != nil && !a.(bool)) {
			m := make(map[string]interface{})
			m[query.SUCCESS] = false
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}
		u, err := model.GetUserByID(id.(string))
		booking, err := model.GetBookingByBookingID(bid)
		if err != nil {
			m := make(map[string]interface{})
			m[query.SUCCESS] = false
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}

		fmt.Println("\n\n", booking, "\n\n", u, "\n\n", rm)

		if bookings.MatchDeletion(booking, u, rm) {
			e := booking.Remove()
			if e == nil {
				m := make(map[string]interface{})
				m[query.SUCCESS] = true
				b, _ := json.MarshalIndent(m, "", "  ")
				w.Header().Set("Content-Type", "application/json")
				w.Write(b)
				return
			}
		}
		m := make(map[string]interface{})
		m[query.SUCCESS] = false
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}

	if r.Method == http.MethodPost { //create a new booking for the client to the corresponding checks

		//TODO Auth Checks
		fmt.Println("\n\nhttp:POST:rooms")

		v := mux.Vars(r)
		rid := v[query.ROOM_ID]
		rm, err := model.GetRoomBy_RoomID(rid)
		session, _ := auth.GetCookieStore().Get(r, auth.GetSessionCookie())
		fmt.Println(session)

		a := session.Values[auth.USER_COOKIE_AUTH]
		//l := session.Values[auth.USER_ACCESS_LEVEL]
		id := session.Values[auth.USER_ID]
		fmt.Println(id)
		if a == nil || !a.(bool) || err != nil || rm == nil || !rm.Bookable() || id == nil {
			fmt.Println("Escaping0")
			m := make(map[string]interface{})
			m[query.SUCCESS] = false
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}
		us, e := model.GetUserByID(id.(string))
		if e != nil {
			fmt.Println("Escaping1")
			m := make(map[string]interface{})
			m[query.SUCCESS] = false
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}
		/*
			//fetch data from query
			bd := r.FormValue(query.BOOKING_TIME_WEEK)
			bw := r.FormValue(query.BOOKING_TIME_WEEK)
			by := r.FormValue(query.BOOKING_TIME_YEAR)
			bm := r.FormValue(query.BOOKING_TIME_MONTH)

			bs := r.FormValue(query.BOOKING_TIME_START)
			be := r.FormValue(query.BOOKING_TIME_END)

			br := r.FormValue(query.BOOKING_TIME_REOCCURING)

			btp := r.FormValue(query.BOOKING_TOPIC)
			bds := r.FormValue(query.BOOKING_DESCRIPTION)

			bed := ""
			bew := ""
			bey := ""
			bem := ""

			if b, _ := strconv.ParseBool(br); b {
				bed = r.FormValue(query.BOOKING_TIME_END_DAY)
				bew = r.FormValue(query.BOOKING_TIME_END_WEEK)
				bey = r.FormValue(query.BOOKING_TIME_END_YEAR)
				bem = r.FormValue(query.BOOKING_TIME_END_MONTH)
			}

			fmt.Println("UID: ", id)
			fmt.Println("ALevel: ", l)
			fmt.Println("RID: ", rid)
			fmt.Println("DATE", bd, bw, by, bm, "START/END:", bs, be, br, bed, "ENDDATE: ", bew, bey, bem, "TOPIC: ", btp, "Description: ", bds)
			fmt.Println(a, l, id)*/

		m, _ := bookings.CreateNewBooking(a.(bool), rm, us, r)
		fmt.Println(m)
		m[query.SUCCESS] = true
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return

	}

	if r.Method == protocol.MethodSelect {
		fmt.Println("\n\nhttp:SELECT:bookings_selection")

		session, _ := auth.GetCookieStore().Get(r, auth.GetSessionCookie())
		fmt.Println(session)
		//l := session.Values[auth.USER_ACCESS_LEVEL]
		id := session.Values[auth.USER_ID]

		v := mux.Vars(r)
		rid := v[query.ROOM_ID]
		rm, err := model.GetRoomBy_RoomID(rid)
		if err != nil || rm == nil {
			fmt.Println("Escaping")
			m := make(map[string]interface{})
			m[query.SUCCESS] = false
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}

		cy := r.FormValue(query.CALENDER_YEAR)
		cw := r.FormValue(query.CALENDER_WEEK)
		cd := r.FormValue(query.CALENDER_DAY)
		ch := r.FormValue(query.CALENDER_HOUR)

		//TODO get bookings by week year hour and day
		yi, _ := strconv.Atoi(cy)
		wi, _ := strconv.Atoi(cw)
		di, _ := strconv.Atoi(cd)
		hi, _ := strconv.Atoi(ch)

		B, err := model.GetBookingsBy_RoomID_YWDH(rid, yi, wi, di, hi)

		if err != nil || B == nil || (B != nil && len(*B) == 0) {
			fmt.Println("Escaping")
			m := make(map[string]interface{})
			m[query.SUCCESS] = false
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}

		m, _ := bookings.CreateSelectionData(id, rm, B)
		m[query.SUCCESS] = true
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
}
