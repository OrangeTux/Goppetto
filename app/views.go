package app

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := loadTemplate("index")
	m := make(map[string]string)
	fmt.Fprint(w, template.HTML(executeTemplate(tmpl, m)))
}

func specs(w http.ResponseWriter, r *http.Request) {
	tmpl := loadTemplate("specs")
	m := make(map[string]string)
	fmt.Fprint(w, template.HTML(executeTemplate(tmpl, m)))
}

func loadTemplate(tmplName string) *template.Template {
	tmplFolder := os.Getenv("GOPATH") + "/src/github.com/OrangeTux/Goppetto/templates/"

	// Load template from file.
	data, err := ioutil.ReadFile(tmplFolder + tmplName + ".html")
	if err != nil {
		log.Fatalln(err)
	}

	tmpl, err := template.New(tmplName).Parse(string(data))
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Loaded template %s.", tmplName)

	return tmpl
}

func executeTemplate(tmpl *template.Template, v map[string]string) []byte {
	output := new(bytes.Buffer)
	tmpl.Execute(output, v)

	return output.Bytes()
}
