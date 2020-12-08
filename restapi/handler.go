package restapi

import (
	"encoding/json"
	"net/http"
)

var db map[string]*UserData

func init() {
	db = make(map[string]*UserData)
}

func RegisterHandler(rw http.ResponseWriter, r *http.Request) {
	user := &UserData{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}
	status := Register(user)
	if status {
		rw.WriteHeader(http.StatusOK)
	}
	response := &Response{
		Message: "Failed to Register",
	}
	byteresp, _ := json.Marshal(response)
	rw.Write(byteresp)
}

func LoginHandler(rw http.ResponseWriter, r *http.Request) {
	data := &Login{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}
	status := CheckLogin(data)
	if status {
		rw.WriteHeader(http.StatusOK)
	}
	rw.WriteHeader(http.StatusUnauthorized)
}

func Register(user *UserData) bool {
	if _, ok := db[user.Username]; ok {
		return false
	}
	db[user.Username] = user
	return true
}

func CheckLogin(user *Login) bool {
	if val, ok := db[user.Username]; ok {
		if val.Password == user.Password {
			//session[val.Username]
			return true
		}
	}
	return false
}
