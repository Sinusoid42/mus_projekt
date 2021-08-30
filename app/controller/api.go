package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"mus_projekt/app/auth"
	"mus_projekt/app/controller/api"
	"mus_projekt/app/controller/users"
	"mus_projekt/app/model"
	"mus_projekt/app/model/query"
	"net/http"
)
/*
	The api multiplexer for all HTTP requests, when accessing the api functionalities of the web application
 */
func ApiMux(w http.ResponseWriter, r *http.Request) {

	v := mux.Vars(r)
	c := v["options"]

	session, _ := auth.GetCookieStore().Get(r, auth.GetSessionCookie())
	a := session.Values[auth.USER_COOKIE_AUTH]
	l := session.Values[auth.USER_ACCESS_LEVEL]

	id := session.Values[auth.USER_ID]

	switch c {
	case "server":
		{

			fmt.Println("http:API:Server")
			m := getServerDataJson()
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
		}
	case "arduinos":
		{
			if a == nil || !a.(bool) || l == nil {
				http.Redirect(w, r, "/", http.StatusUnauthorized)
			}

			u, err := model.GetUserByID(id.(string))
			_, j, _ := users.CheckAccess(u)
			if j && err == nil {
				Arduino(w, r)
			}
		}
	case "rooms":
		{
			if a == nil || !a.(bool) || l == nil {
				http.Redirect(w, r, "/", http.StatusUnauthorized)
			}

			u, err := model.GetUserByID(id.(string))
			i, _, _ := users.CheckAccess(u)
			if i && err == nil {
				RoomsApi(w, r)
			}
		}
	case "menu":
		{
			MenuApi(w, r)
		}
	case "templates":
		{

			if a == nil || !a.(bool) || l == nil {
				http.Redirect(w, r, "/", http.StatusUnauthorized)
			}

			u, err := model.GetUserByID(id.(string))
			_, _, k := users.CheckAccess(u)
			if k && err == nil {
				TemplateApi(w, r)
			}
		}
	case "upgrades":
		{
			if a == nil || !a.(bool) || l == nil {
				http.Redirect(w, r, "/", http.StatusUnauthorized)
			}
			u, err := model.GetUserByID(id.(string))
			_, _, k := users.CheckAccess(u)
			if k && err == nil {
				UpgradeApi(w, r)
			}
		}
	default:
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

func getServerDataJson() map[string]interface{} {

	m := make(map[string]interface{})
	m[api.SERVER_ADDRESS_KEY] = api.SERVER_URL
	m[api.SERVER_PORT_KEY] = api.SERVER_PORT

	return m

}

func ElementMux(w http.ResponseWriter, r *http.Request) {

	v := mux.Vars(r)
	m := make(map[string]interface{})
	option := v["options"]
	element := v["element"]
	fmt.Println(element)
	if r.Method == http.MethodGet {
		switch option {
		case "arduinos":
			{
				/*if len(element) != 16 { //check fir the element is type service_id and refers to the arduino
					m[query.SUCCESS] = false
					b, _ := json.MarshalIndent(m, "", "  ")
					w.Header().Set("Content-Type", "application/json")
					w.Write(b)
					return
				}*/
				ArduinoFetch(element, w, r)
				return
			}
		}
	}
	m[query.SUCCESS] = false
	b, _ := json.MarshalIndent(m, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
	return
}
