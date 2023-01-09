package user

import "fmt"

type User interface{}

var _ fmt.Stringer = &user{}
var _ User = &user{}

type user struct {
	name string
	age  int
}

func (u *user) Name() string {
	return u.name
}

func (u *user) SetName(name string) {
	u.name = name
}

func (u *user) Age() int {
	return u.age
}

func (u *user) SetAge(age int) {
	u.age = age
}

func NewUser(name string, age int) *user {
	return &user{name: name, age: age}
}

func NewEmptyUser() *user {
	return &user{name: "", age: 0}
}

func (u *user) String() string {
	return fmt.Sprintf("%s - %d", u.name, u.age)
}
