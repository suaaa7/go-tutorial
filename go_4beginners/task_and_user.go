package main

import (
	"fmt"
)

type User struct {
	FirstName string
	LastName string
}

// User or *User ?
func (u *User) FullName() string {
	fullname := fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return fullname
}

func NewUser(firstName, lastName string) *User {
	return &User{
		FirstName: firstName,
		LastName: lastName,
	}
}

type Task struct {
	ID int
	Detail string
	done bool
	*User
}

func NewTask(
	id int,
	detail, firstName, lastName string) *Task {
	task := &Task{
		ID: id,
		Detail: detail,
		done: false,
		User: NewUser(firstName, lastName),
	}
	return task
}

func main() {
	fmt.Println("task and user")
	user := NewUser("user", "1")
	task := NewTask(1, "study Go", "user", "2")

	fmt.Println(user.LastName)
	fmt.Println(user.FullName())
	fmt.Println(user)

	fmt.Println(task.LastName)
	fmt.Println(task.FullName())
	fmt.Println(task.User)
}
