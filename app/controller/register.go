package controller

import (
	"encoding/json"
	"fmt"
	"mus_projekt/app/auth"
	"mus_projekt/app/model"
	"mus_projekt/app/model/query"
	"net/http"
)

/*
	This file manages the http communication when trying to register a new account

	@author ben
 */
func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		fmt.Println("http:GET:Register")
		session, _ := auth.GetCookieStore().Get(r, auth.GetSessionCookie())
		if e := session.Values[auth.USER_COOKIE_AUTH]; e != nil && e.(bool) {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

	}

	if r.Method == http.MethodPost {
		m := make(map[string]interface{})
		fmt.Println("http:POST:Register")
		e := r.FormValue(query.USER_EMAIL_ADDRESS)

		ex, _, err := model.Check_User_By_Email(e, "")
		if ex || err != nil {
			m[query.SUCCESS] = false
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}

		u, err := model.CreateUserEP(e, "")
		u.SendWelcomeEmail()

		if err != nil {
			println(err)
		}
		u.StoreUser()

		m[query.SUCCESS] = true
		b, _ := json.MarshalIndent(m, "", "  ")

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
}
