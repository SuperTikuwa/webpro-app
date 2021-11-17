package main

import (
	"net/http"

	"github.com/SuperTikuwa/webpro-app/handler"
)

func main() {
	http.HandleFunc("/users", handler.CreateUserHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/schedules", handler.ScheduleHandler)

	http.ListenAndServe(":8088", nil)
}
