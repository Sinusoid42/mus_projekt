package model

import (
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	authentaction "mus_projekt/app/auth"
	"mus_projekt/app/model/query"
	"mus_projekt/app/model/utils"
	"mus_projekt/app/service"
	"time"
)

/*

	Manages the User Info

	Datamanagement with the Database and Admin Access auth data

*/

type User struct {
	_id  string `json:"_id"`  //The UUID of the User within the database
	_rev string `json:"_rev"` //The REV service_id of the user within the database

	user_id string `json:"user_id"`

	confirmed bool `json:"confirmed"`

	user_name string `json:"user_name"`

	user_password      string   `json:"user_password"`
	admin_access_level int      `json:"admin_access_level"`
	user_bookings      []string `json:"user_bookings"`

	user_email_address string `json:"user_email_address"` //probably deprecated sooner or later

}

/*
	Proprietary user data, to check
	if the userdata, that is checked exists
	already in the database
	Checks wether a given Username and Password matches
	some within the database
	returns
*/
func Check_User(u string, p string) (bool, bool, error) {
	exist := false
	auth := false
	query := `{"selector":{"user_name":"%s"}}`
	um, err := user_DB.QueryJSON(fmt.Sprintf(query, u))
	if err != nil || len(um) == 0 {
		fmt.Println("The user does probbaly not exist")
		return exist, auth, err
	}
	if u != um[0]["user_name"].(string) {
		return exist, auth, err
	}
	exist = true
	if len(p) == 0 || !um[0]["confirmed"].(bool) {
		return exist, auth, nil
	}
	us := m2u(um[0])
	passwordDB, _ := base64.StdEncoding.DecodeString(us.user_password)
	err = bcrypt.CompareHashAndPassword(passwordDB, []byte(authentaction.DecodeRot13(p)))
	if err != nil {
		return exist, auth, nil
	}
	auth = true
	return exist, auth, nil
}

func (u *User) GetID() string {
	return u._id
}

func Check_User_By_Email(e string, p string) (bool, bool, error) {
	exist := false
	auth := false
	query := `{"selector":{"user_email_address":"%s"}}`
	um, err := user_DB.QueryJSON(fmt.Sprintf(query, e))

	fmt.Println(um)

	if err != nil || len(um) == 0 {
		fmt.Println("The user does probbaly not exist")
		return exist, auth, err
	}
	if e != um[0]["user_email_address"].(string) {
		return exist, auth, err
	}
	exist = true
	if len(p) == 0 {
		return exist, auth, nil
	}
	us := m2u(um[0])
	passwordDB, _ := base64.StdEncoding.DecodeString(us.user_password)
	err = bcrypt.CompareHashAndPassword(passwordDB, []byte(authentaction.DecodeRot13(p)))
	if err != nil {
		return exist, auth, nil
	}
	auth = true
	return exist, auth, nil
}

func GetUserByEmailAddress(e string) (*User, error) {
	if len(e) == 0 {
		return nil, errors.New("User does not exist")
	}
	query := `{"selector":{"user_email_address":"%s"}}`
	um, err := user_DB.QueryJSON(fmt.Sprintf(query, e))
	if err != nil || len(um) == 0 {
		return nil, errors.New("User not found")
	}
	u := m2u(um[0])
	return u, nil
}

func GetUserByUserName(un string) (*User, error) {
	query := `{"selector":{"user_name":"%s"}}`
	um, err := user_DB.QueryJSON(fmt.Sprintf(query, un))
	if err != nil || len(um) == 0 {
		return nil, err
	}
	u := m2u(um[0])
	return u, nil
}

func (u *User) GetAccessLevel() int {
	return u.admin_access_level
}

func CreateUser(name string, pw string) (*User, error) {
	p, _ := utils.EncryptUserPassword(pw)
	u := User{
		"",
		"",
		"",
		false,
		name,
		p,
		int(0),
		[]string{},
		"",
	}
	return &u, nil
}

func CreateUserEP(email string, pw string) (*User, error) {
	p, _ := utils.EncryptUserPassword(pw)
	u := User{
		"_",
		"_",
		"_",
		false,
		"_",
		p,
		int(0),
		[]string{},
		email,
	}
	return &u, nil
}

func (u *User) GetEmailAddress() string {
	return u.user_email_address
}

func (u *User) StoreUser() (bool, error) {
	m, _ := u.u2m()
	delete(m, "_id")
	delete(m, "_rev")
	id, rev, err := user_DB.Save(m, nil)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	u._id = id
	u._rev = rev
	u.user_id = id
	return true, nil
}

func (u *User) Confirmed() bool {
	return u.confirmed
}

func (u *User) Save() (bool, error) {
	m, _ := u.u2m()
	id, rev, err := user_DB.Save(m, nil)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	u._id = id
	u._rev = rev
	return true, nil
}

func (u *User) u2m() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m[query.USER_NM] = u.user_name
	m[query.USER_DB_ID] = u.user_id
	m[query.USER_PW] = u.user_password
	m[query.USER_PRIVILEGES] = u.admin_access_level
	m[query.USER_EMAIL_ADDRESS] = u.user_email_address
	m[query.USER_BOOKINGS] = u.user_bookings
	m[query.USER_DB_ID] = u._id
	m[query.USER_DB_REV] = u._rev
	m[query.USER_CONFIRMED] = u.confirmed
	return m, nil
}

func m2u(m map[string]interface{}) *User {
	b := make([]string, len(m[query.USER_BOOKINGS].([]interface{})))
	for i, v := range m[query.USER_BOOKINGS].([]interface{}) {
		b[i] = v.(string)
	}
	u := User{
		user_name:          m[query.USER_NM].(string),
		user_password:      m[query.USER_PW].(string),
		user_id:            m[query.USER_DB_ID].(string),
		admin_access_level: int(m[query.USER_PRIVILEGES].(float64)),
		user_bookings:      b,
		_id:                m[query.USER_DB_ID].(string),
		_rev:               m[query.USER_DB_REV].(string),
		user_email_address: m[query.USER_EMAIL_ADDRESS].(string),
		confirmed:          m[query.USER_CONFIRMED].(bool),
	}
	return &u
}

func (u *User) SendWelcomeEmail() {

	d, _ := service.CreateConfirmationAccess(u.user_email_address)
	service.Send(u.user_email_address, "Welcome to ROS - Verify your account", d, "email_welcome.tmpl")

}

func (u *User) SendPasswordResetEmail() {
	d, _ := service.CreatePWResetData(u.user_email_address)
	fmt.Println(d)
	fmt.Println("Sending the Password reset email")
	fmt.Println(u.user_email_address)
	service.Send(u.user_email_address, "Password Reset for your ROS Account", d, "email_pw.tmpl")
}

func (u *User) ResendWelcomeEmail() {
	d, e := service.GetConfirmationAccess(u.user_email_address)
	fmt.Println(d, e)
	if e != nil {
		return
	}
	service.Send(u.user_email_address, "Welcome to ROS - Verify your account", d, "email_resend_welcome.tmpl")
}

func (u *User) UpgradeAccount(level int) {
	u.admin_access_level = level
	u.Save()
}

func (u *User) SendAccountUpgradeEmail(access int) {
	m, err := service.CreateUpgradeAccess(u.GetID(), u.GetEmailAddress(), access)
	if err != nil {
		fmt.Println("could not send the upgrade email")
	}
	fmt.Println("SENDENSENDENSENDENSENDEN")
	fmt.Println(m, err)
	service.Send(u.user_email_address, "Your Rooom Occupancy System Account Upgrade", m, "email_upgrade.tmpl")
}

func (u *User) Confirm(a bool) error {
	service.RemoveLSHByEmail(u.user_email_address)
	u.confirmed = a
	return nil
}

func (u *User) PasswordResetTimeout() (int, string) {
	pwsh, err := service.GetPWSHByEmail(u.user_email_address)

	if err != nil {
		return service.PW_NEW, "Error"
	}
	fmt.Println("\n\n\nChecking the Time diff for new PW reset")
	if pwsh.GetTime().Add(time.Duration(time.Hour * 12)).Before(time.Now()) {
		pwsh.RemovePWSH()
		fmt.Println("The time already passed")
		return service.PW_RENEW, ""
	}

	t := pwsh.GetTime().Add(time.Duration(time.Hour * 12)).Sub(time.Now())
	hrs := fmt.Sprint(t)
	return service.PW_LOCK, hrs
}

func (u *User) Set(n string, pw string) error {
	u.user_name = n
	pw = authentaction.DecodeRot13(pw)
	p, _ := utils.EncryptUserPassword(pw)
	u.user_password = p
	return nil
}

func (u *User) HandlePasswordReset(r13pw string) error {
	//pwsh, err := service.GetPWSHByEmail(u.user_email_address)
	//pwsh.RemovePWSH()
	u.setPassword(r13pw)
	u.Save()
	err := service.RemovePWSHByEmail(u.user_email_address)
	return err
}

func (u *User) setPassword(r13pw string) error {
	r13pw = authentaction.DecodeRot13(r13pw)
	p, err := utils.EncryptUserPassword(r13pw)
	u.user_password = p
	return err
}

func GetUserByID(id string) (*User, error) {
	if len(id) == 0 {
		return nil, errors.New("No id")
	}
	m, err0 := user_DB.Get(id, nil)
	if err0 != nil || len(m) == 0 {
		return nil, err0
	}
	u := m2u(m)
	return u, err0
}

func (u *User) GetBookings() []string {
	return u.user_bookings
}

func Check_User_By_ID(id interface{}) (bool, error) {

	if id == nil || len(id.(string)) == 0 {
		return false, errors.New("No Id")
	}
	u, err := GetUserByID(id.(string))
	if err != nil {
		return false, err
	}
	return u.confirmed, nil
}

func (u *User) AttachNewBooking(b *Booking) {
	u.user_bookings = append(u.user_bookings, b.booking_id)
	u.Save()
}

func (u *User) RemoveBookingByID(id string) error {
	for i, j := range u.user_bookings {
		if j == id {
			u.user_bookings[i] = u.user_bookings[len(u.user_bookings)-1]
			return nil
		}
	}
	u.Save()
	return errors.New("The id was not found")
}

func (u *User) GetUserName() string {
	return u.user_name
}

/**
Checks wether there already exists a upgrade request for the user account
*/
func (u *User) Upgradeable() bool {
	if u.admin_access_level == -1 || u.admin_access_level == 5 {
		return false
	}
	a, _, err := service.AuthUSHBYUSERID(u.GetID())
	if a && err == nil {
		return false
	}
	return true
}

func (u *User) AccessName() string {
	switch u.admin_access_level {
	case -1:
		{
			return "server_admin"
		}
	case 0:
		{
			return "public"
		}
	case 1:
		{
			return "student"
		}
	case 2:
		{
			return "tutor"
		}
	case 3:
		{
			return "professor"
		}
	case 4:
		{
			return "admin_professor"
		}
	case 5:
		{
			return "admin"
		}
	}
	return ""
}
