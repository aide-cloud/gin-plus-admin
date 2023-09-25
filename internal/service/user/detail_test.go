package user

import (
	"testing"
)

type AUser struct {
	ID uint `json:"id"`
}

func TestUser_GetDetail(t *testing.T) {
	t.Log(Afun[AUser]())
}

func Afun[T any]() *T {
	var m T
	return &m
}
