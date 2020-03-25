package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){
	p1:= doubleZero{
		person:        person{
			Name: "Ian Borat",
			Age:  56,
		},
		LicenseToKill: true,
	}
	err:= tpl.Execute(os.Stdout, p1)
	if err != nil{
		log.Fatalln(err)
	}
}