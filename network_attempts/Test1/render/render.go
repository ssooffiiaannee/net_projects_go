package render

import (
	"log"
	"html/template"
	"net/http"
)


// func RenderTemplate(w http.ResponseWriter, tmpl string){
// 	parsedTemplate, _ := template.ParseFiles("templates/" + tmpl, "templates/base.layout.html")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil{
// 		log.Println("error : ", err)
// 	}
// }

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, tmpl string){
	parTmpl, exists := tc[tmpl]

	if !exists{
		log.Printf("Parsed tmpl created !")
		parTmpl = getParsedTemplates(tmpl)
	}else{
		log.Printf("Parsed tmpl exists !")
	}

	tc[tmpl] = parTmpl
	var err error = parTmpl.Execute(w, nil)
	if err != nil{
		log.Println("error : ", err)
	}

}

func getParsedTemplates(tmpl string) *template.Template{
	templates := []string{"templates/" + tmpl, "templates/base.layout.html"}
	parsedTemplate, _ := template.ParseFiles(templates...)
	
	return parsedTemplate
}