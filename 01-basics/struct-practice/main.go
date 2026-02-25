package main

import (
	"math/rand"
	"net/mail"
)

type User struct {
	Id    int
	Name  string
	Email string
	Age   int
}

func (u User) IsAdult() bool {
	if u.Age >= 18 {
		return true
	}
	return false
}

func (u *User) UpdateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}
	u.Email = email
	return nil
}

func New(name, email string, ege int) (*User, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return nil, err
	}
	return &User{
		Id:    rand.Intn(100),
		Name:  name,
		Email: email,
	}, nil
}
