package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("helloworld")
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}

//Index  默认首页方法
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloworld！")
}
