package service

import (
	"errors"
	"fmt"
	"mus_projekt/app/controller/api"
	"mus_projekt/app/controller/protocol"
	"mus_projekt/app/model/utils"
	"time"
)

const PW_SH = "psh"
const PW_LOCK = 0
const PW_NEW = 1
const PW_RENEW = 2

type password_service_helper struct {
	_id          string `json:"_id"`
	_rev         string `json:"_rev"`
	service_id   string `json:"service_id"`
	user_email   string `json:"user_email"`
	service_type string `json:"service_type"`
	timer        string `json:"timer"`
}

func create_PW_SH(id string, e string) *password_service_helper {
	s := password_service_helper{
		_id:          "",
		_rev:         "",
		service_id:   id,
		user_email:   e,
		service_type: PW_SH,
		timer:        time.Now().Format(time.RFC3339),
	}
	return &s
}

func getPWSHBYID(i string) (*password_service_helper, error) {
	query := `{"selector" :{"service_id":"` + i + `", "service_type" :"` + PW_SH + `"}}`
	m, err := service_db.QueryJSON(query)
	if err != nil || len(m) == 0 {
		return nil, err
	}
	s, _ := m2pwsh(m[0])
	return s, nil
}

func getPWSHBYE(e string) (*password_service_helper, error) {
	query := `{"selector" :{"user_email":"` + e + `", "service_type" :"` + PW_SH + `"}}`
	m, err := service_db.QueryJSON(query)
	if err != nil {
		return nil, err
	}
	if len(m) == 0 {
		return nil, errors.New("The PWSH could not be found")
	}
	s, _ := m2pwsh(m[0])
	return s, nil
}

func CreatePWResetData(e string) (map[string]interface{}, error) {
	id := utils.GenerateIDL(64)
	s := create_PW_SH(id, e)
	s.store()
	fmt.Println("Created the PW RESET")
	m := make(map[string]interface{})
	m[IMPRESS_JSON] = "http://" + api.SERVER_URL + ":" + api.SERVER_PORT + protocol.Impressum
	m[PWRS_JSON] = "http://" + api.SERVER_URL + ":" + api.SERVER_PORT + protocol.PasswordReset + "?id=" + id
	//TODO
	//return proper user_email data with id
	return m, nil
}

func AuthPWSHByID(id string) (bool, string, error) {
	sh, err := getPWSHBYID(id)
	if sh == nil || err != nil {
		return false, "", err
	}
	return true, sh.user_email, nil
}

func GetPWSHByEmail(e string) (*password_service_helper, error) {
	s, err := getPWSHBYE(e)
	return s, err
}

func m2pwsh(m map[string]interface{}) (*password_service_helper, error) {
	s := password_service_helper{
		_id:          m["_id"].(string),
		_rev:         m["_rev"].(string),
		service_id:   m["service_id"].(string),
		user_email:   m["user_email"].(string),
		service_type: m["service_type"].(string),
		timer:        m["timer"].(string),
	}
	return &s, nil
}

func (s *password_service_helper) GetTime() time.Time {
	t, _ := time.Parse(time.RFC3339, s.timer)
	return t
}

func pwsh2m(s *password_service_helper) map[string]interface{} {
	m := make(map[string]interface{})
	m["_id"] = s._id
	m["_rev"] = s._rev
	m["service_id"] = s.service_id
	m["user_email"] = s.user_email
	m["service_type"] = s.service_type
	m["timer"] = s.timer
	return m
}

func (s *password_service_helper) store() error {
	m := pwsh2m(s)
	delete(m, "_id")
	delete(m, "_rev")
	_id, _rev, err := service_db.Save(m, nil)
	if err != nil {
		return err
	}
	s._id = _id
	s._rev = _rev
	return nil
}

func (s *password_service_helper) save() error {
	m := pwsh2m(s)
	_id, _rev, err := service_db.Save(m, nil)
	if err != nil {
		return err
	}
	s._id = _id
	s._rev = _rev
	return nil
}

func RemovePWSHByEmail(e string) error {
	s, err := getPWSHBYE(e)
	fmt.Println(s)
	if err != nil || s == nil {
		return err
	}
	fmt.Println("Removed the PW Ticket for:", e)
	err = service_db.Delete(s._id)
	return err
}

func (s *password_service_helper) RemovePWSH() error {
	err := service_db.Delete(s._id)
	return err
}
