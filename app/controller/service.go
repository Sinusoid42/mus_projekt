package controller

import (
	"encoding/json"
	"fmt"
	"mus_projekt/app/model"
	"mus_projekt/app/model/query"
	"mus_projekt/app/service"
	"net/http"
)

/**
	When accessing any service related content, this file contains all functionalities to respond to account confirmation, password reset requests and asynchronous data

	@author ben
 */

func Confirm(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		id := r.FormValue("id")
		a, e, err := service.AuthLSHByID(id)
		if !a || err != nil {
			//auth didnt work, the id does not exist or has already been removed
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		u, _ := model.GetUserByEmailAddress(e)
		u.Save()
		http.Redirect(w, r, "/confirmation?id="+id, http.StatusFound)
		return
	}
}

func Confirmation(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		m := make(map[string]interface{})
		//TODO

		i := r.FormValue("id")
		m["id"] = i
		fmt.Println(i)
		a, e, er := service.AuthLSHByID(i)
		if !a || er != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		if a {
			//The confirmation db works
			fmt.Println("Confirmation Works", e)
			m["user_email_address"] = e
			html_templates.ExecuteTemplate(w, "account_confirmation.tmpl", m)
			return
		}
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if r.Method == http.MethodPost {
		m := make(map[string]interface{})
		//TODO

		em := r.FormValue(query.USER_EMAIL_ADDRESS)
		fmt.Println(em)
		a, e, er := service.AuthLSHByEmail(em)
		if !a || er != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		if a {
			passwordRot13 := r.FormValue(query.USER_PW)
			username := r.FormValue(query.USER_NM)
			u, _ := model.GetUserByEmailAddress(e)

			u.Set(username, passwordRot13)

			m[query.SUCCESS] = true

			u.Confirm(a)
			u.Save()
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}
	}
}

func UpgradeUser(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	a, e, err := service.AuthUSHBYID(id)

	fmt.Println(a, e, err)

}
