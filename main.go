package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	//Parse after sumitting form in "form.html" has error
       if err:= r.ParseForm(); err != nil{
		fmt.Fprintf(w,"ParseForm() err: %v",err)
		return
	   }
	   fmt.Fprintf(w,"POST request successful")
	   name:=r.FormValue("name");
	   address:=r.FormValue("address");
 
	   fmt.Fprintf(w,"Name = %s\n",name)
	   fmt.Fprintf(w,"Address = %s\n",address)

}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	//if request from user is not exact as "/hello" so
	if r.URL.Path != "/hello" {
		//then response will be Page not found
		http.Error(w, "404 Page not Found", http.StatusNotFound)
		return
	}
	//we want only GET method to be used so
	//if any other method is used it will return error
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	//In ideal condition it will print "hello!" on the screen 
	fmt.Fprintf(w,"hello!")
}

func main() {
	//We are telling golang that we want to check out the static directory
	fileServer := http.FileServer(http.Dir("./static"))
	//Handling Root Route
	http.Handle("/", fileServer)
	//Handling Form Route
	http.HandleFunc("/form", formHandler)
	//Handling Hello Route
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080 \n")
	//Creating Server
	//It will either redirect to 8080 port or show error --> nil
	if err := http.ListenAndServe(":8080", nil);
	//If error is not nil it will show the other error
	err != nil {
		log.Fatal(err)
	}
}
