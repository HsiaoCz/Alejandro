package api

import (
	"alejandro/storage"
)

type userApi struct {
	store *storage.Storage
}

func NewUserApi(store *storage.Storage) *userApi {
	return &userApi{store: store}
}
