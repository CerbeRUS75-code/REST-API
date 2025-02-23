package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var counter int

func GetHendler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		fmt.Fprintln(w, "Counter равен", strconv.Itoa(counter))
	} else {
		fmt.Fprintln(w, "Поддерживается только метод GET")

	}
}
func PostHendler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		counter++
		fmt.Fprintln(w, "Counter увеличен на 1")
	} else {
		fmt.Fprintln(w, "Поддерживается только метод POST")

	}
}

func DelHendler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		counter--
		fmt.Fprintln(w, "Counter уменьшен на 1")
	} else {
		fmt.Fprintln(w, "Поддерживается только метод DELETE")

	}
}

func main() {
	http.HandleFunc("/get", GetHendler)
	http.HandleFunc("/post", PostHendler)
	http.HandleFunc("/delete", DelHendler)
	http.ListenAndServe("localhost:8080", nil)

}
