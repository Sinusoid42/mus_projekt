package controller

import (
	"encoding/json"
	"fmt"
	"mus_projekt/app/controller/protocol"
	"mus_projekt/app/controller/upgrades"
	"mus_projekt/app/model"
	"mus_projekt/app/model/query"
	"mus_projekt/app/service"
	"net/http"
	"strconv"
)

/**
	When upgrading an account of any ticketed user, this file contains all functionalities for
	a webbrowser to access the api/upgrades page

	relevant persmissions are checked in the api/mux section

	@author ben
 */

func UpgradeApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Upgrade API")

	if r.Method == http.MethodGet {
		html_templates.ExecuteTemplate(w, "upgrade_template.tmpl", nil)
	}

	if r.Method == protocol.MethodFetch {

		m := upgrades.CreateApiData()
		fmt.Println(m)
		m[query.SUCCESS] = true
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}

	if r.Method == http.MethodPost {

		t := r.FormValue("t")

		if t == "t0" {

			//TODO send email about change and create ticket

			level := r.FormValue("l")
			id := r.FormValue("id")

			access, e := strconv.Atoi(level)
			if e != nil {
				return
			}
			ush, err := service.GetUSHBYID(id)
			if err != nil {
				fmt.Println(err)
				return
			}
			u, err := model.GetUserByID(ush.GetUserID())
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Y>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
			fmt.Println("Hier sende ich jetzt was")
			fmt.Println(access)
			ush.SaveAccessLevel(access)

			q, err := service.GetUSHBYID(id)

			fmt.Println("\n\n", q, "\n\n")

			u.UpgradeAccount(access)
			fmt.Println(ush)
			fmt.Println(u)
			u.SendAccountUpgradeEmail(access)

			m := make(map[string]interface{})
			m[query.SUCCESS] = true
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
			return
		}

		if t == "t1" {

			//level := 0
			//id := r.FormValue("id")

		}

		m := make(map[string]interface{})
		m[query.SUCCESS] = false
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return

	}

}
