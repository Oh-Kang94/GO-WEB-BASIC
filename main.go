package main

import (
	"fmt"
	"net/http"
)

func main() {
	// "/"경로로 접속했을때 처리할 핸들러 처리
	http.HandleFunc("/", printHelloWorld)

	http.ListenAndServe(":8000", nil)
}

func printHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println(fmt.Fprintln(w, "welcome!"))
}
