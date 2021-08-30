package templates

import (
	"mus_projekt/app/model"
	"mus_projekt/app/model/query"
)

const DYNAMIC_ROOM_NAME string = "room_name"
const DYNAMIC_ROOM_MISC_NAME string = "room_name_misc"
const DYNAMIC_ROOM_LOCATION_NAME string = "room_location_name"
const DYNAMIC_ROOM_LOCATION_FLOOR string = "room_location_floor"
const DYNAMIC_ROOM_LOCATION_CORRIDOR string = "room_location_corridor"
const DYNAMIC_ROOM_LOCATION_ROOM_NUMBER string = "room_location_number"
const DYNAMIC_ROOM_MAX_CAPACITY string = "room_max_capacity"
const DYNAMIC_ROOM_CURRENT_OCCUPANCY string = "room_current_occupancy"

/**
Current Booking

If current Booking is nil, data will be set to default
*/
const DYNAMIC_BOOKING_USER_NAME string = "booking_user_name"
const DYNAMIC_BOOKING_TOPIC string = "booking_topic"
const DYNAMIC_BOOKING_START_TIME string = "booking_start_time"
const DYNAMIC_BOOKING_END_TIME string = "booking_end_time"
const DYNAMIC_BOOKING_NUMBER_OF_STUDENTTS string = "booking_occupancy"

const DYNAMIC_BOOKING_NEXT_USER_NAME string = "booking_next_user_name"
const DYNAMIC_BOOKING_NEXT_TOPIC string = "booking_next_topic"
const DYNAMIC_BOOKING_NEXT_START_TIME string = "booking_next_start_time"
const DYNAMIC_BOOKING_NEXT_END_TIME string = "booking_next_end_time"
const DYNAMIC_BOOKING_NEXT_MUMBER_OF_STUDENTS string = "booking_next_occupancy"

const DC_ERROR_CODE_ROOM_UNAVAILABLE string = "ERR_101"
const DC_ERROR_CODE_BOOKINGS_UNAVAILABLE string = "ERR_102"
const DC_ERROR_CODE_DYNAMIC_CONTENT_UNAVAILABLE string = "ERR_103"

const DC_CODE_BOOKINGS_UNAVAILABLE string = "CODE_101"

func CreateNew() *model.Template {
	t, _ := model.CreateTemplate()
	return t
}

func CreateAPIData(t *model.Template) map[string]interface{} {
	j := make(map[string]interface{})

	j[query.TEMPLATE_ID] = t.GetID()
	j[query.TEMPLATE_NAME] = t.GetName()

	l := []map[string]interface{}{}

	for _, e := range *t.GetElements() {

		l = append(l, CreateAPIDataElement(t.GetID(), e))
	}

	j[query.TEMPLATE_ELEMENTS] = l
	return j
}

func GetAllTemplates() (map[string]interface{}, error) {
	m := []map[string]interface{}{}
	k := make(map[string]interface{})
	t, err := model.GetAllTemplates()
	if err != nil {
		return k, err
	}
	for _, i := range t {
		m = append(m, CreateAPIData(i))
	}
	k["templates"] = m
	return k, err
}

func CreateAPIDataElement(id string, e *model.Element) map[string]interface{} {
	m := make(map[string]interface{})
	m[query.TEMPLATE_ID] = id
	m[query.ELEMENT_ID] = e.GetElementID()
	m[query.ELEMENT_CONTENT] = e.GetContent()
	m[query.ELEMENT_CONTENT_STATIC] = e.GetContentStatic()
	m[query.ELEMENT_POSITION_X] = e.GetX()
	m[query.ELEMENT_POSITION_Y] = e.GetY()
	m[query.ELEMENT_POSITION_W] = e.GetW()
	m[query.ELEMENT_POSITION_H] = e.GetH()
	m[query.ELEMENT_COLOR] = e.GetColor()
	m[query.ELEMENT_FILL_COLOR] = e.GetFillColor()
	m[query.ELEMENT_FONT_SIZE] = e.GetFontSize()
	m[query.ELEMENT_PIXEL_SIZE] = e.GetPixelSize()
	m[query.ELEMENT_PIXEL_STYLE] = e.GetStyle()
	m[query.ELEMENT_FORM] = e.GetForm()
	return m
}

func CheckTemplateById(id string) (bool, *model.Template, error) {
	if len(id) == 0 {
		return false, nil, nil
	}
	return model.Check_Template(id)
}

/**
TODO mit jannis
*/
func GetContentDynamic(arduino *model.Arduino, r *model.Room, content string) string {
	if r == nil {
		return DC_ERROR_CODE_ROOM_UNAVAILABLE
	}
	b, j := r.GetCurrentBookings()
	if j != nil {
		return DC_ERROR_CODE_BOOKINGS_UNAVAILABLE
	}
	if len(*b) == 0 {
		return DC_CODE_BOOKINGS_UNAVAILABLE
	}
	e := ""
	k := getHighestPriorityBooking(b)
	switch content {
	case "current_booking_user_name":
		{
			e = k.UserName()
			return e
		}

	}
	return DC_ERROR_CODE_DYNAMIC_CONTENT_UNAVAILABLE
}

func getHighestPriorityBooking(b *[]*model.Booking) *model.Booking {
	var e = -1
	var p = 0
	for i, k := range *b {
		if k.ComparePriority(p) == 1 || k.ComparePriority(p) == 0 {
			e = i
			p = k.UserAccess()
		}
	}
	if e != -1 {
		bk := *b
		return bk[e]
	}
	return nil
}
