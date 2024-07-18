package domain

import "time"

type User struct {
	UserId         uint64
	UserName       string
	PassWord       string
	Email          string
	UserRank       uint64
	IncorrectLogin uint64
	LastLogin      *time.Time
	CreateDate     *time.Time
}
