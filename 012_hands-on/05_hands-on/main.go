package main

import (
	"log"
	"os"
	"text/template"
)

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
	*/
	menu := Menu{
		BreakfastItems: []string{"Scrambled Eggs", "Fried Spam", "Eggs Benedict", "Eggs and Spam", "Eggs, Eggs, Spam, and Eggs"},
		LunchItems:     []string{"BLT", "Spam Sandwich", "Spam à l'orange", "Hamburger", "Spamburger"},
		DinnerItems:    []string{"Chicken Cordon Spam", "Lasagna", "Ham Spamwich", "Spaghetti and Snails", "Spam Flambé"},
	}

	err := tpl.Execute(os.Stdout, menu)

	if err != nil {
		log.Fatalln(err)
	}
}
