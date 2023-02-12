package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func (c *Collection) UserTakeData() {
	file, err := os.Open("user.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		u := strings.Split(line, ",")
		n, err := strconv.Atoi(u[0])
		if err != nil {
			fmt.Println(err)
		}
		user := User{n, u[1], u[2]}
		c.Users = append(c.Users, user)
		c.UserIterator = n
	}
}

func (c *Collection) UserSaveData() {
	if err := os.Truncate("user.txt", 0); err != nil {
		fmt.Println(err)
	}
	file, err := os.OpenFile("user.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for _, user := range c.Users {
		if _, err := file.WriteString(strconv.Itoa(user.Id) + "," + user.Username + "," + user.Password + "\n"); err != nil {
			fmt.Println(err)
		}
	}
}
