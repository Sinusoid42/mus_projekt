package model

import (
	"fmt"
	"log"
	"mus_projekt/app/model/utils"
)

type Arduino struct {
	_id  string `json:"_id"`  //The UUID of the User within the database
	_rev string `json:"_rev"` //The REV service_id of the user within the database

	microcontroller_id string `json:"microcontroller_id"`

	room_id       string `json:"room_id"`
	room_location string `json:"room_location"`

	microcontroller_ip_adress string `json:"microcontroller_ip_address"`
	microcontroller_inet_port string `json:"microcontroller_inet_port"`
	microcontroller_password  string `json:"microcontroller_password"`

	microcontroller_color_options []string `json:"microcontroller_color_option"`

	microcontroller_type string `json:"microcontroller_type"`

	microcontroller_firmware string `json:"microcontroller_firmware"`

	microcontroller_active_template string `json:"microcontroller_active_template"`
}

func GetArduinoBy_ArduinoID(id string) (*Arduino, error) {
	query := `{"selector":{"microcontroller_id":"%s"}}`
	m, err := arduino_DB.QueryJSON(fmt.Sprintf(query, id))
	if err != nil || len(m) == 0 {
		fmt.Println(err)
		return nil, err
	}
	a := M2A(m[0])
	return a, nil
}
func M2A(m map[string]interface{}) *Arduino {
	if m == nil {
		return nil
	}
	a := Arduino{
		_id:                             m["_id"].(string),
		_rev:                            m["_rev"].(string),
		microcontroller_id:              m["microcontroller_id"].(string),
		room_id:                         m["room_id"].(string),
		room_location:                   m["room_location"].(string),
		microcontroller_ip_adress:       m["microcontroller_ip_address"].(string),
		microcontroller_inet_port:       m["microcontroller_inet_port"].(string),
		microcontroller_password:        m["microcontroller_password"].(string),
		microcontroller_type:            m["microcontroller_type"].(string),
		microcontroller_firmware:        m["microcontroller_firmware"].(string),
		microcontroller_active_template: m["microcontroller_active_template"].(string),
	}
	return &a
}

func (A *Arduino) A2M() map[string]interface{} {
	if A == nil {
		return nil
	}
	m := make(map[string]interface{})
	m["_id"] = A._id
	m["_rev"] = A._rev
	m["microcontroller_id"] = A.microcontroller_id
	m["room_id"] = A.room_id
	m["room_location"] = A.room_location
	m["microcontroller_ip_address"] = A.microcontroller_ip_adress
	m["microcontroller_inet_port"] = A.microcontroller_inet_port
	m["microcontroller_password"] = A.microcontroller_password
	m["microcontroller_type"] = A.microcontroller_type
	m["microcontroller_firmware"] = A.microcontroller_firmware
	m["microcontroller_active_template"] = A.microcontroller_active_template
	return m
}
func CreateMicrocontroller() (*Arduino, error) {
	a := Arduino{
		"",
		"",
		"nil",
		"nil",
		"nil",
		"nil",
		"nil",
		"nil",
		[]string{},
		"nil",
		"nil",
		"nil",
	}
	a.microcontroller_id = utils.GenerateID()
	m := a.A2M()
	delete(m, "_id")
	delete(m, "_rev")
	id, rev, err := arduino_DB.Save(m, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	a._id = id
	a._rev = rev
	return &a, nil
}

func GetAllActiveArduinos() []map[string]interface{} {

	i, _ := arduino_DB.DocIDs()
	m := make([]map[string]interface{}, len(i))
	for k, v := range i {
		d, _ := arduino_DB.Get(v, nil)
		m[k] = d
	}
	return m
}

func RemoveArduinoByID(id string) bool {

	a, e := GetArduinoBy_ArduinoID(id)
	if e != nil {
		return false
	}
	arduino_DB.Delete(a._id)
	return true
}

func (a *Arduino) StoreUpdate(update map[string]interface{}) bool {

	a.microcontroller_id = update["microcontroller_id"].(string)
	a.room_id = update["room_id"].(string)
	a.microcontroller_active_template = update["microcontroller_active_template"].(string)
	a.microcontroller_password = update["microcontroller_password"].(string)
	a.microcontroller_type = update["microcontroller_type"].(string)
	a.microcontroller_inet_port = update["microcontroller_inet_port"].(string)
	a.microcontroller_ip_adress = update["microcontroller_ip_address"].(string)
	a.room_location = update["room_location"].(string)
	a.microcontroller_firmware = update["microcontroller_firmware"].(string)

	if len(a.room_id) > 0 {
		r, _ := GetRoomBy_RoomID(a.room_id)
		r.microcontroller_available = true
		r.Save()

	}
	_, _, err := arduino_DB.Save(a.A2M(), nil)
	if err != nil {
		return false
	}

	return true
}

func (a *Arduino) GetTemplateName() string {
	return a.microcontroller_active_template
}

func (a *Arduino) GetRoomID() string {
	return a.room_id
}
