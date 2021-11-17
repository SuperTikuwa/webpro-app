package model

import (
	"crypto/sha256"
	"encoding/hex"
)

type UserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *UserRequest) HashPassword() {
	b := sha256.Sum256([]byte(u.Password))
	u.Password = hex.EncodeToString(b[:])
}

func (u *UserRequest) validate() bool {
	return u.Name != "" && u.Password != ""
}

func (u *UserRequest) ToUser() *User {
	if !u.validate() {
		return nil
	}

	return &User{
		Name:     u.Name,
		Password: u.Password,
	}
}

type UserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ScheduleRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Place       string `json:"place"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
}

func (s *ScheduleRequest) ToSchedule() *Schedule {
	if s.Title == "" || s.StartTime == "" || s.EndTime == "" || s.Place == "" {
		return nil
	}

	return &Schedule{
		Title:       s.Title,
		Description: s.Description,
		Place:       s.Place,
		StartTime:   s.StartTime,
		EndTime:     s.EndTime,
	}
}
