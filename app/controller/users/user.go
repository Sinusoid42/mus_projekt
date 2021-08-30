package users

import (
	"fmt"
	"mus_projekt/app/auth"
	"mus_projekt/app/model"
	"mus_projekt/app/model/query"
	"mus_projekt/app/service"
)
/**
	provides all data relevant to the api page
 */
func CreateApiData(u *model.User) map[string]interface{} {
	m := make(map[string]interface{})
	if u == nil {
		return m
	}

	m[query.USER_EMAIL_ADDRESS] = u.GetEmailAddress()
	m[query.ADMIN_LEVEL] = u.GetAccessLevel()
	m[query.USER_BOOKINGS] = createApiDataBookingsByIDs(u.GetBookings())

	return m
}

/**
	creates a map[string]interface{} based on all bookings for which IDS are provided
 */
func createApiDataBookingsByIDs(ids []string) map[string]interface{} {
	m := make(map[string]interface{})
	g := []map[string]interface{}{}
	if ids == nil || len(ids) == 0 {
		return m
	}
	bs, err := model.GetBookingsByIDs(ids)
	if err != nil || len(*bs) == 0 {
		return m
	}
	for i, j := range *bs {
		fmt.Println(i, j)
	}
	fmt.Println(g)

	return m

}

/**
returns api_rooms, api_arduinos, api_templates access
*/
func CheckAccess(u *model.User) (bool, bool, bool) {
	switch u.GetAccessLevel() {
	case auth.ACCOUNT_LEVEL_PUBLIC:
		{
			return false, false, false
		}
	case auth.ACCOUNT_LEVEL_TH_STUDENT:
		{
			return false, false, false
		}
	case auth.ACCOUNT_LEVEL_TH_TUTOR:
		{
			return false, false, false
		}
	case auth.ACCOUNT_LEVEL_TH_PROFESSOR:
		{
			return true, false, false
		}
	case auth.ACCOUNT_LEVEL_TH_PROFESSOR_ADMIN:
		{
			return true, false, false
		}
	case auth.ACCOUNT_LEVEL_TH_ROOM_ADMIN:
		{
			return true, false, false
		}
	case auth.ACCOUNT_LEVEL_SERVICE_ADMIN:
		{
			return true, true, true
		}
	}
	return false, false, false
}

/**
	serves all bookings as json containers relevant to the user page displaying all bookings by a user
 */
func CreateUserPageBookings(id string) map[string]interface{} {
	bks, _ := model.GetBookingsByUserID(id)
	m := make(map[string]interface{})
	m["bookings"] = []map[string]interface{}{}
	for _, j := range *bks {
		q := make(map[string]interface{})
		q[query.BOOKING_USER_ID] = j.UserID()
		r, _ := model.GetRoomBy_RoomID(j.RoomID())
		q[query.LOCATION_ROOM_NAME] = r.Location().Name()
		q[query.BOOKING_TIME_START] = j.StartTime()
		q[query.BOOKING_TIME_END] = j.EndTime()
		q[query.BOOKING_TIME_YEAR] = j.GetTime().Year()
		q[query.BOOKING_TIME_MONTH] = j.GetTime().Month()
		q[query.BOOKING_TIME_DAY] = j.GetTime().Day()
		q[query.BOOKING_TIME_WEEK] = j.GetTime().Week()
		q[query.BOOKING_TOPIC] = j.Topic()

		q[query.BOOKING_ID] = j.ID()
		if j.GetTime().Reoccurring() {
			q[query.BOOKING_TIME_REOCCURING] = true
			q[query.BOOKING_TIME_END_YEAR] = j.GetTime().EndYear()
			q[query.BOOKING_TIME_END_MONTH] = j.GetTime().EndMonth()
			q[query.BOOKING_TIME_END_DAY] = j.GetTime().EndDay()
			q[query.BOOKING_TIME_END_WEEK] = j.GetTime().EndWeek()
		} else {
			q[query.BOOKING_TIME_REOCCURING] = false
		}

		m["bookings"] = append(m["bookings"].([]map[string]interface{}), q)
	}
	return m
}

/*
	with a room booking selected from the user page overview, serves the information for a single booking to be displayed with its form
 */
func CreateSelectionData(bk *model.Booking) map[string]interface{} {
	m := make(map[string]interface{})

	//check max count for that booking time slot

	m[query.BOOKING_ID] = bk.ID()
	m[query.BOOKING_TOPIC] = bk.Topic()
	m[query.BOOKING_DESCRIPTION] = bk.Description()
	m[query.BOOKING_ROOM_ID] = bk.RoomID()
	m[query.BOOKING_TIME_REOCCURING] = false
	m[query.BOOKING_OCCUPANCY] = bk.GetCount()
	if bk.GetTime().Reoccurring() {
		m[query.BOOKING_TIME_REOCCURING] = true

		m[query.BOOKING_TIME_END_YEAR] = bk.GetTime().EndYear()
		m[query.BOOKING_TIME_END_MONTH] = bk.GetTime().EndMonth()
		m[query.BOOKING_TIME_END_WEEK] = bk.GetTime().EndWeek()
		m[query.BOOKING_TIME_END_DAY] = bk.GetTime().EndDay()
	}
	return m
}

/**
	serves the functionalities to create a new upgrade ticket for a user
 */
func CreateUpgradeTicket(u *model.User) {
	service.CreateUpgradeAccessTicket(u.GetID(), u.GetEmailAddress(), u.GetAccessLevel())
}
