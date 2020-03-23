package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	logError(err)

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	logError(err)

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	logError(err)

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	logError(err)
}

func logError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
