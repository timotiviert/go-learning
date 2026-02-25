package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/mail"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func New(name string, email string, age int) (*User, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    rand.Intn(100),
		Name:  name,
		Email: email,
		Age:   age,
	}, nil
}

func (u *User) IsAdult() bool {
	return u.Age >= 18
}

func (u *User) UpdateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}
	u.Email = email
	return nil
}

func UserToJSON(u *User) ([]byte, error) {
	jsonData, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func NewFromJSON(jsonData []byte) (*User, error) {
	var u User
	err := json.Unmarshal(jsonData, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func main() {
	gandalf, err := New("Gandalf the Grey", "mithrandir@wizzard.com", 2000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*gandalf)

	gandalfJSON, err := UserToJSON(gandalf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(gandalfJSON))

	frodoJSON := `{"username": "Frodo", "email": "frodo.baggins@shire.com", "age": 18}`
	frodo, err := NewFromJSON([]byte(frodoJSON))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*frodo)
}
