package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
w (ResponseWriter): This interface is used to construct the HTTP response. You write your data (like HTML or text) to this.
r (Request): This is a pointer to a struct containing all the details about the incoming request (URL, headers, method, etc.).
*/
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Path validation.
	if r.URL.Path != "/hello" { // Checks if the requested URL path is exactly '/hello'.
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	// Method validation.
	if r.Method != "GET" { // Ensures the HTTP method used is GET.
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!") // If the checks pass, it writes the string 'Hello!' to the response writer, which sends it back to the client.
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil { // r.ParseForm() reads the request body and URL query params to populate the r.Form map.
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful") // Writes a success message which is send back to the client.
	name := r.FormValue("name") // Retrieves the value associated with the key 'name' from the parsed form data (returns an empty string if the key doesn't exist).
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name) // Prints the name to the client.
	fmt.Fprintf(w, "Address = %s\n", address) // Prints the address to the client.
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))// Creates a handler that serves static files.
	http.Handle("/", fileServer) // Registers the file server handler for the root path.
	http.HandleFunc("/form", formHandler) // Registers a func named formHandler to handle requests send to the /form endpoint.
	http.HandleFunc("/hello", helloHandler) // Registers a func named helloHandler to handle requests send to the /hello endpoint. 

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil { // func that starts the web server. The second argument is the handler. Passing nil tells Go to use the DefaultServeMux (the default request multiplexer).
		log.Fatal(err) // Logs the error message with a timestamp and immediately terminates the pgm with a non-zero exit code.
	}
}