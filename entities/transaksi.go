package entities

import "time"

type Transaksi struct {
	ID               int
	USER_ID          int
	ACTION           string
	NOMINAL          int
	USER_ID_PENERIMA int
	CREATED_AT       time.Time
}
