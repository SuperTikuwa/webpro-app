package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/SuperTikuwa/webpro-app/dboperation"
	"github.com/SuperTikuwa/webpro-app/model"
)

func ScheduleHandler(w http.ResponseWriter, r *http.Request) {
}

func createScheduleHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var s model.ScheduleRequest
	if err := json.Unmarshal(b, &s); err != nil {
		panic(err)
	}

	schedule := s.ToSchedule()

	if err := dboperation.CreateSchedule(schedule); err != nil {
		panic(err)
	}
}

func searchScheduleHandler(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("UserName")

	user, err := dboperation.SelectUserByName(name)
	if err != nil {
		panic(err)
	}

	schedules, err := dboperation.SelectAllScheduleByUserID(user.ID)
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(schedules); err != nil {
		panic(err)
	}
}
