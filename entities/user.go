package entities

import "time"

type User struct {
	ID         int
	NAMA       string
	NO_TELP    string
	PASSWORD   string
	CREATED_AT time.Time
	UPDATED_AT time.Time
}
