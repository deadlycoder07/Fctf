package models

import "time"

type Users struct {
	Id        int `pg:"id,pk"`
	Name      string
	Email     string `pg:"email,unique"`
	Password  string
	Type      string
	Website   string
	Country   string
	Hidden    bool
	Verified  bool
	Banned    bool
	TeamId    int
	CreatedAt time.Time
}
