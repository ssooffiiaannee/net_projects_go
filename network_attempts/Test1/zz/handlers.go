package kokoa

import(
	// "fmt"
	"net/http"
	"example.com/vla/render"
)

// name of functions in capital letter, it can be seen from outside the file.
func Handler1(w http.ResponseWriter, r *http.Request){
	// _, err := fmt.Fprintf(w, "this is the handler 1 page")
	// if err != nil{
	// 	log.Println("error : ", err)
	// }
	render.RenderTemplate(w, "first.page.html")
}

func Handler2(w http.ResponseWriter, r *http.Request){
	// _, err := fmt.Fprintf(w, "this is the handler 2 page")
	// if err != nil{
	// 	log.Println("error : ", err)
	// }
	render.RenderTemplate(w, "second.page.html")
}

// func RenderTemplate(w http.ResponseWriter, tmpl string){
// 	parsedTemplate, _ := template.ParseFiles("templates/" + tmpl, "templates/base.layout.html")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil{
// 		log.Println("error : ", err)
// 	}
// }