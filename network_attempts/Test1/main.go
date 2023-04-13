package main

import (
	// "fmt"
	"net/http"
	"log"
	"example.com/vla/zz"
)



func main(){
	http.HandleFunc("/first", kokoa.Handler1)
	http.HandleFunc("/second", kokoa.Handler2)

	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Println("Listen and serve error : ", err)
	}

}