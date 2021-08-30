package model

import (
	"errors"
	"fmt"
	"mus_projekt/app/auth"
	"mus_projekt/app/model/query"
	"mus_projekt/app/model/utils"
	"net/http"
	"strconv"
)

type Booking struct {
	_id  string `json:"_id"`  //The UUID of the User within the database
	_rev string `json:"_rev"` //The REV service_id of the user within the database

	booking_id          string       `json:"booking_id"`
	user_id             string       `json:"user_id"`
	room_id             string       `json:"user_id"`
	occupancy           int          `json:"number_of_students"`
	booking_topic       string       `json:"booking_topic"`
	booking_description string       `json:"booking_description"`
	booking_user_access int          `json:"booking_user_access"`
	time                *BookingTime `json:"time"`
}

type BookingTime struct {
	year  int `json:"year"`
	week  int `json:"week"`
	month int `json:"month"`
	day   int `json:"day"`

	start int `json:"start"` //represents the preset hours
	end   int `json:"end"`

	reocurring bool `json:"reocurring"`
	end_year   int  `json:"end_year"`
	end_week   int  `json:"end_week"`
	end_month  int  `json:"end_month"`
	end_day    int  `json:"end_day"`
}

/**
rid : The Room service_id
uid : The User IF
*/
func CreateNewBooking(rid string, uid string, uaccess int, y int, m int, d int, w int, o int, tpc string, desc string, s int, e int, ey int, em int, ed int, ew int, rc bool) (*Booking, error) {
	//TODO Check for existing room and user
	b := Booking{
		_id:                 "",
		_rev:                "",
		booking_id:          utils.GenerateID(),
		user_id:             uid,
		room_id:             rid,
		occupancy:           o,
		booking_topic:       tpc,
		booking_description: desc,
		booking_user_access: uaccess,
		time: &BookingTime{
			year:       y,
			week:       w,
			month:      m,
			day:        d,
			start:      s,
			end:        e,
			reocurring: rc,
			end_year:   ey,
			end_week:   ew,
			end_month:  em,
			end_day:    ed,
		},
	}
	fmt.Println("Ausgabe hier")
	fmt.Println(b)
	fmt.Println(b.time)

	storeBooking(&b)
	return &b, nil
}
func GetBookingByBookingID(id string) (*Booking, error) {
	query := `{"selector":{"booking_id":"` + id + `"}}`

	fmt.Println("The booking id now query", query)
	m, err := booking_DB.QueryJSON(query)
	fmt.Println("\n\n\n", m)
	if len(m) == 0 || err != nil {
		return nil, err
	}
	return m2b(m[0])
}

func GetBookingsByUserID(id string) (*[]*Booking, error) {
	query := `{"selector" : {"booking_user_id" : "` + id + `"}}`

	m, err := booking_DB.QueryJSON(query)
	if err != nil {
		return nil, err
	}
	b := []*Booking{}
	for _, i := range m {
		booking, _ := m2b(i)

		b = append(b, booking)
	}
	return &b, nil
}

func GetBookingsByRoomID(id string) (*[]*Booking, error) {
	return nil, nil
}

func GetBookingsByRoomID_Date(day int, month int, year int, id string) (*[]*[]*Booking, error) {
	return nil, nil
}

func GetBookingsBy_RoomID_YWDH(rid string, y int, w int, d int, h int) (*[]*Booking, error) {
	query :=
		`{` +
			`"selector":{` +
			`"` + query.BOOKING_ROOM_ID + `"` + `:` + `"` + rid + `"` + `,` +
			`"` + query.BOOKING_TIME + `" : {` +
			`"` + query.BOOKING_TIME_YEAR + `"` + `:` + strconv.Itoa(y) + `,` +
			`"` + query.BOOKING_TIME_WEEK + `"` + `:` + strconv.Itoa(w) + `,` +
			`"` + query.BOOKING_TIME_DAY + `"` + `:` + strconv.Itoa(d) + `,` +
			`"` + query.BOOKING_TIME_REOCCURING + `"` + `:` + strconv.FormatBool(false) +
			`}` +
			`}` +
			`}`
	m, err := booking_DB.QueryJSON(query)

	bnks := []*Booking{}
	/*if err != nil {
		return &bnks, err
	}*/

	rec_query := `{
			"selector": {
                "booking_room_id": "` + rid + `",
                "booking_time": {
                    "booking_time_reoccuring": true,
                        "$or": [
                                {
                                    "booking_time_end_year": ` + strconv.Itoa(y+1) + `,
                                    "booking_time_year": ` + strconv.Itoa(y) + `,
									"booking_time_day": ` + strconv.Itoa(d) + `,
                                    "booking_time_week":{
                                        "$lte": ` + strconv.Itoa(w) + `
                                    },
									"booking_time_end": {
                                        "$gte": ` + strconv.Itoa(h) + `
                                    },
									"booking_time_start": {
                                        "$lte": ` + strconv.Itoa(h) + `
                                    }
                                },
                                {
                                    "booking_time_end_year": ` + strconv.Itoa(y) + `,
                                    "booking_time_year": ` + strconv.Itoa(y) + `,
									"booking_time_day": ` + strconv.Itoa(d) + `,
                                    "booking_time_week": {
                                        "$lte": ` + strconv.Itoa(w) + `
                                    },
                                    "booking_time_end_week": {
                                        "$gte": ` + strconv.Itoa(w) + `
                                    },
									"booking_time_end": {
                                        "$gte": ` + strconv.Itoa(h) + `
                                    },
									"booking_time_start": {
                                        "$lte": ` + strconv.Itoa(h) + `
                                    }
                                },
                                {
                                    "booking_time_end_year": ` + strconv.Itoa(y) + `,
                                    "booking_time_year": ` + strconv.Itoa(y-1) + `,
									"booking_time_day": ` + strconv.Itoa(d) + `,
                                    "booking_time_week": {
                                        "$gte": ` + strconv.Itoa(w) + `
                                    },
                                    "booking_time_end_week": {
                                        "$gte": ` + strconv.Itoa(w) + `
                                    },
									"booking_time_end": {
                                        "$gte": ` + strconv.Itoa(h) + `
                                    },
									"booking_time_start": {
                                        "$lte": ` + strconv.Itoa(h) + `
                                    }
                                }
                            ]
                        }
                    }
				}`

	q, err := booking_DB.QueryJSON(rec_query)
	if err != nil {
		return &bnks, err
	}
	for _, i := range m {
		if checkHour(i, h) {
			b, _ := m2b(i)
			bnks = append(bnks, b)
		}
	}
	for _, i := range q {
		if checkHour(i, h) {
			b, _ := m2b(i)
			bnks = append(bnks, b)
		}
	}
	return &bnks, nil
}

func checkHour(m map[string]interface{}, h int) bool {
	return int(m[query.BOOKING_TIME].(map[string]interface{})[query.BOOKING_TIME_START].(float64)+0.5) <= h && int(m[query.BOOKING_TIME].(map[string]interface{})[query.BOOKING_TIME_END].(float64)+0.5) >= h
}

func GetBookingsByRoomID_Week(year int, week int, id string) (*[]*[]*[]*Booking, error) {
	query := `{"selector":{"` +
		query.BOOKING_ROOM_ID + `" : "` + id + `","` +
		query.BOOKING_TIME + `" :{"` +
		query.BOOKING_TIME_WEEK + `": ` + strconv.Itoa(week) + `, "` +
		query.BOOKING_TIME_REOCCURING + `": ` + strconv.FormatBool(false) + `, "` +
		query.BOOKING_TIME_YEAR + `": ` + strconv.Itoa(year) + `}}}`
	m, err := booking_DB.QueryJSON(query)
	if err != nil {
		return nil, err
	}

	//TODO Sort By Time & By Day => Pack data into useable information about the room
	//Fetch the data to the client and build the booking creation now

	//TODO get all the reoccuring bookings

	rec_query := `{
			"selector": {
                "booking_room_id": "` + id + `",
                "booking_time": {
                    "booking_time_reoccuring": true,
                        "$or": [
                                {
                                    "booking_time_end_year": ` + strconv.Itoa(year+1) + `,
                                    "booking_time_year": ` + strconv.Itoa(year) + `,
                                    "booking_time_week": {
                                        "$lte": ` + strconv.Itoa(week) + `
                                    }
                                },
                                {
                                    "booking_time_end_year": ` + strconv.Itoa(year) + `,
                                    "booking_time_year": ` + strconv.Itoa(year) + `,
                                    "booking_time_week": {
                                        "$lte": ` + strconv.Itoa(week) + `
                                    },
                                    "booking_time_end_week": {
                                        "$gte": ` + strconv.Itoa(week) + `
                                    }
                                },
                                {
                                    "booking_time_end_year": ` + strconv.Itoa(year) + `,
                                    "booking_time_year": ` + strconv.Itoa(year-1) + `,
                                    "booking_time_week": {
                                        "$gte": ` + strconv.Itoa(week) + `
                                    },
                                    "booking_time_end_week": {
                                        "$gte": ` + strconv.Itoa(week) + `
                                    }
                                }
                            ]
                        }
                    }
				}`

	q, err := booking_DB.QueryJSON(rec_query)
	wk := []*[]*[]*Booking{}
	mo := []*[]*Booking{}
	di := []*[]*Booking{}
	mi := []*[]*Booking{}
	do := []*[]*Booking{}
	fr := []*[]*Booking{}
	i := 0

	for i = 0; i < 14; i++ {
		mo = append(mo, &[]*Booking{})
		di = append(di, &[]*Booking{})
		mi = append(mi, &[]*Booking{})
		do = append(do, &[]*Booking{})
		fr = append(fr, &[]*Booking{})
	}

	wk = append(wk, &mo)
	wk = append(wk, &di)
	wk = append(wk, &mi)
	wk = append(wk, &do)
	wk = append(wk, &fr)
	for _, b := range q {
		booking, _ := m2b(b)
		switch booking.time.day {
		case 0:
			{
				l := append(*mo[booking.time.start], booking)
				mo[booking.time.start] = &l
				break
			}
		case 1:
			{
				l := append(*di[booking.time.start], booking)
				di[booking.time.start] = &l
				break
			}
		case 2:
			{
				l := append(*mi[booking.time.start], booking)
				mi[booking.time.start] = &l
				break
			}
		case 3:
			{
				l := append(*do[booking.time.start], booking)
				do[booking.time.start] = &l
				break
			}
		case 4:
			{
				l := append(*fr[booking.time.start], booking)
				fr[booking.time.start] = &l
				break
			}
		}
	}
	for _, b := range m {
		booking, _ := m2b(b)
		switch booking.time.day {
		case 0:
			{
				l := append(*mo[booking.time.start], booking)
				mo[booking.time.start] = &l
				break
			}
		case 1:
			{
				l := append(*di[booking.time.start], booking)
				di[booking.time.start] = &l
				break
			}
		case 2:
			{
				l := append(*mi[booking.time.start], booking)
				mi[booking.time.start] = &l
				break
			}
		case 3:
			{
				l := append(*do[booking.time.start], booking)
				do[booking.time.start] = &l
				break
			}
		case 4:
			{
				l := append(*fr[booking.time.start], booking)
				fr[booking.time.start] = &l
				break
			}
		}
	}
	return &wk, nil
}

func b2m(b *Booking) map[string]interface{} {
	m := make(map[string]interface{})
	m[query.DB_ID] = b._id
	m[query.DB_REV] = b._rev
	m[query.BOOKING_ID] = b.booking_id
	m[query.BOOKING_ROOM_ID] = b.room_id
	m[query.BOOKING_USER_ID] = b.user_id
	m[query.BOOKING_OCCUPANCY] = b.occupancy
	m[query.BOOKING_TOPIC] = b.booking_topic
	m[query.BOOKING_USER_ACCESS] = b.booking_user_access
	m[query.BOOKING_DESCRIPTION] = b.booking_description

	t := make(map[string]interface{})
	t[query.BOOKING_TIME_YEAR] = b.time.year
	t[query.BOOKING_TIME_MONTH] = b.time.month
	t[query.BOOKING_TIME_DAY] = b.time.day
	t[query.BOOKING_TIME_WEEK] = b.time.week

	t[query.BOOKING_TIME_REOCCURING] = b.time.reocurring
	t[query.BOOKING_TIME_START] = b.time.start
	t[query.BOOKING_TIME_END] = b.time.end

	t[query.BOOKING_TIME_END_YEAR] = b.time.end_year
	t[query.BOOKING_TIME_END_MONTH] = b.time.end_month
	t[query.BOOKING_TIME_END_DAY] = b.time.end_day
	t[query.BOOKING_TIME_END_WEEK] = b.time.end_week

	m[query.BOOKING_TIME] = t
	return m
}

func m2b(m map[string]interface{}) (*Booking, error) {
	b := Booking{
		_id:                 m[query.DB_ID].(string),
		_rev:                m[query.DB_REV].(string),
		booking_id:          m[query.BOOKING_ID].(string),
		user_id:             m[query.BOOKING_USER_ID].(string),
		room_id:             m[query.BOOKING_ROOM_ID].(string),
		occupancy:           int(m[query.BOOKING_OCCUPANCY].(float64) + 0.5),
		booking_topic:       m[query.BOOKING_TOPIC].(string),
		booking_description: m[query.BOOKING_DESCRIPTION].(string),
		booking_user_access: int(m[query.BOOKING_USER_ACCESS].(float64) + 0.5),
		time: &BookingTime{
			year:       int(m[query.BOOKING_TIME].(map[string]interface{})[query.BOOKING_TIME_YEAR].(float64) + 0.5),
			month:      int(m[query.BOOKING_TIME].(map[string]interface{})[query.BOOKING_TIME_MONTH].(float64) + 0.5),
			day:        int(m[query.BOOKING_TIME].(map[string]interface{})[query.BOOKING_TIME_DAY].(float64) + 0.5),
			end_year:   int(m[query.BOOKING_TIME].(map[string]interface{})[query.BOOKING_TIME_END_YEAR].(float64) + 0.5),
			end_month:  int(m[query.BOOKING_TIME].(map[string]interface{})[query.BOOKING_TIME_END_MONTH].(float64) + 0.5),
			end_day:    int(m[query.BOOKING_TIME].(map[string]interface{})[query.BOOKING_TIME_END_DAY].(float64) + 0.5),
			start:      int(m[query.BOOKING_TIME].(map[string]interface{})[query.BOOKING_TIME_START].(float64) + 0.5),
			end:        int(m[query.BOOKING_TIME].(map[string]interface{})[query.BOOKING_TIME_END].(float64) + 0.5),
			week:       int(m[query.BOOKING_TIME].(map[string]interface{})[query.BOOKING_TIME_WEEK].(float64) + 0.5),
			end_week:   int(m[query.BOOKING_TIME].(map[string]interface{})[query.BOOKING_TIME_END_WEEK].(float64) + 0.5),
			reocurring: m[query.BOOKING_TIME].(map[string]interface{})[query.BOOKING_TIME_REOCCURING].(bool),
		},
	}
	return &b, nil
}

func (b *Booking) Save() error {
	_id, _rev, err := booking_DB.Save(b2m(b), nil)
	b._id = _id //should here be the same as b4
	b._rev = _rev
	if err != nil {
		return err
	}
	return nil
}

func storeBooking(b *Booking) error {
	m := b2m(b)
	delete(m, query.DB_ID) //couchdb preparation
	delete(m, query.DB_REV)

	_id, _rev, err := booking_DB.Save(m, nil)
	b._id = _id //should here be the same as b4
	b._rev = _rev

	if err != nil {
		return err
	}
	return nil
}

func GetBookingByID(id string) (*Booking, error) {
	query := `{"selector":{"booking_id":"` + id + `"}}`
	b, err := booking_DB.QueryJSON(query)
	if err != nil || len(b) == 0 {
		return nil, errors.New("User not found")
	}
	bk, err := m2b(b[0])
	return bk, nil
}

func GetBookingsByIDs(ids []string) (*[]*Booking, error) {
	b := []*Booking{}
	for _, i := range ids {
		bk, err := GetBookingByID(i)
		if err != nil {
			return nil, err
		}
		b = append(b, bk)
	}
	return &b, nil
}

func (b *Booking) ID() string {
	return b.booking_id
}

func (b *Booking) SetNonRepeating() {
	b.time.reocurring = false
	b.time.end_day = -1
	b.time.end_year = -1
	b.time.end_week = -1
	b.time.end_month = -1
}

func (b *Booking) Topic() string {
	return b.booking_topic
}

func (b *Booking) Description() string {
	return b.booking_description
}

func (b *Booking) UserID() string {
	return b.user_id
}

func (b *Booking) RoomID() string {
	return b.room_id
}

func (b *Booking) GetTime() *BookingTime {
	return b.time
}

func (b *Booking) StartTime() int {
	return b.time.start
}

func (b *Booking) EndTime() int {
	return b.time.end
}

func (b *Booking) Remove() error {
	u, _ := GetUserByID(b.user_id)
	u.RemoveBookingByID(b.booking_id)
	r, _ := GetRoomBy_RoomID(b.room_id)
	r.RemoveBookingByID(b.booking_id)
	err := booking_DB.Delete(b._id)
	return err
}

func (b *Booking) GetCount() int {
	return b.occupancy
}

func (b *Booking) Duration() int {
	return b.time.end - b.time.start
}

func (b *Booking) ComparePriority(p int) int {
	switch b.booking_user_access {
	case auth.ACCOUNT_LEVEL_PUBLIC:
		{
			switch p {
			case auth.ACCOUNT_LEVEL_PUBLIC:
				{
					return 0
				}
			case auth.ACCOUNT_LEVEL_TH_STUDENT:
				{
					return -1
				}
			case auth.ACCOUNT_LEVEL_TH_TUTOR:
				{
					return -1
				}
			case auth.ACCOUNT_LEVEL_TH_PROFESSOR:
				{
					return -1
				}
			case auth.ACCOUNT_LEVEL_TH_PROFESSOR_ADMIN:
				{
					return -1
				}
			case auth.ACCOUNT_LEVEL_TH_ROOM_ADMIN:
				{
					return -1
				}
			case auth.ACCOUNT_LEVEL_SERVICE_ADMIN:
				{
					return -1
				}
			}
		}
	case auth.ACCOUNT_LEVEL_TH_STUDENT:
		{
			switch p {
			case auth.ACCOUNT_LEVEL_PUBLIC:
				{
					return 1
				}
			case auth.ACCOUNT_LEVEL_TH_STUDENT:
				{
					return 0
				}
			case auth.ACCOUNT_LEVEL_TH_TUTOR:
				{
					return -1
				}
			case auth.ACCOUNT_LEVEL_TH_PROFESSOR:
				{
					return -1
				}
			case auth.ACCOUNT_LEVEL_TH_PROFESSOR_ADMIN:
				{
					return -1
				}
			case auth.ACCOUNT_LEVEL_TH_ROOM_ADMIN:
				{
					return -1
				}
			case auth.ACCOUNT_LEVEL_SERVICE_ADMIN:
				{
					return -1
				}
			}
		}
	case auth.ACCOUNT_LEVEL_TH_TUTOR:
		{
			switch p {
			case auth.ACCOUNT_LEVEL_PUBLIC:
				{
					return 1
				}
			case auth.ACCOUNT_LEVEL_TH_STUDENT:
				{
					return 1
				}
			case auth.ACCOUNT_LEVEL_TH_TUTOR:
				{
					return 0
				}
			case auth.ACCOUNT_LEVEL_TH_PROFESSOR:
				{
					return -1
				}
			case auth.ACCOUNT_LEVEL_TH_PROFESSOR_ADMIN:
				{
					return -1
				}
			case auth.ACCOUNT_LEVEL_TH_ROOM_ADMIN:
				{
					return -1
				}
			case auth.ACCOUNT_LEVEL_SERVICE_ADMIN:
				{
					return -1
				}
			}
		}
	case auth.ACCOUNT_LEVEL_TH_PROFESSOR:
		{
			switch p {
			case auth.ACCOUNT_LEVEL_PUBLIC:
				{
					return 1
				}
			case auth.ACCOUNT_LEVEL_TH_STUDENT:
				{
					return 1
				}
			case auth.ACCOUNT_LEVEL_TH_PROFESSOR:
				{
					return 0
				}
			case auth.ACCOUNT_LEVEL_TH_PROFESSOR_ADMIN:
				{
					return -1
				}
			case auth.ACCOUNT_LEVEL_TH_ROOM_ADMIN:
				{
					return -1
				}
			case auth.ACCOUNT_LEVEL_SERVICE_ADMIN:
				{
					return -1
				}
			}
		}
	case auth.ACCOUNT_LEVEL_TH_PROFESSOR_ADMIN:
		{
			switch p {
			case auth.ACCOUNT_LEVEL_PUBLIC:
				{
					return 1
				}
			case auth.ACCOUNT_LEVEL_TH_STUDENT:
				{
					return 1
				}
			case auth.ACCOUNT_LEVEL_TH_TUTOR:
				{
					return 1
				}
			case auth.ACCOUNT_LEVEL_TH_PROFESSOR:
				{
					return -1
				}
			case auth.ACCOUNT_LEVEL_TH_PROFESSOR_ADMIN:
				{
					return 0
				}
			case auth.ACCOUNT_LEVEL_TH_ROOM_ADMIN:
				{
					return -1
				}
			case auth.ACCOUNT_LEVEL_SERVICE_ADMIN:
				{
					return -1
				}
			}
		}
	case auth.ACCOUNT_LEVEL_TH_ROOM_ADMIN:
		{
			switch p {
			case auth.ACCOUNT_LEVEL_PUBLIC:
				{
					return 1
				}
			case auth.ACCOUNT_LEVEL_TH_STUDENT:
				{
					return 1
				}
			case auth.ACCOUNT_LEVEL_TH_TUTOR:
				{
					return 1
				}
			case auth.ACCOUNT_LEVEL_TH_PROFESSOR:
				{
					return 1
				}
			case auth.ACCOUNT_LEVEL_TH_PROFESSOR_ADMIN:
				{
					return 0
				}
			case auth.ACCOUNT_LEVEL_TH_ROOM_ADMIN:
				{
					return -1
				}
			case auth.ACCOUNT_LEVEL_SERVICE_ADMIN:
				{
					return -1
				}
			}
		}
	case auth.ACCOUNT_LEVEL_SERVICE_ADMIN:
		{
			switch p {
			case auth.ACCOUNT_LEVEL_PUBLIC:
				{
					return 1
				}
			case auth.ACCOUNT_LEVEL_TH_STUDENT:
				{
					return 1
				}
			case auth.ACCOUNT_LEVEL_TH_TUTOR:
				{
					return 1
				}
			case auth.ACCOUNT_LEVEL_TH_PROFESSOR:
				{
					return 1
				}
			case auth.ACCOUNT_LEVEL_TH_PROFESSOR_ADMIN:
				{
					return 1
				}
			case auth.ACCOUNT_LEVEL_TH_ROOM_ADMIN:
				{
					return 1
				}
			case auth.ACCOUNT_LEVEL_SERVICE_ADMIN:
				{
					return 0
				}
			}
		}
	}
	return -2
}

func (b *Booking) UserAccess() int {
	return b.booking_user_access
}

func (b *Booking) UserName() string {
	u, _ := GetUserByID(b.user_id)
	return u.user_name
}

func (t *BookingTime) Year() int {
	return t.year
}

func (t BookingTime) Month() int {
	return t.month
}

func (t *BookingTime) Day() int {
	return t.day
}

func (t *BookingTime) Week() int {
	return t.week
}

func (t *BookingTime) EndWeek() int {
	return t.end_week
}

func (t *BookingTime) EndYear() int {
	return t.end_year
}

func (t BookingTime) EndMonth() int {
	return t.end_month
}

func (t *BookingTime) EndDay() int {
	return t.end_day
}

func (t *BookingTime) Reoccurring() bool {
	return t.reocurring
}

func (b *Booking) SaveBooking(r *http.Request) (map[string]interface{}, error) {

	//fetch data from query

	br := r.FormValue(query.BOOKING_TIME_REOCCURING)
	btp := r.FormValue(query.BOOKING_TOPIC)
	bds := r.FormValue(query.BOOKING_DESCRIPTION)

	bo := r.FormValue(query.BOOKING_OCCUPANCY)

	fmt.Println("ANKOMMENDE LEUTE ", bo)

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

	//the booking date already exists, unchangeable

	//start time end time

	//if reoccuring, end date
	booking_end_day, _ := strconv.Atoi(bed)
	booking_end_week, _ := strconv.Atoi(bew)
	booking_end_month, _ := strconv.Atoi(bem)
	booking_end_year, _ := strconv.Atoi(bey)

	booking_occupancy, _ := strconv.Atoi(bo)

	b.booking_topic = btp
	b.booking_description = bds
	b.time.end_week = booking_end_week
	b.time.end_month = booking_end_month
	b.time.end_year = booking_end_year
	b.time.end_day = booking_end_day
	b.time.reocurring = booking_reoccuring
	b.occupancy = booking_occupancy

	b.Save()
	m := make(map[string]interface{})
	return m, nil
}
