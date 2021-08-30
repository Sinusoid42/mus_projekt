package bookings

import (
	"fmt"
	"mus_projekt/app/auth"
	"mus_projekt/app/model"
	"mus_projekt/app/model/query"
	"net/http"
	"strconv"
)

func CreateSelectionData(s interface{}, r *model.Room, b *[]*model.Booking) (map[string]interface{}, error) {

	j := false

	if s != nil && len(s.(string)) != 0 {
		j = true
	}

	var us *model.User
	if j {
		us, _ = model.GetUserByID(s.(string))

	}

	m := make(map[string]interface{})
	m["bookings"] = []map[string]interface{}{}
	c := 0
	for _, i := range *b {

		k := make(map[string]interface{})
		k[query.BOOKING_ID] = i.ID()
		k[query.BOOKING_TOPIC] = i.Topic()
		c = c + i.GetCount()
		u, _ := model.GetUserByID(i.UserID())

		k[query.BOOKING_USER_ID] = u.GetID()
		k[query.USER_NM] = u.GetUserName()
		if j {
			k["o"] = MatchDeletion(i, us, r)
		}

		k[query.BOOKING_TIME_START] = i.StartTime()
		k[query.BOOKING_TIME_END] = i.EndTime()
		m["bookings"] = append(m["bookings"].([]map[string]interface{}), k)

	}
	m[query.BOOKING_OCCUPANCY] = c
	return m, nil
}

func CreateNewBooking(auth bool, room *model.Room, user *model.User, r *http.Request) (map[string]interface{}, error) {

	fmt.Println("55 bookings.go", r.Form)
	//fetch data from query
	bd := r.FormValue(query.BOOKING_TIME_DAY)
	bw := r.FormValue(query.BOOKING_TIME_WEEK)
	by := r.FormValue(query.BOOKING_TIME_YEAR)
	bm := r.FormValue(query.BOOKING_TIME_MONTH)

	bs := r.FormValue(query.BOOKING_TIME_START)
	be := r.FormValue(query.BOOKING_TIME_END)

	br := r.FormValue(query.BOOKING_TIME_REOCCURING)
	fmt.Println("ADKIöjoegnpoasingv")
	fmt.Println(br)

	btp := r.FormValue(query.BOOKING_TOPIC)
	bds := r.FormValue(query.BOOKING_DESCRIPTION)

	bo := r.FormValue(query.BOOKING_OCCUPANCY)

	bed := ""
	bew := ""
	bey := ""
	bem := ""

	booking_reoccuring, _ := strconv.ParseBool(br)

	if booking_reoccuring {
		bed = r.FormValue(query.BOOKING_TIME_END_DAY)
		bew = r.FormValue(query.BOOKING_TIME_END_WEEK)
		bey = r.FormValue(query.BOOKING_TIME_END_YEAR)
		bem = r.FormValue(query.BOOKING_TIME_END_MONTH)
	}

	//the booking date
	booking_day, _ := strconv.Atoi(bd)
	booking_week, _ := strconv.Atoi(bw)
	booking_month, _ := strconv.Atoi(bm)
	booking_year, _ := strconv.Atoi(by)

	//start time end time
	booking_start, _ := strconv.Atoi(bs)
	booking_end, _ := strconv.Atoi(be)

	//if reoccuring, end date
	booking_end_day, _ := strconv.Atoi(bed)
	booking_end_week, _ := strconv.Atoi(bew)
	booking_end_month, _ := strconv.Atoi(bem)
	booking_end_year, _ := strconv.Atoi(bey)

	booking_occupancy, _ := strconv.Atoi(bo)

	fmt.Println(">>>>>>>> POST  BOOKING CREATION REOCCURING")
	fmt.Println(booking_reoccuring, bew, bed, bey, bem)

	booking, _ := model.CreateNewBooking(
		room.ID(),
		user.GetID(),
		user.GetAccessLevel(),
		booking_year,
		booking_month-1,
		booking_day-1,
		booking_week,
		booking_occupancy,
		btp,
		bds,
		booking_start, booking_end,
		booking_end_year,
		booking_end_month-1,
		booking_end_day-1,
		booking_end_week,
		booking_reoccuring)

	if !booking_reoccuring {
		booking.SetNonRepeating()
	}
	user.AttachNewBooking(booking)
	room.AttachNewBooking(booking)

	B, err := model.GetBookingsBy_RoomID_YWDH(room.ID(), booking_year, booking_week, booking_day, booking_start)

	if err != nil {
		return nil, err
	}
	return CreateSelectionData(user.GetID(), room, B)
}

func MatchDeletion(b *model.Booking, u *model.User, r *model.Room) bool {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", "CHECKING NOW")

	if b == nil || u == nil || r == nil {
		return false
	}
	fmt.Println(b.RoomID())
	fmt.Println(u.GetID())
	fmt.Println(u.GetAccessLevel())
	fmt.Println(b.UserID())
	fmt.Println(r.ID())

	if r.ID() == b.RoomID() {
		if u.GetAccessLevel() == auth.ACCOUNT_LEVEL_SERVICE_ADMIN ||
			u.GetAccessLevel() == auth.ACCOUNT_LEVEL_TH_PROFESSOR ||
			u.GetAccessLevel() == auth.ACCOUNT_LEVEL_TH_PROFESSOR_ADMIN ||
			u.GetID() == b.UserID() {
			return true
		}
		return false
	}
	return false
}
