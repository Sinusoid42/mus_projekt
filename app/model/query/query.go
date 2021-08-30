package query

/**
Status Information for current session
*/

const SUCCESS string = "success"
const AUTH string = "auth"
const ADMIN_LEVEL string = "admin_level"

const DB_ID string = "_id"
const DB_REV string = "_rev"

/**
User Information
*/

const USER_EXIST string = "user_exist"
const USER_PW string = "user_password"
const USER_NM string = "user_name"
const USER_PRIVILEGES string = "admin_access_level"
const USER_BOOKINGS string = "user_bookings"
const USER_DB_ID string = "_id"
const USER_DB_REV string = "_rev"
const USER_EMAIL_ADDRESS string = "user_email_address"
const USER_AUTHENTICATED string = "auth"
const USER_CONFIRMED string = "confirmed"

/**
Room Information
*/

const ROOM_ID string = "room_id"
const ROOM_DB_ID string = "_id"
const ROOM_DB_REV string = "_rev"
const ROOM_NAME string = "room_name"
const ROOM_NAME_MISC string = "room_name_misc"
const ROOM_BOOKING_IDS string = "room_booking_ids"
const ROOM_MAX_CAPACITY string = "room_max_capacity"
const ROOM_MAX_BOOKING_DURATION string = "room_max_booking_duration"
const ROOM_LOCATION string = "room_location"
const ROOM_BOOKABLE string = "room_bookable"
const ROOM_ARDUINO_AVAILBLE string = "room_arduino_available"

/**
Location Information
*/

const LOCATION_FLOOR string = "location_room_floor_level"
const LOCATION_ROOM_CORRIDOR string = "location_room_corridor"
const LOCATION_ROOM_NUMBER string = "location_room_number"
const LOCATION_ROOM_NAME string = "location_room_name"

/**
Template Information
*/
const TEMPLATE_ID string = "template_id"
const TEMPLATE_NAME string = "template_name"
const TEMPLATE_ELEMENTS string = "elements"

const ELEMENT_ID string = "element_id"
const ELEMENT_CONTENT string = "content"
const ELEMENT_CONTENT_STATIC string = "content_static"
const ELEMENT_POSITION_X string = "x"
const ELEMENT_POSITION_Y string = "y"
const ELEMENT_POSITION_W string = "w"
const ELEMENT_POSITION_H string = "h"
const ELEMENT_COLOR string = "color"
const ELEMENT_FILL_COLOR string = "fill_color"
const ELEMENT_FONT_SIZE string = "font_size"
const ELEMENT_PIXEL_SIZE string = "pixel_size"
const ELEMENT_PIXEL_STYLE string = "pixel_style"
const ELEMENT_FORM string = "form"

/**
Booking Information
*/

const BOOKING_ID string = "booking_id"
const BOOKING_USER_ID string = "booking_user_id"
const BOOKING_ROOM_ID string = "booking_room_id"
const BOOKING_OCCUPANCY string = "booking_occupancy"
const BOOKING_TOPIC string = "booking_topic"
const BOOKING_DESCRIPTION string = "booking_description"
const BOOKING_USER_ACCESS string = "booking_user_access"
const BOOKING_TIME string = "booking_time"

const BOOKING_TIME_YEAR string = "booking_time_year"
const BOOKING_TIME_MONTH string = "booking_time_month"
const BOOKING_TIME_DAY string = "booking_time_day"
const BOOKING_TIME_WEEK string = "booking_time_week"

const BOOKING_TIME_REOCCURING string = "booking_time_reoccuring"
const BOOKING_TIME_END_YEAR string = "booking_time_end_year"
const BOOKING_TIME_END_MONTH string = "booking_time_end_month"
const BOOKING_TIME_END_DAY string = "booking_time_end_day"
const BOOKING_TIME_END_WEEK string = "booking_time_end_week"

const BOOKING_TIME_START string = "booking_time_start"
const BOOKING_TIME_END string = "booking_time_end"

const CALENDER_YEAR string = "selected_calender_year"
const CALENDER_WEEK string = "selected_calender_week"
const CALENDER_DAY string = "selected_calender_day"
const CALENDER_HOUR string = "selected_calender_hour"

const HTML_API_ROOMS_ACCESS string = "api_rooms_access"
const HTML_API_ARDUINOS_ACCESS string = "api_arduinos_access"
const HTML_API_TEMPLATES_ACCESS string = "api_templates_access"

const SERVICE_UPGRADE_HELPER_ID string = "upgrade_helper_id"
