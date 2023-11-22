package main

import (
	"fmt"
	"log"
	"net/http"
)



func main(){
	 fileserver := http.FileServer((http.Dir("./static")))
	 http.Handle("/", fileserver)
	 http.HandleFunc("/form", formhandler)
	 http.HandleFunc("/hello", hellohandler)
	 fmt.Print("Starting server at port: 8080\n")

	 if err :=http.ListenAndServe(":8080", nil); err!=nil{
		log.Fatal(err)
	 }
	 
}

func hellohandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hell0")
}

func formhandler(w http.ResponseWriter, r *http.Request){
	if err:= r.ParseForm(); err!=nil{
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return 
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s", address)

}