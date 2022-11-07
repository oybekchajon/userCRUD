package main

import "time"

type User struct {
	Id int
	Name string
	Age int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
