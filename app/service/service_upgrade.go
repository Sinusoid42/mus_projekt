package service

import (
	"errors"
	"fmt"
	"mus_projekt/app/controller/api"
	"mus_projekt/app/controller/protocol"
	"mus_projekt/app/model/utils"
)

const UPGRADE_JSON = "upgrade"

const U_SH = "ush"
const UA_SH = "uash"

type upgrade_service_helper struct {
	_id           string `json:"_id"`
	_rev          string `json:"_rev"`
	service_id    string `json:"service_id"`
	user_id       string `json:"user_id"`
	email_address string `json:"user_email_address"`
	user_access   int    `json:"user_access"`
	service_type  string `json:"service_type"`
}

type upgrade_access_service_helper struct {
	_id           string `json:"_id"`
	_rev          string `json:"_rev"`
	ID            string `json:"service_id"`
	user_id       string `json:"user_id"`
	email_address string `json:"user_email_address"`
	user_access   int    `json:"user_access"`
	t             string `json:"service_type"`
}

func ush2m(s *upgrade_service_helper) map[string]interface{} {
	m := make(map[string]interface{})
	m["_id"] = s._id
	m["_rev"] = s._rev
	m["service_id"] = s.service_id
	m["user_id"] = s.user_id
	m["user_email_address"] = s.email_address
	m["user_access"] = s.user_access
	m["service_type"] = s.service_type
	return m
}

func uash2m(s *upgrade_access_service_helper) map[string]interface{} {
	m := make(map[string]interface{})
	m["_id"] = s._id
	m["_rev"] = s._rev
	m["service_id"] = s.ID
	m["user_id"] = s.user_id
	m["user_email_address"] = s.email_address
	m["user_access"] = s.user_access
	m["service_type"] = s.t
	return m
}

func m2ush(m map[string]interface{}) *upgrade_service_helper {
	s := upgrade_service_helper{
		_id:           m["_id"].(string),
		_rev:          m["_rev"].(string),
		service_id:    m["service_id"].(string),
		user_id:       m["user_id"].(string),
		email_address: m["user_email_address"].(string),
		user_access:   int(m["user_access"].(float64)),
		service_type:  m["service_type"].(string),
	}
	return &s
}

func m2uash(m map[string]interface{}) *upgrade_access_service_helper {
	s := upgrade_access_service_helper{
		_id:           m["_id"].(string),
		_rev:          m["_rev"].(string),
		ID:            m["service_id"].(string),
		user_id:       m["user_id"].(string),
		email_address: m["user_email_address"].(string),
		user_access:   int(m["user_access"].(float64)),
		t:             m["service_type"].(string),
	}
	return &s
}

func (s *upgrade_service_helper) GetServiceID() string {
	return s.service_id
}

func getUSHBYID(id string) (*upgrade_service_helper, error) {
	query := `{"selector" :{"service_id":"` + id + `", "service_type" :"` + U_SH + `"}}`
	m, err := service_db.QueryJSON(query)
	if err != nil || len(m) == 0 {
		return nil, errors.New("No Service Element Found")
	}
	return m2ush(m[0]), nil
}

func getUSHBYUSERID(id string) (*upgrade_service_helper, error) {
	query := `{"selector" :{"user_id":"` + id + `", "service_type" :"` + U_SH + `"}}`
	m, err := service_db.QueryJSON(query)
	if err != nil || len(m) == 0 {
		return nil, errors.New("No Service Element Found")
	}
	return m2ush(m[0]), nil
}

func (s *upgrade_service_helper) store() error {
	m := ush2m(s)
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

func (s *upgrade_access_service_helper) store() error {
	m := uash2m(s)
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

func (s *upgrade_service_helper) save() error {
	m := ush2m(s)
	fmt.Println("TICKETTICKETTCKET >>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println(m)

	_id, _rev, err := service_db.Save(m, nil)
	if err != nil {
		return err
	}
	s._id = _id
	s._rev = _rev
	return nil
}

func CreateUpgradeAccess(user_id string, user_emailadress string, accesslevel int) (map[string]interface{}, error) {
	//i := utils.GenerateIDL(64)
	//s := createUASH(user_id, i, user_emailadress, accesslevel)
	//err := s.store()
	//if err != nil {
	//	return nil, err
	//}

	m := make(map[string]interface{})
	m["login"] = "http://" + api.SERVER_URL + ":" + api.SERVER_PORT + "/login"

	switch accesslevel {
	case 0:
		{
			m["access_name"] = "Public"
		}
	case 1:
		{

			m["access_name"] = "Th Cologne Student"
		}

	case 2:
		{
			m["access_name"] = "Tutor"

		}
	case 3:
		{
			m["access_name"] = "Prof."
		}

	case 4:
		{
			m["access_name"] = "Prof. Admin"

		}
	case 5:
		{
			m["access_name"] = "Room Admin"
		}
	}
	m[IMPRESS_JSON] = "http://" + api.SERVER_URL + ":" + api.SERVER_PORT + protocol.Impressum
	return m, nil
}

func getUASHBYID(id string) (*upgrade_access_service_helper, error) {
	query := `{"selector" :{"service_id":"` + id + `", "service_type" :"` + UA_SH + `"}}`
	m, err := service_db.QueryJSON(query)
	if err != nil || len(m) == 0 {
		return nil, errors.New("No Service Element Found")
	}
	return m2uash(m[0]), nil
}

func CreateUpgradeAccessTicket(id string, e string, a int) (map[string]interface{}, error) {
	i := utils.GenerateIDL(64)
	s := createUSH(id, i, e, a)
	err := s.store()
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	m[UPGRADE_JSON] = "http://" + api.SERVER_URL + ":" + api.SERVER_PORT + protocol.Upgrade + "?id=" + s.service_id
	m[IMPRESS_JSON] = "http://" + api.SERVER_URL + ":" + api.SERVER_PORT + protocol.Impressum
	return m, nil
}

func GetUpgradeAccess(id string) (map[string]interface{}, error) {
	s, err := getUSHBYID(id)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	m[UPGRADE_JSON] = "http://" + api.SERVER_URL + ":" + api.SERVER_PORT + protocol.Upgrade + "?id=" + s.service_id
	m[IMPRESS_JSON] = "http://" + api.SERVER_URL + ":" + api.SERVER_PORT + protocol.Impressum
	return m, nil
}

func createUASH(id string, i string, e string, a int) *upgrade_access_service_helper {
	s := upgrade_access_service_helper{
		_id:           "",
		_rev:          "",
		email_address: e,
		user_id:       id,
		user_access:   a,
		ID:            i,
		t:             UA_SH,
	}
	return &s
}

func createUSH(id string, i string, e string, a int) *upgrade_service_helper {
	s := upgrade_service_helper{
		_id:           "",
		_rev:          "",
		email_address: e,
		user_id:       id,
		user_access:   a,
		service_id:    i,
		service_type:  U_SH,
	}
	return &s
}

func RemoveUSHByID(id string) error {
	s, err := getUSHBYID(id)
	if err != nil || s == nil {
		return err
	}
	service_db.Delete(s._id)
	return nil
}

func RemoveUSHBYUSERID(id string) error {
	s, err := getUSHBYUSERID(id)
	if err != nil || s == nil {
		return err
	}
	service_db.Delete(s._id)
	return nil
}

func AuthUSHBYID(id string) (bool, string, error) {
	s, err := getUSHBYID(id)
	if err != nil || s == nil {
		return false, "", err
	}
	return true, s.user_id, nil
}

func AuthUSHBYUSERID(id string) (bool, string, error) {
	s, err := getUSHBYUSERID(id)
	if err != nil || s == nil {
		return false, "", err
	}
	return true, s.user_id, nil
}

func GetAllUSH() (*[]*upgrade_service_helper, error) {
	query := `{"selector" : {"service_type":"ush"}}`
	m, err := service_db.QueryJSON(query)
	if err != nil {
		return nil, err
	}
	s := []*upgrade_service_helper{}
	for _, i := range m {
		s = append(s, m2ush(i))
		fmt.Println(i)
	}
	return &s, nil
}

func (s *upgrade_service_helper) EmailAddress() string {
	return s.email_address
}

func (s *upgrade_service_helper) UserAccess() int {
	return s.user_access
}

func GetUSHBYID(id string) (*upgrade_service_helper, error) {
	return getUSHBYID(id)
}

func (s *upgrade_service_helper) GetUserID() string {
	return s.user_id
}

func (s *upgrade_service_helper) SaveAccessLevel(a int) {
	s.user_access = a
	s.save()
}
