package main

import(
	"fmt"
	"log"
	"net/http"

)

// Given a form input, write out the form values
func formHandler(w http.ResponseWriter, r *http.Request) {

		// if unable to parse form, fail early
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Parseform() error : %v", err)
			return

		}

		fmt.Fprintf(w, "POST request successful\n")
		name := r.FormValue("name")
		address := r.FormValue("address")

		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)

}

// Given a form input, write out the form values
func samFormHandler(w http.ResponseWriter, r *http.Request) {

	// if unable to parse form, fail early
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() error : %v", err)
		return

	}

	fmt.Fprintf(w, "POST request successful\n")
	scale := r.FormValue("scale")

	fmt.Fprintf(w, "This is how much I love Sam ---> %v\n", scale)


}

// Say Hello :)
func helloHandler(w http.ResponseWriter, r *http.Request) {

	// expect PATH = "/hello"
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// by default it should be GET
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello!")

}

func main() {
	// tell golang to looks at the ./static folder for static files
	// fileServer serves HTTP requests within the given file system
	fileServer := http.FileServer(http.Dir("./static"))
	
	// handle the root server
	http.Handle("/", fileServer)

	// handle "/form"
	http.HandleFunc("/form", formHandler)

	// handle  "/hello"
	http.HandleFunc("/hello", helloHandler)

	// handle  "/samForm"
	http.HandleFunc("/samForm", samFormHandler)

	fmt.Println("Starting server at port 8080")

	// listen on TCP port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}	

}