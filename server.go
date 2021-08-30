package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"mus_projekt/app/controller"
	"mus_projekt/app/controller/api"
	"mus_projekt/utils"
	"net/http"
)

func main() {
	MuxRouter := mux.NewRouter()
	fmt.Print("Booting the new RosRockServer")

	MuxRouter.HandleFunc("/", controller.Index)

	MuxRouter.HandleFunc("/login", controller.Login)

	//MuxRouter.HandleFunc("/users", controller.User)

	MuxRouter.HandleFunc("/register", controller.Register)

	MuxRouter.HandleFunc("/logout", controller.Logout)

	MuxRouter.HandleFunc("/rooms", controller.Rooms)
	MuxRouter.HandleFunc("/rooms/{room_id}", controller.RoomsHandler)

	MuxRouter.HandleFunc("/register/confirm", controller.Confirm)

	MuxRouter.HandleFunc("/confirmation", controller.Confirmation)

	MuxRouter.HandleFunc("/pw_reset", controller.PasswordReset)

	MuxRouter.HandleFunc("/users", controller.User)

	//TODO: Weiterleitung auf verschiedene URLS
	//TODO: Datenanbindung über speziellen pfad => /users?id'öjbsegüaosidbf'
	//TODO: CRUD Methoden bereitstellen für die Datenbanken => Create/Read/Update/Delete
	//TODO: Pfad methoden erkennen und zuweisen

	MuxRouter.HandleFunc("/api/{options}", controller.ApiMux)
	MuxRouter.HandleFunc("/api/{options}/{element}", controller.ElementMux)

	MuxRouter.PathPrefix("/css/").Handler(http.StripPrefix("/css", http.FileServer(http.Dir(utils.GetLocalEnv()+"static/css"))))
	MuxRouter.PathPrefix("/script/").Handler(http.StripPrefix("/script", http.FileServer(http.Dir(utils.GetLocalEnv()+"static/scripts"))))
	MuxRouter.PathPrefix("/js/").Handler(http.StripPrefix("/js", http.FileServer(http.Dir(utils.GetLocalEnv()+"static/js"))))
	MuxRouter.PathPrefix("/images/").Handler(http.StripPrefix("/images", http.FileServer(http.Dir(utils.GetLocalEnv()+"static/images"))))
	MuxRouter.PathPrefix("/svg/").Handler(http.StripPrefix("/svg", http.FileServer(http.Dir(utils.GetLocalEnv()+"static/svg"))))
	http.Handle("/", MuxRouter)
	server := http.Server{
		Addr: ":8080",
	}
	fmt.Println("****************\n\n    Server has been booted\n    IP  : " + api.SERVER_URL + "\n" + "    PORT: " + api.SERVER_PORT + "\n\n****************")
	server.ListenAndServe()

}
