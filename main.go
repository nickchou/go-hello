package main

import (
	"fmt"
	"net/http"
)

func main() {
	//加一个测试，1
	fmt.Println("helloworld")
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8081", nil)
}

//Index  默认首页方法
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloworld！")
}
