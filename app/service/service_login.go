package service

import (
	"github.com/leesper/couchdb-golang"
	"mus_projekt/app/controller/api"
	"mus_projekt/app/controller/protocol"

	"mus_projekt/app/model/utils"
)

const CORFIRM_JSON = "confirm_account"
const PWRS_JSON = "reset_password"
const IMPRESS_JSON = "impressum"

const LGN_SH = "lsh"

var service_db *couchdb.Database

func init() {
	service_db, _ = couchdb.NewDatabase("http://localhost:5984/ros_rocks_service")
}

type login_service_helper struct {
	_id          string `json:"_id"`
	_rev         string `json:"_rev"`
	service_id   string `json:"service_id"`
	user_email   string `json:"user_email"`
	service_type string `json:"service_type"`
}

func sh2m(s *login_service_helper) map[string]interface{} {
	m := make(map[string]interface{})
	m["_id"] = s._id
	m["_rev"] = s._rev
	m["service_id"] = s.service_id
	m["user_email"] = s.user_email
	m["service_type"] = s.service_type
	return m
}

func m2sh(m map[string]interface{}) *login_service_helper {
	s := login_service_helper{
		_id:          m["_id"].(string),
		_rev:         m["_rev"].(string),
		service_id:   m["service_id"].(string),
		user_email:   m["user_email"].(string),
		service_type: m["service_type"].(string),
	}
	return &s
}

func RemoveByID(e string) error {
	s, err := getLSHBYID(e)
	if err != nil {
		return err
	}
	service_db.Delete(s._id)
	return nil
}

func RemoveLSHByEmail(e string) error {
	s, err := getLSHBYE(e)
	if err != nil || s == nil {
		return err
	}
	service_db.Delete(s._id)
	return nil
}

func AuthLSHByID(id string) (bool, string, error) {
	s, err := getLSHBYID(id)
	if err != nil || s == nil {
		return false, "", err
	}
	return true, s.user_email, nil
}

func AuthLSHByEmail(e string) (bool, string, error) {
	s, err := getLSHBYE(e)
	if err != nil || s == nil {
		return false, "", err
	}
	return true, s.user_email, nil
}

func createLSH(id string, e string) *login_service_helper {
	s := login_service_helper{
		_id:          "",
		_rev:         "",
		service_id:   id,
		user_email:   e,
		service_type: LGN_SH,
	}
	return &s
}

func (s *login_service_helper) store() error {
	m := sh2m(s)
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

func GetLSHByEmail(e string) *login_service_helper {
	s, _ := getLSHBYE(e)
	return s
}

func (s *login_service_helper) save() error {
	m := sh2m(s)
	_id, _rev, err := service_db.Save(m, nil)
	if err != nil {
		return err
	}
	s._id = _id
	s._rev = _rev
	return nil
}

func getLSHBYID(i string) (*login_service_helper, error) {
	query := `{"selector" :{"service_id":"` + i + `", "service_type" :"` + LGN_SH + `"}}`
	m, err := service_db.QueryJSON(query)
	if err != nil || len(m) == 0 {
		return nil, err
	}
	return m2sh(m[0]), nil
}

func getLSHBYE(e string) (*login_service_helper, error) {
	query := `{"selector" :{"user_email":"` + e + `", "service_type" :"` + LGN_SH + `"}}`
	m, err := service_db.QueryJSON(query)
	if err != nil || len(m) == 0 {
		return nil, err
	}
	return m2sh(m[0]), nil
}

func CreateConfirmationAccess(e string) (map[string]interface{}, error) {
	id := utils.GenerateIDL(64)
	s := createLSH(id, e)
	err := s.store()
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	m[CORFIRM_JSON] = "http://" + api.SERVER_URL + ":" + api.SERVER_PORT + protocol.Confirm + "?id=" + id
	m[IMPRESS_JSON] = "http://" + api.SERVER_URL + ":" + api.SERVER_PORT + protocol.Impressum
	return m, nil

}

func GetConfirmationAccess(e string) (map[string]interface{}, error) {
	sh, er := getLSHBYE(e)
	if er != nil {
		return nil, er
	}
	m := make(map[string]interface{})
	m[CORFIRM_JSON] = "http://" + api.SERVER_URL + ":" + api.SERVER_PORT + protocol.Confirm + "?id=" + sh.service_id
	m[IMPRESS_JSON] = "http://" + api.SERVER_URL + ":" + api.SERVER_PORT + protocol.Impressum
	return m, nil
}
