package dao

import "context"

type UserDAO struct {
}

func (u *UserDAO) Insert(ctx context.Context) {
}

type User struct {
	Id       int64
	Email    string
	Password string
	Ctime    int64
	Utime    int64
}
