package auth

/*
	The authenticator is managing the access
	to privately and admin interactions from
	http requests

	@author ben
*/

import (
	"github.com/gorilla/sessions"
	"math/rand"
)

var cookieStore *sessions.CookieStore
var session_cookie string = "room_occupancy_system"

const USER_COOKIE_AUTH string = "authenticated"
const USER_ID string = "user_id"
const USER_ACCESS_LEVEL string = "admin_access_level"

/**

Access priorities

*/

/**
Priority is unable to create Bookings, can see rooms and self user account
*/
const ACCOUNT_LEVEL_PUBLIC = 0

/**
Priority has the ability to create Bookings and remove its own bookings
cannot exeed bookings if max room occupancy number is exeeding

bookings by account level 1 can be removed by admins and professors
*/
const ACCOUNT_LEVEL_TH_STUDENT = 1

/**
Priority has the ability to create bookings and reoccuring bookings as a tutor
until a fixed date
Priority cannot remove other students bookings from the calender
*/

const ACCOUNT_LEVEL_TH_TUTOR = 2

/**
Priority has the ability to create bookings and reocurring bookings as a professor

Has the ability to remove bookings from days or within timeslots, if rooms are blocked


*/
const ACCOUNT_LEVEL_TH_PROFESSOR = 3

/**
SAME AS LEVEL 3

This priority gives the user access to server api insights

With this priority the user may see, but not create or store anything within the api site
*/
const ACCOUNT_LEVEL_TH_PROFESSOR_ADMIN = 4

/**
This priority can create rooms from the api section and change rooms from the backend

*/
const ACCOUNT_LEVEL_TH_ROOM_ADMIN = 5

/**
SERVICE ADMIN

This priority enables everything

*/

const ACCOUNT_LEVEL_SERVICE_ADMIN = -1

func init() {
	key := make([]byte, 32)
	rand.Read(key)
	cookieStore = sessions.NewCookieStore(key)
}

func GetCookieStore() *sessions.CookieStore {
	return cookieStore
}

func GetSessionCookie() string {
	return session_cookie
}

/*
	Decodes the user_password from rot13 to utf8 visible in pipeline for encrypting with bcrypt
*/
func DecodeRot13(user_password_shift_13 string) string {
	var i = 0
	var s = ""
	for ; i < len(user_password_shift_13); i++ {
		var j = uint32(byte(user_password_shift_13[i]))
		j -= 13
		if j < 0 {
			j += 255
		}
		s += string(j & 0xff)
	}
	return s
}
