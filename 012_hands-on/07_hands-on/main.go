package main

import (
	"log"
	"os"
	"text/template"
)

type Restaurant struct {
	Name string
	Menu Menu
}

type Menu struct {
	BreakfastItems []string
	LunchItems     []string
	DinnerItems    []string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	/*
			1. Create a data structure to pass to a template which
		 contains information about restaurant's menu including Breakfast, Lunch,
		and Dinner items

		... Using the data structure created in the previous folder, modify it to
		hold menu information for an unlimited number of restaurants.
	*/

	restaurants := []Restaurant{
		{
			Name: "Dennys",
			Menu: Menu{
				BreakfastItems: []string{"Scrambled Eggs", "Fried Spam", "Eggs Benedict", "Eggs and Spam", "Eggs, Eggs, Spam, and Eggs"},
				LunchItems:     []string{"BLT", "Spam Sandwich", "Spam à l'orange", "Hamburger", "Spamburger"},
				DinnerItems:    []string{"Chicken Cordon Spam", "Lasagna", "Ham Spamwich", "Spaghetti and Snails", "Spam Flambé"},
			},
		},
		{
			Name: "Lennys",
			Menu: Menu{
				BreakfastItems: []string{"Scrambled Spams", "Fried Taters", "Eggs n Browns", "Eggs and Tacos", "Eggs, Eggs, Tacos, and Eggs"},
				LunchItems:     []string{"SOS", "Hard Tacos", "Soft Tacos", "Tacos a la Chalupa", "Tacoburger"},
				DinnerItems:    []string{"Chicken Flaming Tacos", "Taco Explosion", "Spam Tacos", "Spaghetti and Tacos", "Spam Under Glass"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, restaurants)

	if err != nil {
		log.Fatalln(err)
	}
}
