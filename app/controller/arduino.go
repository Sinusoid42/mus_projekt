package controller

import (
	"encoding/json"
	"fmt"
	"mus_projekt/app/controller/protocol"
	"mus_projekt/app/controller/templates"
	"mus_projekt/app/model"
	"mus_projekt/app/model/query"
	"net/http"
	"time"
)

/*
	The arduino.go controller is managing all
	incoming http requests for the arduino
	the path variables will be set by this go script

	@author ben

*/

//const registered_arduinos _ARDUINO = database.re

func Arduino(w http.ResponseWriter, r *http.Request) {

	fmt.Println("\nhttp:API:Arduino")

	if r.Method == http.MethodGet {
		ExecuteTemplate("arduino_settings_template.tmpl", w, r, nil)
		return
	}
	if r.Method == protocol.MethodFetch {
		//m := testdata()
		a := model.GetAllActiveArduinos()
		fmt.Println(a)
		m := make(map[string]interface{})
		m["data"] = a
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
	if r.Method == protocol.MethodSelect {
		i := r.FormValue("microcontroller_id")
		fmt.Println("The Arduino service_id", i)
		a, _ := model.GetArduinoBy_ArduinoID(i)

		m := a.A2M()
		rm, err := model.GetRoomBy_RoomID(a.GetRoomID())
		if err != nil || rm == nil {
			m["room_location"] = ""

		} else {
			m["success"] = true
			m["room_location"] = rm.Location().Name()
		}
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
	if r.Method == http.MethodPost {
		//Create new arduino for database
		a, _ := model.CreateMicrocontroller()
		m := a.A2M()
		fmt.Println(m)
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
	if r.Method == http.MethodDelete {
		fmt.Println("http:Arduini:DELETE")
		//delete the arduino from the database and send confirmation success to the client
		if e := r.FormValue("microcontroller_id"); &e != nil {
			success := model.RemoveArduinoByID(e)
			m := make(map[string]interface{})
			m["success"] = success
			b, _ := json.MarshalIndent(m, " ", "  ")
			//response := `{"success" : "` + strconv.FormatBool(success) + `"}`
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}
	}
	if r.Method == http.MethodPut {
		fmt.Println("http:Arduino:PUT")
		q := make(map[string]interface{})
		q["id"] = r.FormValue("id")
		q["microcontroller_id"] = r.FormValue("microcontroller_id")
		q["microcontroller_ip_address"] = r.FormValue("microcontroller_ip_address")
		q["microcontroller_inet_port"] = r.FormValue("microcontroller_inet_port")
		q["microcontroller_password"] = r.FormValue("microcontroller_password")
		q["microcontroller_firmware"] = r.FormValue("microcontroller_firmware")
		q["microcontroller_active_template"] = r.FormValue("microcontroller_active_template")
		q["microcontroller_type"] = r.FormValue("microcontroller_type")
		q["room_location"] = r.FormValue("room_location")
		q["room_id"] = r.FormValue("room_id")
		a, _ := model.GetArduinoBy_ArduinoID(q["id"].(string))
		m := make(map[string]interface{})

		o, e := model.GetRoomBy_RoomID(a.GetRoomID())
		if e == nil {
			o.RemoveArduino()
		}

		rm, err := model.GetRoomBy_RoomID(r.FormValue("room_id"))
		if err != nil || rm == nil {
			m["success"] = false
			q["room_id"] = ""
			m["room_location"] = ""

		} else {
			m["success"] = true
			m["room_location"] = rm.Location().Name()
		}

		a.StoreUpdate(q)

		b, _ := json.MarshalIndent(m, " ", "  ")
		//response := `{"success" : "` + strconv.FormatBool(success) + `"}`
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

/**
Arduino Fetch Method

@param id string    The service_id for the Arduino
		=>  As a standard definition we settled on using the hardware MAC Address of the device
			that we are working with, that way
*/
func ArduinoFetch(id string, w http.ResponseWriter, r *http.Request) {
	//already in Method Fetch, now serve Data via json to the arduino
	fmt.Println(id)
	a, _ := model.GetArduinoBy_ArduinoID(id)
	t, err := model.GetTemplateByTemplateName(a.GetTemplateName())
	if err != nil {
		m := make(map[string]interface{})
		m["success"] = false
		m["time"] = time.Now().String()
		b, _ := json.MarshalIndent(m, " ", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
	m := templates.CreateAPIData(t)

	var R *model.Room
	R, err = model.GetRoomBy_RoomID(a.GetRoomID())
	fmt.Println(R, err)

	for i, e := range *t.GetElements() {
		if e.GetContentStatic() {
			//TODO Serve a dynamic content for the keyword, given by t.content
			c := " - "

			c = templates.GetContentDynamic(a, R, e.GetContent())

			fmt.Println(query.ELEMENT_CONTENT)
			m[query.TEMPLATE_ELEMENTS].([]map[string]interface{})[i][query.ELEMENT_CONTENT] = c
		}
	}
	m["noe"] = len(m[query.TEMPLATE_ELEMENTS].([]map[string]interface{}))
	m["time"] = time.Now().String()
	b, _ := json.MarshalIndent(m, " ", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
	return
}
