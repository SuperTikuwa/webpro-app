package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/SuperTikuwa/webpro-app/dboperation"
	"github.com/SuperTikuwa/webpro-app/model"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var u model.UserRequest
	err = json.Unmarshal(b, &u)
	if err != nil {
		panic(err)
	}

	u.HashPassword()

	user := u.ToUser()
	if user == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = dboperation.CreateUser(user)
	if err != nil {
		panic(err)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var u model.UserRequest
	err = json.Unmarshal(b, &u)
	if err != nil {
		panic(err)
	}

	user := u.ToUser()

	if user == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = dboperation.Login(user)
	if err != nil {
		panic(err)
	}

	res := model.UserResponse{
		ID:   user.ID,
		Name: user.Name,
	}

	jsonBytes, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, string(jsonBytes))

}
