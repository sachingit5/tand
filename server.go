package main

import (

	"log"
	// "errors"
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"Hello World!")
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"Welcome to Home Page!!!")

}

func Sachin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"Welcome Sachin :)")
	

}

func iniT(){
	// http.HandleFunc("/hello/sachin",Sachin)
	http.HandleFunc("/hello",Hello)
	http.HandleFunc("/",Home)
	log.Fatal(http.ListenAndServe(":23431",nil))
}

func main(){
	iniT()
}




















// type helloHandler struct {
// 	name string
// }

// func (h helloHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
// 	w.Write([]byte("Hello "+h.name))
// }

// func main(){
// 	mux:=http.NewServeMux()
// 	var user helloHandler
// 	fmt.Scanf("%s",&user)

// 	helloSachin:=helloHandler{name:"Sachin"}
// 	helloSurya:=helloHandler{name:"Surya"}
// 	mux.Handle("/sachin",helloSachin)
// 	mux.Handle("/surya",helloSurya)

// 	http.ListenAndServe(":2343",mux)
// }