package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("t.gohtml"))
}

type hotel struct {
	Name, Address, City, Zip, Region string
}

type region struct {
	Name   string
	Hotels []hotel
}

type Regions []region

func main() {

	r := Regions{
		region{
			Name: "Southern",
			Hotels: []hotel{
				hotel{
					"hotel 1", "Angel Kunchev 17", "Plovdiv", "1000", "Southern",
				},
				hotel{
					"hotel 1", "Angel Kunchev 17", "Plovdiv", "1000", "Southern",
				},
			},
		},
		region{
			Name: "Central",
			Hotels: []hotel{
				hotel{
					"hotel 1", "Angel Kunchev 17", "Plovdiv", "1000", "Southern",
				},
				hotel{
					"hotel 1", "Angel Kunchev 17", "Plovdiv", "1000", "Southern",
				},
			},
		},
		region{
			Name: "Northern",
			Hotels: []hotel{
				hotel{
					"hotel 1", "Angel Kunchev 17", "Plovdiv", "1000", "Southern",
				},
				hotel{
					"hotel 1", "Angel Kunchev 17", "Plovdiv", "1000", "Southern",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
