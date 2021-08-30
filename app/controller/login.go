package controller

import (
	"encoding/json"
	"fmt"
	"mus_projekt/app/auth"
	"mus_projekt/app/controller/protocol"
	"mus_projekt/app/model"
	"mus_projekt/app/model/query"
	"net/http"
	"strings"
)

/**
	This file manages all serverside communication using the http protocol for the login page

	@author ben
 */
func Serve_Login_Checks(exist bool) map[string]interface{} {
	m := make(map[string]interface{})
	m[query.USER_EXIST] = exist
	return m
}

func Serve_Login(e bool, a bool, err error) map[string]interface{} {
	m := make(map[string]interface{})
	m[query.USER_EXIST] = e
	m[query.USER_AUTHENTICATED] = a && err == nil
	return m
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		fmt.Println("http:Get:Login")
		session, _ := auth.GetCookieStore().Get(r, auth.GetSessionCookie())
		if e := session.Values[auth.USER_COOKIE_AUTH]; e != nil && e.(bool) {
			http.Redirect(w, r, "/", http.StatusFound)
		}

		if strings.Contains(r.UserAgent(), "Mobile") {
			html_templates.ExecuteTemplate(w, "login_mobile.tmpl", nil)
		} else {
			html_templates.ExecuteTemplate(w, "login.tmpl", nil)
		}

	}

	if r.Method == protocol.MethodCheck {
		fmt.Println("http:CHECK:Login")
		exist, _, _ := model.Check_User_By_Email(r.FormValue(query.USER_EMAIL_ADDRESS), "")
		m := Serve_Login_Checks(exist)
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}

	if r.Method == http.MethodPost {

		fmt.Println("http:POST:Login")
		n := r.FormValue(query.USER_EMAIL_ADDRESS)
		p := r.FormValue(query.USER_PW)
		m := Serve_Login(model.Check_User_By_Email(n, p))

		if m[query.USER_AUTHENTICATED].(bool) {
			fmt.Println("Login Successful")
			session, _ := auth.GetCookieStore().Get(r, auth.GetSessionCookie())
			session.Values[auth.USER_COOKIE_AUTH] = true
			u, _ := model.GetUserByEmailAddress(n)
			session.Values[auth.USER_ACCESS_LEVEL] = u.GetAccessLevel()
			session.Values[auth.USER_ID] = u.GetID()
			session.Save(r, w)
		}

		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		fmt.Println("Sending the json login data")
		return

	}
}
