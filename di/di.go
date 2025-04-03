package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}
func main() {
	Greet(os.Stdout, "Elodie")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
