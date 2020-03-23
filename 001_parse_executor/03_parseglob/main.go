package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
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
