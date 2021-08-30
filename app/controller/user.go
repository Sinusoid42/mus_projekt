package controller

import (
	"encoding/json"
	"fmt"
	"mus_projekt/app/auth"
	"mus_projekt/app/controller/protocol"
	"mus_projekt/app/controller/users"
	"mus_projekt/app/model"
	"mus_projekt/app/model/query"
	"mus_projekt/app/service"
	"net/http"
)

/**
	This file contains all functionalities relevant for a user account management

	Pasword resets, upgrades, useraccount page functionalities and async connections are handled here

	@author ben
 */
func PasswordReset(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println("http:Reset:GET")
		id := r.FormValue("id")
		if len(id) == 0 {
			html_templates.ExecuteTemplate(w, "pw_reset_request.tmpl", nil)
			return
		}
		a, e, err := service.AuthPWSHByID(id)
		if !a || err != nil {
			http.Redirect(w, r, "/pw_reset", http.StatusFound)
			html_templates.ExecuteTemplate(w, "pw_reset_request.tmpl", nil)
			return
		}

		fmt.Println(e)
		html_templates.ExecuteTemplate(w, "pw_reset.tmpl", nil)

	}

	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		fmt.Println(id)
		if len(id) != 0 {
			a, em, err := service.AuthPWSHByID(id)
			if !a || err != nil {
				fmt.Println(id)
				fmt.Println(em)
				m := make(map[string]interface{})
				m[query.SUCCESS] = false
				b, _ := json.MarshalIndent(m, "", "  ")
				w.Header().Set("Content-Type", "application/json")
				w.Write(b)
				return
			}
			fmt.Println(id)
			fmt.Println(em)
			r13pw := r.FormValue(query.USER_PW)
			u, _ := model.GetUserByEmailAddress(em)
			u.HandlePasswordReset(r13pw)
			m := make(map[string]interface{})
			m[query.SUCCESS] = true
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		} else {
			fmt.Println("http:Reset:POST")
			e := r.FormValue(query.USER_EMAIL_ADDRESS)
			u, er := model.GetUserByEmailAddress(e)

			if er != nil {
				fmt.Println("http:Reset:POST:NO_USER_FOUND")
				m := make(map[string]interface{})
				m[query.USER_EXIST] = false
				m["answer0"] = "There was no account registered for this email, forwarding to login. . . "
				b, _ := json.MarshalIndent(m, "", "  ")
				w.Header().Set("Content-Type", "application/json")
				w.Write(b)
				return
			}
			if !u.Confirmed() {
				fmt.Println("http:Reset:POST:NOT_CONFIRMED_ACCOUNT")
				u.ResendWelcomeEmail()
				m := make(map[string]interface{})
				m["answer0"] = "Your email address was not confirmed yet, we will resend the confirmation to you"
				m["answer1"] = "Maybe you should check your Junk-Emails (happens to me sometimes at least while testing this system"
				m[query.USER_EXIST] = true
				b, _ := json.MarshalIndent(m, "", "  ")
				w.Header().Set("Content-Type", "application/json")
				w.Write(b)
				return
			}
			p, t := u.PasswordResetTimeout()
			if p == service.PW_LOCK {
				fmt.Println("http:Reset:POST:PW_RESET_TIMEOUT")
				m := make(map[string]interface{})
				m["answer0"] = "You already changed your password within the last 12h, please wait " + t
				m["answer1"] = ""
				m[query.USER_EXIST] = true
				b, _ := json.MarshalIndent(m, "", "  ")
				w.Header().Set("Content-Type", "application/json")
				w.Write(b)
				return
			}
			u.SendPasswordResetEmail()
			m := make(map[string]interface{})
			if p == service.PW_RENEW {
				m["answer0"] = "There already existed a password reset request for your account, creating a new one now ..."
				m["answer1"] = "Maybe you should check your Junk-Emails"
			}
			if p == service.PW_NEW {
				m["answer0"] = "Check your mail inbox for an email to reset your password"
				m["answer1"] = ""
			}
			m[query.USER_EXIST] = true
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}
	}

	if r.Method == protocol.MethodCheck {
		e := r.FormValue(query.USER_EMAIL_ADDRESS)
		_, er := model.GetUserByEmailAddress(e)
		m := make(map[string]interface{})
		if er == nil {
			m[query.USER_EXIST] = true
		} else {
			m[query.USER_EXIST] = false
		}
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
}

func User(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		var m map[string]interface{}
		m = make(map[string]interface{})
		session, _ := auth.GetCookieStore().Get(r, auth.GetSessionCookie())
		a := session.Values[auth.USER_COOKIE_AUTH]

		id := session.Values[auth.USER_ID]

		j, er := model.Check_User_By_ID(id)
		if !j || er != nil || a == nil || !a.(bool) {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		u, _ := model.GetUserByID(id.(string))
		m = users.CreateApiData(u)
		m[query.SUCCESS] = true
		m[query.AUTH] = a.(bool)
		m[query.ADMIN_LEVEL] = u.GetAccessLevel()

		m["user_id"] = u.GetID()
		m["access_name"] = u.AccessName()

		m["upgradeable"] = u.Upgradeable()
		api_rooms, api_arduinos, api_templates := users.CheckAccess(u)
		m[query.HTML_API_ROOMS_ACCESS] = api_rooms
		m[query.HTML_API_ARDUINOS_ACCESS] = api_arduinos
		m[query.HTML_API_TEMPLATES_ACCESS] = api_templates

		m["reoccurring_available"] = u.GetAccessLevel() == -1 || u.GetAccessLevel() >= 2

		html_templates.ExecuteTemplate(w, "account.tmpl", m)
		return
	}

	if r.Method == http.MethodPost {

		//TODO check for upgradeability for the user that is logged in
		//if true => create upgrade helper list
		fmt.Println("POST FROM THE USER FOR AN UPGRADE")

		session, _ := auth.GetCookieStore().Get(r, auth.GetSessionCookie())
		a := session.Values[auth.USER_COOKIE_AUTH]
		id := session.Values[auth.USER_ID]

		j, er := model.Check_User_By_ID(id)
		if !j || er != nil || a == nil || !a.(bool) {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		u, _ := model.GetUserByID(id.(string))
		q := u.Upgradeable()
		if !q {
			m := make(map[string]interface{})
			m[query.SUCCESS] = false
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}
		users.CreateUpgradeTicket(u)
		m := make(map[string]interface{})
		m[query.SUCCESS] = true
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}

	if r.Method == protocol.MethodFetch {
		session, _ := auth.GetCookieStore().Get(r, auth.GetSessionCookie())
		a := session.Values[auth.USER_COOKIE_AUTH]
		id := session.Values[auth.USER_ID]
		j, er := model.Check_User_By_ID(id)
		if a == nil || !a.(bool) || !j || er != nil {
			fmt.Println("http:User:POST:NO_USER_FOUND")
			m := make(map[string]interface{})
			m[query.SUCCESS] = false
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}
		m := users.CreateUserPageBookings(id.(string))
		m[query.SUCCESS] = true
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}

	if r.Method == protocol.MethodSelect {
		session, _ := auth.GetCookieStore().Get(r, auth.GetSessionCookie())
		a := session.Values[auth.USER_COOKIE_AUTH]
		id := session.Values[auth.USER_ID]
		j, er := model.Check_User_By_ID(id)
		if a == nil || !a.(bool) || !j || er != nil {
			fmt.Println("http:User:POST:NO_USER_FOUND")
			m := make(map[string]interface{})
			m[query.SUCCESS] = false
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}
		bkid := r.FormValue(query.BOOKING_ID)
		bk, _ := model.GetBookingByID(bkid)

		m := users.CreateSelectionData(bk)
		m[query.SUCCESS] = true
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
	if r.Method == http.MethodPut {
		session, _ := auth.GetCookieStore().Get(r, auth.GetSessionCookie())
		a := session.Values[auth.USER_COOKIE_AUTH]
		id := session.Values[auth.USER_ID]
		j, er := model.Check_User_By_ID(id)
		if a == nil || !a.(bool) || !j || er != nil {
			fmt.Println("http:User:POST:NO_USER_FOUND")
			m := make(map[string]interface{})
			m[query.SUCCESS] = false
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}
		bid := r.FormValue(query.BOOKING_ID)
		bk, err := model.GetBookingByID(bid)
		if err != nil {
			m := make(map[string]interface{})
			m[query.SUCCESS] = false
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}

		m, _ := bk.SaveBooking(r)
		m[query.SUCCESS] = true
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
}
