package storage

import "alejandro/data"

type userMysqlI interface {
	GetUserById(int) *data.User
}

type userMysql struct{}

func newUserMysql() *userMysql {
	return &userMysql{}
}

func (u *userMysql) GetUserById(id int) *data.User {
	return &data.User{}
}
