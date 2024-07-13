package domain

import "time"

type ModelTransaction struct {
	MtId            uint64
	RequestPrompt   []byte
	ResponseData    []byte
	TransactionDate *time.Time
}
