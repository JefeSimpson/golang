package main

import (
	c "firstAssignment/controller"
	"fmt"
)

func main() {
	sms := c.Collection{
		Users:        make([]c.User, 0),
		Items:        make([]c.Item, 0),
		ItemIterator: 0,
		UserIterator: 0,
	}
	var name, password string
	//fmt.Print("Hello, this is mini-program, in which you can:\n1)sing up a user," +
	//	"\n2)authorize by it,\n3)manage items.\n")
	//fmt.Println("First, you have to create account, because you're gonna be the first user.")
	//fmt.Print("Type a username: ")
	//fmt.Scanln(&name)
	//fmt.Print("Type a password: ")
	//fmt.Scanln(&password)
	//sms.SignUp(name, password)
	//fmt.Println("To start program, choose the operation you want to do!")
	sms.UserTakeData()
	sms.ItemTakeData()
	var oper int
	//fmt.Scanln(&oper)
	isAuthorized := false
	for true {
		fmt.Print("1 - add another user\n2 - authorize\n3 - manage items\n4 - exit the program\n5 - log out\n")
		fmt.Scanln(&oper)
		if oper == 1 {
			fmt.Print("Type a new username: ")
			fmt.Scanln(&name)
			fmt.Print("Type a new password: ")
			fmt.Scanln(&password)
		} else if oper == 2 {
			fmt.Println("Type existing user to authorize")
			fmt.Print("Type a username: ")
			fmt.Scanln(&name)
			fmt.Print("Type a password: ")
			fmt.Scanln(&password)
			if sms.SignIn(name, password) {
				fmt.Println("Now, you can manage the items!")
				isAuthorized = true
			} else {
				fmt.Println("The data you have wrote is invalid! Try again.")
			}
		} else if oper == 3 {
			var itemOper int
			var itemName string
			var rating, price float64
			if isAuthorized {
				fmt.Println("1 - add item, 2 - filter item by price, 3 - filter item by rating, 4 - search item by name, 5 - set rating for item")
				fmt.Print("Choose operation for item: ")
				fmt.Scanln(&itemOper)
				if itemOper == 1 {
					fmt.Println("Write item you want to add")
					fmt.Print("Write name: ")
					fmt.Scanln(&itemName)
					fmt.Print("Write price: ")
					fmt.Scanln(&price)
					fmt.Print("Write rating: ")
					fmt.Scanln(&rating)
					sms.ItemPush(itemName, price, rating)
				} else if itemOper == 2 {
					fmt.Print("Write item price you want to sort: ")
					fmt.Scanln(&price)
					fmt.Println(sms.FilterItemsByPrice(price))
				} else if itemOper == 3 {
					fmt.Print("Write item rating you want to sort: ")
					fmt.Scanln(&rating)
					fmt.Println(sms.FilterItemsByRating(rating))
				} else if itemOper == 4 {
					fmt.Print("Write item name you want to find: ")
					fmt.Scanln(&itemName)
					fmt.Println(sms.SearchItemsByName(itemName))
				} else if itemOper == 5 {
					var itemId int
					fmt.Println("To set rating of existing item, you have to write id and rating.")
					fmt.Print("item id: ")
					fmt.Scanln(&itemId)
					fmt.Print("rating of item: ")
					fmt.Scanln(&rating)
					sms.SetRating(rating, itemId)
				} else {
					fmt.Println("You wrote incorrect operation for item managing!")
				}
			} else {
				fmt.Println("Firstly, you have to authorize. This operation is not available for you!")
			}
		} else if oper == 4 {
			fmt.Print("You're quitting from the program, are you sure. Type yes or no: ")
			var isSure string
			fmt.Scanln(&isSure)
			if isSure == "yes" {
				sms.UserSaveData()
				sms.ItemSaveData()
				break
			} else {
				continue
			}
		} else if oper == 5 {
			isAuthorized = false
			fmt.Println("You have just logged out.")
		} else {
			fmt.Println("Operation doesn't exist, try again.")
		}
	}
	//isAuthorized := sms.SignIn("Nuradil25", "123")
	//for isAuthorized {
	//	fmt.Println("You're authorized. What do you want to do?")
	//	fmt.Println("If you want to add items, then use 1")
	//	fmt.Println("If you want to filter items, then use 2")
	//	fmt.Println("If you want to search items, then use 3")
	//
	//}
}
