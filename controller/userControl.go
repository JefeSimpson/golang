package controller

import (
	"fmt"
)

func (c *Collection) SignUp(username, password string) {
	c.UserIterator++
	user := User{c.UserIterator, username, password}
	c.Users = append(c.Users, user)
	fmt.Println("User:", user, "was created successfully.")
}

func (c *Collection) SignIn(username, password string) bool {
	for _, col := range c.Users {
		if col.Username == username && col.Password == password {
			fmt.Println("User:", username, "has authorized recently.")
			return true
		}
	}
	fmt.Println("User:", username, "failed the authorization.")
	return false
}

func (c *Collection) GetUser(username string) []User {
	var result []User
	for _, user := range c.Users {
		if user.Username == username {
			result = append(result, user)
		}
	}
	return result
}
