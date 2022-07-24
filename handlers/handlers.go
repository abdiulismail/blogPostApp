package handlers

import (
	"fmt"
	"net/http"
)

func MyHelloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"hello world")
}

func GoodBye(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"goodbye")
}
