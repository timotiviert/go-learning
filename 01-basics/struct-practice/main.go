package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/mail"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func New(name string, email string, age int) *User {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return nil
	}
	return &User{
		Id:    rand.Intn(100),
		Name:  name,
		Email: email,
		Age:   age,
	}
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
	gandalf := New("Gandalf the Grey", "mithrandir@wizzard.com", 2000)
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
