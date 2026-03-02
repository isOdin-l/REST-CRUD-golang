package models

import "github.com/google/uuid"

type User struct {
	Id            uuid.UUID
	Name          string
	Username      string
	Password_hash string
}

type List struct {
	Id          uuid.UUID
	Author_id   uuid.UUID
	Title       string
	Description string
}

type Item struct {
	Id          uuid.UUID
	List_id     uuid.UUID
	Title       string
	Description string
	Done        bool
}
