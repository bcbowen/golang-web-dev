package main

import (
	"log"
	"os"
	"text/template"
)

// Region is the Hotel Region
type Region int

const (
	// Southern is in Southern California
	Southern Region = iota
	// Central is in Central California
	Central Region = iota
	// Northern is in Northern California
	Northern Region = iota
)

// Hotel contains info about a hotel
type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  Region
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	/*
			1. Create a data structure to pass to a template which
		* contains information about California hotels including Name, Address, City, Zip, Region
		* region can be: Southern, Central, Northern
		* can hold an unlimited number of hotels
	*/

	hotels := []Hotel{
		{
			Name:    "Mariott",
			Address: "412 Central Ave",
			City:    "Palo Flojo",
			Zip:     "23222",
			Region:  Southern,
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}

}
