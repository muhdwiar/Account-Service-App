package entities

import "time"

type Balance struct {
	ID         int
	SALDO      int
	UPDATED_AT time.Time
}
