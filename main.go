package main

import (
	"fmt"
	"log"
	"net/http"
)

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "hello server")

// }
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "parsefrom() err: %v", err)
		return
	}
	fmt.Fprint(w, "POST request successful! \n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name= %s \n", name)
	fmt.Fprintln(w, "Address = %s", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found ", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not suppport", http.StatusNotFound)
		return

	}
	fmt.Fprintf(w, "hello")
}

func main() {
	//import file
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println(" Starting server at port 8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
