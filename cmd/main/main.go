package main

import (
	. "affordmed/fibb"
	. "affordmed/restapi"
	"fmt"
	"net/http"
)

var Memory map[int]int

func init() {
	Memory = make(map[int]int)
}

func main() {
	fmt.Println(Fibbo(8, &Memory))
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/login", LoginHandler)
	http.ListenAndServe(":8080", nil)
}
