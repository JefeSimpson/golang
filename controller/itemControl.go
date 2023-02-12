package controller

import (
	"fmt"
)

func (c *Collection) ItemPush(name string, price, rating float64) {
	c.ItemIterator++
	item := Item{c.ItemIterator, name, price, rating}
	c.Items = append(c.Items, item)
	fmt.Println("Item:", item, "was created successfully.")
}

func (c *Collection) SearchItemsByName(name string) []Item {
	var result []Item
	for _, item := range c.Items {
		if item.ItemName == name {
			result = append(result, item)
		}
	}
	return result
}

func (c *Collection) FilterItemsByPrice(price float64) []Item {
	var result []Item
	for _, item := range c.Items {
		if item.Price <= price {
			result = append(result, item)
		}
	}
	return result
}

func (c *Collection) FilterItemsByRating(rating float64) []Item {
	var result []Item
	for _, item := range c.Items {
		if item.Rating <= rating {
			result = append(result, item)
		}
	}
	return result
}

func (c *Collection) SetRating(rating float64, id int) {
	for _, item := range c.Items {
		if item.Id == id {
			item.Rating = rating
			fmt.Println("rating of item:", item, "has successfully changed.")
			return
		}
	}
	fmt.Println("item was not found.")
	return
}
