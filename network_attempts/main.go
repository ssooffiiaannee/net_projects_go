package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		n, err := fmt.Fprintf(w, "hi, this is sofiane !")
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(n, " characters printed !")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		fmt.Println("Listen and serve error : ",err)
	}
}