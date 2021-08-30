package utils

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"time"
)

var dictionary []string

func init() {
	dictionary = make([]string, 62)
	fillDictionary()
}

func fillDictionary() {
	var i = 0
	var asciiOffset = 65
	for ; i < 26; i++ {
		dictionary[i] = string(i + asciiOffset)
		dictionary[i+26] = string(i + asciiOffset + 32)
	}
	for i = 0; i < 10; i++ {
		dictionary[i+52] = string(i + 48)
	}
}

/*
	Genrerates a 16 byte long string utf8 representation of a new id
	for either a room or a new stream user

	the id can be found within the database for the room and user stream_ids
	or within the byte stream itself at stream.stream_api.DATA_STREAM_STREAM_MEMBER_ID
*/
func GenerateID() string {
	var i = 0
	var s = ""
	rand.Seed(time.Now().UnixNano())
	for ; i < 16; i++ {
		s += dictionary[rand.Intn(len(dictionary))]
	}
	return s
}

func GenerateIDL(l int) string {
	var i = 0
	var s = ""
	rand.Seed(time.Now().UnixNano())
	for ; i < l; i++ {
		s += dictionary[rand.Intn(len(dictionary))]
	}
	return s
}

/*
	Convert the User Password using base 64 and bcrypt
*/
func EncryptUserPassword(pw string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
	b64HashedPwd := base64.StdEncoding.EncodeToString(hashedPwd)

	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return b64HashedPwd, nil
}

func GetDateFromISOWeek(year int, week int) time.Time {
	date := time.Date(year, 0, 0, 0, 0, 0, 0, time.Local)
	iy, iw := date.ISOWeek()
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
		iy, iw = date.ISOWeek()
	}
	for iy < year {
		date.AddDate(0, 0, 7)
		iy, iw = date.ISOWeek()
	}
	for iw < week {
		date.AddDate(0, 0, 7)
		iy, iw = date.ISOWeek()
	}
	return date
}
