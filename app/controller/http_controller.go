package controller

import (
	"fmt"
	"html/template"
	"mus_projekt/app/auth"
	"mus_projekt/app/controller/api"
	"mus_projekt/app/model/query"
	"mus_projekt/utils"
	"net"
	"net/http"
	"strings"
)

/*
	The Controller to manage all publicly incoming
	server requests handled by the controller package

	This public.go file manages all incoming http
	requests for handling html/css/js website content
	that is publicly available
	
	This file also fetches the network address for the network card and local machine such that the machine can be communicated from the lan, without any special settings necessary

	@author ben

*/

var html_templates *template.Template

func init() {
	html_templates = template.Must(template.ParseGlob(utils.GetLocalEnv() + "static/template/*.tmpl"))

	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if !strings.Contains(ip.String(), ":") && ip.String() != "127.0.0.1" {
				fmt.Println(ip)
				api.SERVER_URL = ip.String()
				//api.SERVER_URL = "localhost"
				return
			}
		}
		//api.SERVER_URL = "localhost"
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving the Index Website")

	if r.Method == http.MethodGet {

		session, _ := auth.GetCookieStore().Get(r, auth.GetSessionCookie())
		a := session.Values[auth.USER_COOKIE_AUTH]
		l := session.Values[auth.USER_ACCESS_LEVEL]
		m := make(map[string]interface{})
		m["auth"] = a
		m[api.SERVER_ADDRESS_KEY] = api.SERVER_URL
		m[api.SERVER_PORT_KEY] = api.SERVER_PORT
		m[query.ADMIN_LEVEL] = l
		html_templates.ExecuteTemplate(w, "index.tmpl", m)
		return
	}

}

func ExecuteTemplate(s string, w http.ResponseWriter, r *http.Request, data interface{}) error {
	return html_templates.ExecuteTemplate(w, s, data)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println("http:GET:Logout")
		session, _ := auth.GetCookieStore().Get(r, auth.GetSessionCookie())
		if e := session.Values[auth.USER_COOKIE_AUTH]; e != nil && e.(bool) {
			session.Values[auth.USER_COOKIE_AUTH] = false
			session.Values[auth.USER_ACCESS_LEVEL] = -1
			session.Values[auth.USER_ID] = ""
			session.Save(r, w)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
	}
	http.Redirect(w, r, "/login", http.StatusFound)
	return
}
