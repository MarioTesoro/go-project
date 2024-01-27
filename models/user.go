package models

import (
	"fmt"
	"html"
	"strconv"
	"strings"

	"hash/fnv"

	"go-project/utils/token"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string
	Username string
	Password string
}

type AuthInput struct {
	Username string
	Password string
}

func (u *User) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}

func VerifyPassword(password, hashedPassword string) bool {
	return strings.TrimSpace(password) == strings.TrimSpace(hashedPassword)
}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.Itoa(int(h.Sum32()))
}

func (u *User) SaveUser(db map[string]User) (*User, error) {

	u.ID = hash(u.Username + u.Password)
	//u.BeforeSave()
	db[u.ID] = *u
	//fmt.Println(db)
	return u, nil
}

func LoginCheck(u User, db map[string]User) (string, error) {

	var err error
	//fmt.Println(hash(u.Username + u.Password))
	user := db[hash(u.Username+u.Password)]
	//u.BeforeSave()
	fmt.Println(u.Password, user)
	res := VerifyPassword(u.Password, user.Password)
	fmt.Println(res)
	if res != true {
		return "", err
	}
	fmt.Println(user.ID)
	id, err := strconv.ParseUint(user.ID, 0, 64)
	fmt.Println(id)
	token, err := token.GenerateToken(uint32(id))
	if err != nil {
		return "", err
	}

	return token, nil

}
