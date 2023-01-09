package main

import "fmt"

type User interface {
	Rename(newName string)
	ChangeAge(newAge int)
	String() string
}

var _ fmt.Stringer = &user{} // Явно показывает имплементацию интерфейса Stringer
var _ User = &user{}

type user struct { // Структура, которая имплементирует Stringer
	Name string
	Age  int
}

func (u *user) Rename(newName string) {
	u.Name = newName
}

func (u *user) ChangeAge(newAge int) {
	u.Age = newAge
}

func (u user) String() string { // Имплементация Stringer
	return fmt.Sprintf("%s - %d", u.Name, u.Age)
}

func NewUser(name string, age int) User {
	return &user{
		Name: name,
		Age:  age}
}

func main() {
	u1 := NewUser("V", 12)
	u1.Rename("Valentin")
	u1.ChangeAge(25)
	fmt.Println(u1.String())

	u2 := NewUser("Y", 42)
	fmt.Println(u2.String())
}
