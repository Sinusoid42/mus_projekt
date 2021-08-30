package service

import (
	"bytes"
	"fmt"
	"html/template"
	"mus_projekt/utils"
	"net/smtp"
)

const service_email_address_google = "room.occupancy.system.no.reply@gmail.com"
const service_email_password = "ros_rocks2020$"
const email_provider_smtp_address = "smtp.gmail.com"

func Send(email_address string, title string, data map[string]interface{}, f string) {
	to := []string{
		email_address,
	}
	auth := smtp.PlainAuth("", service_email_address_google, service_email_password, email_provider_smtp_address)
	t, _ := template.ParseFiles(utils.GetLocalEnv() + "static/email/" + f)

	var body bytes.Buffer
	headers := "MIME_version: 1.0;\nContent-Type: text/html"
	body.Write([]byte(fmt.Sprintf("Subject: "+title+"\n%s\n\n", headers)))
	fmt.Println(utils.GetLocalEnv() + "static/email/" + f)
	fmt.Println("Testausgabe")
	fmt.Println(t)
	fmt.Println(data)
	t.Execute(&body, data)
	err := smtp.SendMail("smtp.gmail.com:587", auth, service_email_address_google, to, body.Bytes())
	if err != nil {
		fmt.Println("Error communicating with the user_email server", err)
		return
	}
	fmt.Println("Send successfully")
}
