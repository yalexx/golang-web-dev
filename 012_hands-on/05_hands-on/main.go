package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type menu struct {
	Meal  string
	Meals []meal
}

type meal struct {
	Name  string
	Items []menuItem
}

type menuItem struct {
	Name  string
	Price int
}

func init() {
	tpl = template.Must(template.ParseFiles("t.gohtml"))
}

func main() {

	m := menu{

		Meals: []meal{
			meal{
				Name: "Breakfast",
				Items: []menuItem{
					menuItem{
						"Bread", 10,
					},
					menuItem{
						"Eggs", 20,
					},
				},
			},
			meal{
				Name: "Lunch",
				Items: []menuItem{
					menuItem{
						"Kur", 5,
					},
					menuItem{
						"Hui", 15,
					},
				},
			},
			meal{
				Name: "Dinner",
				Items: []menuItem{
					menuItem{
						"Bum", 1,
					},
					menuItem{
						"Bam", 2,
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
